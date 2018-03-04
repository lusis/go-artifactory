BINARIES := $(shell find cmd/ -maxdepth 1 -type d -name 'artif-*' -exec sh -c 'echo $(basename {})' \;)
BINLIST := $(subst cmd/,,$(BINARIES))

ifeq ($(TRAVIS_BUILD_DIR),)
	GOPATH := $(GOPATH)
else
	GOPATH := $(GOPATH):$(TRAVIS_BUILD_DIR)
endif

all: bindata clean lint test artifactory $(BINLIST)

linux: export GOOS=linux
linux: all linux-zip

osx: export GOOS=darwin
osx: clean artifactory $(BINLIST) osx-zip

lint:
	@script/lint

test:
	@script/test

artifactory:
	@echo "Building for $(GOOS)"
	@script/build

$(BINLIST):
	@echo $@
	@go build -o bin/$@ ./cmd/$@

osx-zip:
	@mkdir target || echo "directory already exists"
	@zip -j target/artifactory-tools-osx-`date +%s`.zip bin/darwin_amd64/*

linux-zip:
	@mkdir target || echo "directory already exists"
	@zip -j target/artifactory-tools-linux-`date +%s`.zip bin/*

bindata:
	@go get -u github.com/jteeuwen/go-bindata/...
	@cd ${GOPATH}/src/github.com/lusis/go-artifactory/pkg/artifactory/responses/testdata;  go-bindata -pkg testdata -o testdata.go *.json *.txt; cd -

clean:
	@rm -rf bin/ pkg/

.PHONY: all clean lint test artifactory osx-zip linux-zip osx linux bindata $(BINLIST)
