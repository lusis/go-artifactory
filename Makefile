BINARIES = artif-list-repos artif-get-repo artif-get-license artif-list-users artif-get-user artif-list-groups artif-get-group artif-list-permission-targets artif-get-permission-target artif-create-user artif-delete-user artif-deploy-artifact

ifeq ($(TRAVIS_BUILD_DIR),)
	GOPATH := $(GOPATH)
else
	GOPATH := $(GOPATH):$(TRAVIS_BUILD_DIR)
endif
all: clean test artifactory artifactory-bin

test:
	@go test artifactory.v401 -test.v

artifactory:
	@mkdir -p bin/
	@go get ./... 
	@go install artifactory.v401

artifactory-bin:
	@mkdir -p bin/
	@$(foreach bin,$(BINARIES),go install $(bin);)

clean:
	@rm -rf bin/ pkg/

.PHONY: all clean test artifactory artifactory-bin
