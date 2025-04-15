package services

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"path/filepath"

	"github.com/soat-46/ms-mail-sender/internal/mail/domain/entities"
)

var (
	ErrParseTemplate  = errors.New("failed to parse template")
	ErrRenderTemplate = errors.New("failed to render template")
)

type RenderMailTemplate struct {
}

func NewRenderMailTemplate() *RenderMailTemplate {
	return &RenderMailTemplate{}
}

func (itself *RenderMailTemplate) Execute(mailType entities.EmailType) (string, error) {
	templateName := itself.getTemplateName(mailType)
	templatePath := filepath.Join("templates", fmt.Sprintf("%s.html", templateName))

	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", fmt.Errorf("%w: %w", ErrParseTemplate, err)
	}

	var buf bytes.Buffer
	if err = tmpl.Execute(&buf, nil); err != nil {
		return "", fmt.Errorf("%w: %w", ErrRenderTemplate, err)
	}
	return buf.String(), nil
}

func (itself *RenderMailTemplate) getTemplateName(mailType entities.EmailType) string {
	if mailType == entities.Success {
		return "mail_success"
	}
	return "mail_fail"
}
