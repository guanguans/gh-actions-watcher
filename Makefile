# note: call scripts from /scripts

checks: mod-tidy goimports fmt

.PHONY: ai-commit
ai-commit:
	ai-commit commit --generator=bito_cli --ansi

.PHONY: ai-commit-no-verify
ai-commit-no-verify:
	ai-commit commit --generator=bito_cli --ansi --no-verify

.PHONY: golangci-lint
golangci-lint:
	golangci-lint run ./...

.PHONY: gosec
gosec:
	gosec ./...

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: fumpt
fumpt:
	gofumpt -d -e -l -w -extra .

.PHONY: vet
vet:
	go vet ./...

.PHONY: goimports
goimports:
	goimports -w .

.PHONY: mod-tidy
mod-tidy:
	go mod tidy

.PHONY: test
test:
	go test -cover -coverprofile=cover.out -race ./... -v

.PHONY: test-cover
test-cover:
	go tool cover -html=cover.out

.PHONY: bench
bench:
	go test -bench=. -benchmem ./... -v

.PHONY: goreleaser
# goreleaser init
# goreleaser check
# goreleaser build --single-target
# goreleaser release --snapshot --rm-dist
# goreleaser release
goreleaser:
	goreleaser

.PHONY: staticcheck
staticcheck:
	staticcheck ./...

.PHONY: errcheck
errcheck:
	errcheck ./...

.PHONY: ineffassign
ineffassign:
	ineffassign ./...

.PHONY: license-check
license-check:
	license-eye header check

.PHONY: license-fix
license-fix:
	license-eye header fix