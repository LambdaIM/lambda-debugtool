GOCMD = go
GOBUILD = $(GOCMD) build
GOMOD = $(GOCMD) mod

export GO111MODULE = on

all:build

build: go-mod-cache
	$(GOBUILD) -o build/lambdadg ./cmd/lambdadg

clean:
	rm -rf build/

go-mod-cache:
	$(GOMOD) download
