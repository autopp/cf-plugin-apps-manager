.PHONY: test
test:
	go test ./...

.PHONY: build
build:
	goreleaser build --single-target --snapshot --rm-dist
