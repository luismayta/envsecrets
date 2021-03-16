# go
.PHONY: go.help

# Bin variables
GOLANGCI-LINT = $(GOBIN)/golangci-lint

go.help:
	@echo '    go:'
	@echo ''
	@echo '        go                 show help'
	@echo '        go.lint            lint go'
	@echo '        go.setup           setup go'
	@echo '        go.fix             fix code of golangci lint'
	@echo '        go.vet             go vet against code'
	@echo '        go.fmt             run fmt for files'
	@echo '        go.build           build application'
	@echo ''

go:
	@if [ -z "${command}" ]; then \
		make go.help;\
	fi

bin/golangci-lint-${GOLANGCI_VERSION}:
	@mkdir -p bin
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | BINARY=golangci-lint bash -s -- v${GOLANGCI_VERSION}
	@mv bin/golangci-lint $@

bin/golangci-lint: bin/golangci-lint-${GOLANGCI_VERSION}
	@ln -sf golangci-lint-${GOLANGCI_VERSION} bin/golangci-lint

.PHONY: go.lint
go.lint: bin/golangci-lint ## Run linter
	bin/golangci-lint run

.PHONY: go.fix
go.fix: bin/golangci-lint ## Fix lint violations
	bin/golangci-lint run --fix

# Run go vet against code
go.vet:
	go vet ./...
.PHONY: go.vet

go.build: bin/goreleaser
	bin/goreleaser build --snapshot --rm-dist
.PHONY: go.build

# gofmt and goimports all go files
go.fmt:
	gofmt -s -l -w $(PROJECT_BUILD_SRCS)
.PHONY: go.fmt

# setup download and install dependence.
go.setup:
	go mod download
	go mod tidy
	go mod vendor
	go generate -v ./...
.PHONY: go.setup