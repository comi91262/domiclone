
.PHONY: build
build: ./cmd/domilike/main.go
	go build -o domilike ./cmd/domilike

.PHONY: goa
goa:
	goa gen github.com/comi91262/domilike/design
