package services_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/soat-46/ms-mail-sender/internal/mail/domain/entities"
	"github.com/soat-46/ms-mail-sender/internal/mail/infrastructure/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupTestTemplates(t *testing.T) (string, func()) {
	// Create a temporary directory for test templates
	tempDir := t.TempDir()

	// Create templates directory
	templatesDir := filepath.Join(tempDir, "templates")
	err := os.MkdirAll(templatesDir, 0755)
	require.NoError(t, err)

	// Create test template files
	successTemplate := `<!DOCTYPE html>
<html>
<body>
	<h1>Success Template</h1>
</body>
</html>`

	failTemplate := `<!DOCTYPE html>
<html>
<body>
	<h1>Fail Template</h1>
</body>
</html>`

	err = os.WriteFile(filepath.Join(templatesDir, "mail_success.html"), []byte(successTemplate), 0644)
	require.NoError(t, err)

	err = os.WriteFile(filepath.Join(templatesDir, "mail_fail.html"), []byte(failTemplate), 0644)
	require.NoError(t, err)

	// Change working directory to temp dir for the test
	originalWd, err := os.Getwd()
	require.NoError(t, err)
	err = os.Chdir(tempDir)
	require.NoError(t, err)

	// Return cleanup function
	cleanup := func() {
		cleanErr := os.Chdir(originalWd)
		if cleanErr != nil {
			t.Errorf("failed to change back to original directory: %v", cleanErr)
		}
		os.RemoveAll(tempDir)
	}
	return tempDir, cleanup
}

func TestRenderMailTemplate(t *testing.T) {
	t.Run("should render success template correctly", func(t *testing.T) {
		// given
		_, cleanup := setupTestTemplates(t)
		defer cleanup()

		service := services.NewRenderMailTemplate()

		// when
		result, err := service.Execute(entities.Success)

		// then
		require.NoError(t, err, "template file does not exist")
		assert.Contains(t, result, "Success Template")
	})

	t.Run("should render fail template correctly", func(t *testing.T) {
		// given
		_, cleanup := setupTestTemplates(t)
		defer cleanup()

		service := services.NewRenderMailTemplate()

		// when
		result, err := service.Execute(entities.Fail)

		// then
		require.NoError(t, err, "template file does not exist")
		assert.Contains(t, result, "Fail Template")
	})

	t.Run("should return error when template file does not exist", func(t *testing.T) {
		// given
		tempDir, cleanup := setupTestTemplates(t)
		defer cleanup()

		// Remove templates directory to simulate missing files
		err := os.RemoveAll(filepath.Join(tempDir, "templates"))
		require.NoError(t, err)

		service := services.NewRenderMailTemplate()

		// when
		result, err := service.Execute(entities.Success)

		// then
		require.Error(t, err, "template file does not exist")
		require.ErrorIs(t, err, services.ErrParseTemplate, "expected ErrParseTemplate error")
		assert.Empty(t, result, "result should be empty")
	})
}
