.PHONY: test
test:
	ginkgo ./...

.PHONY: build
build:
	goreleaser build --single-target --snapshot --rm-dist
