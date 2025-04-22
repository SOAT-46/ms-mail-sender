TOOLS_DIR := $(shell go env GOPATH)/bin

.PHONY: build
build:
	make wire
	go mod tidy
	mkdir -p bin && go build -o ./bin -v ./...

.PHONY: wire
wire:
	go install github.com/google/wire/cmd/wire@latest
	cd cmd; wire

.PHONY: test
test:
	.github/scripts/tests/run_tests.sh

.PHONY: docker-up
docker-up:
	docker compose -f ./.dev/compose.yaml up

.PHONY: lint
lint:
	wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh
	./bin/golangci-lint run --fix \
		--config ".golangci.yaml" \
		--color "always" \
		--timeout "10m" \
		--print-resources-usage \
		--allow-parallel-runners \
		--max-issues-per-linter 0 \
		--max-same-issues 0 ./...

.PHONY: goimports
goimports:
	go install golang.org/x/tools/cmd/goimports@latest
	goimports -l -w .

.PHONY: gci
gci:
	go install github.com/daixiang0/gci@latest
	gci write --skip-generated -s standard -s default .

.PHONY: lint-fix
lint-fix:
	make goimports
	make gci
	make lint

.PHONY: static-analysis
static-analysis: gosec staticcheck govulncheck

.PHONY: gosec
gosec:
	@echo "Running gosec security static analysis..."
	go install github.com/securego/gosec/v2/cmd/gosec@latest
	 gosec -fmt=json ./...

.PHONY: staticcheck
staticcheck:
	@echo "Running staticcheck..."
	go install honnef.co/go/tools/cmd/staticcheck@latest
	staticcheck ./...

.PHONY: govulncheck
govulncheck:
	@echo "Running govulncheck vulnerability analysis..."
	go install golang.org/x/vuln/cmd/govulncheck@latest
	govulncheck ./...

.PHONY: check-all
check-all: lint-fix static-analysis test
	@echo "All checks completed successfully!"
