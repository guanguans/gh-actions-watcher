# note: call scripts from /scripts

checks: mod-tidy fmt

golangci-lint:
	golangci-lint run ./...

gosec:
	gosec ./...

fmt:
	go fmt ./...

fumpt:
	gofumpt -d -e -l -w -extra .

vet:
	go vet ./...

mod-tidy:
	go mod tidy

test:
	go test -cover -coverprofile=cover.out -race ./... -v

test-cover:
	go tool cover -html=cover.out

bench:
	go test -bench=. -benchmem ./... -v

# goreleaser init
# goreleaser check
# goreleaser build --single-target
# goreleaser release --snapshot --rm-dist
# goreleaser release
goreleaser:
	goreleaser

staticcheck:
	staticcheck ./...

errcheck:
	errcheck ./...

ineffassign:
	ineffassign ./...

license-check:
	license-eye header check

license-fix:
	license-eye header fix

.PHONY: golangci-lint gosec fmt fumpt vet mod-tidy test test-cover bench goreleaser staticcheck errcheck ineffassign license-check license-fix