export PATH:=${PATH}:${GOPATH}/bin

.PHONY: all
all: build test

build: cmd/*.go *.go Makefile
	@echo "compile"
	@rm -rf build/ && mkdir -p build/ && \
	go build cmd/gmet_demo.go && \
	mv gmet_demo build/ && \
	cp cmd/seelog.xml build/

vendor: glide.lock glide.yaml
	@echo "download dependency"
	glide install

.PHONY: test
test:
	@echo "test"
	go test -cover *.go

.PHONY: clean
clean:
	rm -rf build

.PHONY: deep_clean
deep_clean:
	rm -rf vendor build
