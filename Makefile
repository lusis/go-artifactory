BINARIES := $(shell find src/ -maxdepth 1 -type d -name 'artif-*' -exec sh -c 'echo $(basename {})' \;)
BINLIST := $(subst src/,,$(BINARIES))

ifeq ($(TRAVIS_BUILD_DIR),)
	GOPATH := $(GOPATH)
else
	GOPATH := $(GOPATH):$(TRAVIS_BUILD_DIR)
endif

all: clean test artifactory $(BINLIST)

osx: clean artifactory $(BINLIST) osx-zip

test:
	@go get -t -d ./...
	@go test artifactory.v401 -v #-test.v
	@go test artifactory.v491 -v #-test.v

artifactory:
	@go get -t -d ./... 
	@go install artifactory.v401
	@go install artifactory.v491

$(BINLIST):
	@echo $@
	@go install $@

osx-zip:
	@mkdir target
	@zip -j target/artifactory-tools-`date +%s`.zip bin/darwin_386/*
	
clean:
	@rm -rf bin/ pkg/ src/github.com src/gopkg.in target/

.PHONY: all clean test artifactory osx-zip $(BINLIST)
