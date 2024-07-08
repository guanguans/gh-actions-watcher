# note: call scripts from /scripts
md-lint = lint-md --config .lintmdrc ./*.md ./.github/ ./docs/ ./src/*/*.md
release = monorepo-builder release --ansi -v

checks: mod-tidy goimports fumpt license-fix test
.PHONY: checks

ai-commit:
	ai-commit commit --generator=bito_cli --ansi
.PHONY: ai-commit

ai-commit-no-verify:
	ai-commit commit --generator=bito_cli --ansi --no-verify
.PHONY: ai-commit-no-verify

golangci-lint:
	golangci-lint run ./...
.PHONY: golangci-lint

gosec:
	gosec ./...
.PHONY: gosec

fmt:
	go fmt ./...
.PHONY: fmt

fumpt:
	gofumpt -d -e -l -w -extra .
.PHONY: fumpt

vet:
	go vet ./...
.PHONY: vet

goimports:
	goimports -w .
.PHONY: goimports

mod-tidy:
	go mod tidy
.PHONY: mod-tidy

test:
	go test -coverprofile=coverage.out -cover -race -v ./...
.PHONY: test

test-cover-view: test
	go tool cover -html=coverage.out
.PHONY: test-cover

bench:
	go test -bench=. -benchmem ./... -v
.PHONY: bench

# goreleaser init
# goreleaser check
# goreleaser build --single-target
# goreleaser release --snapshot --rm-dist
# goreleaser release
goreleaser:
	goreleaser
.PHONY: goreleaser

staticcheck:
	staticcheck ./...
.PHONY: staticcheck

errcheck:
	errcheck ./...
.PHONY: errcheck

ineffassign:
	ineffassign ./...
.PHONY: ineffassign

license-check:
	license-eye header check
.PHONY: license-check

license-fix:
	license-eye header fix
.PHONY: license-fix

vhs:
	vhs < gh-actions-watcher.tape
.PHONY: vhs

md-lint:
	$(md-lint)
.PHONY: lint-md

md-fix:
	$(md-lint) --fix
.PHONY: md-fix

trufflehog:
	trufflehog git https://github.com/guanguans/gh-actions-watcher --only-verified
.PHONY: trufflehog

release-major:
	$(release) major
release-major-dry-run:
	$(release)-major --dry-run
release-minor:
	$(release) minor
release-minor-dry-run:
	$(release)-minor --dry-run
release-patch:
	$(release) patch
release-patch-dry-run:
	$(release)-patch --dry-run
