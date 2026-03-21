# Makefile — Canon (library — no binary)
.PHONY: test vet

test:
	go test ./...

vet:
	go vet ./...
