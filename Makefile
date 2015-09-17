BINARIES = artif-list-projects

GOPATH := $(GOPATH):$(TRAVIS_BUILD_DIR)
all: clean test artifactory artifactory-bin

test:
	@go test artifactory.v401 -v

artifactory:
	@mkdir -p bin/
	@go get ./... 
	@go install artifactory.v401

artifactory-bin:
	@mkdir -p bin/
	$(foreach bin,$(BINARIES),go install $(bin);)

clean:
	@rm -rf bin/ pkg/

.PHONY: all clean test artifactory artifactory-bin
