PKG_NAME     ?= gmailfilter
WEBSITE_REPO  = github.com/hashicorp/terraform-website

default: fmt goimports lint tflint docscheck

clean:
	rm -Rf $(CURDIR)/bin/*

.PHONY: tools
tools:
	GO111MODULE=off go get github.com/x-motemen/gobump/cmd/gobump
	GO111MODULE=off go get golang.org/x/tools/cmd/goimports
	GO111MODULE=off go get github.com/tcnksm/ghr
	GO111MODULE=off go get github.com/bflad/tfproviderdocs
	GO111MODULE=off go get github.com/bflad/tfproviderlint/cmd/tfproviderlintx
	GO111MODULE=off go get github.com/client9/misspell/cmd/misspell
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin v1.24.0


.PHONY: build-envs
build-envs:
	$(eval CURRENT_VERSION ?= $(shell gobump show -r gmailfilter/))
	$(eval BUILD_LDFLAGS := "-s -w \
           -X github.com/yamamoto-febc/terraform-provider-gmailfilter/gmailfilter.Revision=`git rev-parse --short HEAD` \
           -X github.com/yamamoto-febc/terraform-provider-gmailfilter/gmailfilter.Version=$(CURRENT_VERSION)")

.PHONY: build
build: build-envs
	OS=$${OS:-"`go env GOOS`"} ARCH=$${ARCH:-"`go env GOARCH`"} BUILD_LDFLAGS=$(BUILD_LDFLAGS) CURRENT_VERSION=$(CURRENT_VERSION) sh -c "'$(CURDIR)/scripts/build.sh'"

.PHONY: build-x
build-x: build-envs build-darwin build-windows build-linux shasum

.PHONY: build-darwin
build-darwin: build-envs bin/terraform-provider-gmailfilter_$(CURRENT_VERSION)_darwin-386.zip bin/terraform-provider-gmailfilter_$(CURRENT_VERSION)_darwin-amd64.zip

.PHONY: build-windows
build-windows: build-envs bin/terraform-provider-gmailfilter_$(CURRENT_VERSION)_windows-386.zip bin/terraform-provider-gmailfilter_$(CURRENT_VERSION)_windows-amd64.zip

.PHONY: build-linux
build-linux: build-envs bin/terraform-provider-gmailfilter_$(CURRENT_VERSION)_linux-386.zip bin/terraform-provider-gmailfilter_$(CURRENT_VERSION)_linux-amd64.zip

bin/terraform-provider-gmailfilter_$(CURRENT_VERSION)_darwin-386.zip: build-envs
	OS="darwin"  ARCH="386"   ARCHIVE=1 BUILD_LDFLAGS=$(BUILD_LDFLAGS) CURRENT_VERSION=$(CURRENT_VERSION) sh -c "'$(CURDIR)/scripts/build.sh'"

bin/terraform-provider-gmailfilter_$(CURRENT_VERSION)_darwin-amd64.zip: build-envs
	OS="darwin"  ARCH="amd64" ARCHIVE=1 BUILD_LDFLAGS=$(BUILD_LDFLAGS) CURRENT_VERSION=$(CURRENT_VERSION) sh -c "'$(CURDIR)/scripts/build.sh'"

bin/terraform-provider-gmailfilter_$(CURRENT_VERSION)_windows-386.zip: build-envs
	OS="windows" ARCH="386"   ARCHIVE=1 BUILD_LDFLAGS=$(BUILD_LDFLAGS) CURRENT_VERSION=$(CURRENT_VERSION) sh -c "'$(CURDIR)/scripts/build.sh'"

bin/terraform-provider-gmailfilter_$(CURRENT_VERSION)_windows-amd64.zip: build-envs
	OS="windows" ARCH="amd64" ARCHIVE=1 BUILD_LDFLAGS=$(BUILD_LDFLAGS) CURRENT_VERSION=$(CURRENT_VERSION) sh -c "'$(CURDIR)/scripts/build.sh'"

bin/terraform-provider-gmailfilter_$(CURRENT_VERSION)_linux-386.zip: build-envs
	OS="linux"   ARCH="386"   ARCHIVE=1 BUILD_LDFLAGS=$(BUILD_LDFLAGS) CURRENT_VERSION=$(CURRENT_VERSION) sh -c "'$(CURDIR)/scripts/build.sh'"

bin/terraform-provider-gmailfilter_$(CURRENT_VERSION)_linux-amd64.zip: build-envs
	OS="linux"   ARCH="amd64" ARCHIVE=1 BUILD_LDFLAGS=$(BUILD_LDFLAGS) CURRENT_VERSION=$(CURRENT_VERSION) sh -c "'$(CURDIR)/scripts/build.sh'"

.PHONY: shasum
shasum:
	(cd bin/; shasum -a 256 * > terraform-provider-gmailfilter_$(CURRENT_VERSION)_SHA256SUMS)

.PHONY: release
release: build-envs
	ghr v${CURRENT_VERSION} bin/

.PHONY: test
test:
	TF_ACC= go test -v $(TESTARGS) -timeout=30s ./...

.PHONY: testacc
testacc:
	TF_ACC=1 go test -v $(TESTARGS) -timeout 240m ./...

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: tflint
tflint:
	tfproviderlintx \
        -AT001 -AT002 -AT003 -AT004 -AT005 -AT006 -AT007 -AT008 \
        -R001 -R002 -R004 -R005 -R006 -R007 -R008 -R009 -R010 -R011 -R012 -R013 -R014 \
        -S001 -S002 -S003 -S004 -S005 -S006 -S007 -S008 -S009 -S010 -S011 -S012 -S013 -S014 -S015 \
        -S016 -S017 -S018 -S019 -S020 -S021 -S022 -S023 -S024 -S025 -S026 -S027 -S028 -S029 -S030 \
        -S031 -S032 -S033 -S034 \
        -V001 -V002 -V003 -V004 -V005 -V006 -V007 -V008 \
        -XR001 -XR004 \
        ./$(PKG_NAME)

.PHONY: goimports
goimports:
	goimports -l -w $(PKG_NAME)/

.PHONY: fmt
fmt:
	find . -name '*.go' | grep -v vendor | xargs gofmt -s -w

.PHONY: docscheck
docscheck:
	tfproviderdocs check \
		-require-resource-subcategory \
		-require-guide-subcategory

.PHONY: website
website:
ifeq (,$(wildcard $(GOPATH)/src/$(WEBSITE_REPO)))
	echo "$(WEBSITE_REPO) not found in your GOPATH (necessary for layouts and assets), get-ting..."
	git clone https://$(WEBSITE_REPO) $(GOPATH)/src/$(WEBSITE_REPO)
	(cd $(GOPATH)/src/$(WEBSITE_REPO); \
	  git submodule init ext/terraform; git submodule update; \
	  ln -s ../../../ext/providers/gmailfilter/website/gmailfilter.erb content/source/layouts/gmailfilter.erb; \
	  ln -s ../../../../ext/providers/gmailfilter/website/docs content/source/docs/providers/gmailfilter \
	)
endif
	$(MAKE) -C $(GOPATH)/src/$(WEBSITE_REPO) website-provider PROVIDER_PATH=$(shell pwd) PROVIDER_NAME=$(PKG_NAME)

.PHONY: website-lint
website-lint:
	@echo "==> Checking website against linters..."
	misspell -error -source=text website/

.PHONY: website-test
website-test:
ifeq (,$(wildcard $(GOPATH)/src/$(WEBSITE_REPO)))
	echo "$(WEBSITE_REPO) not found in your GOPATH (necessary for layouts and assets), get-ting..."
	git clone https://$(WEBSITE_REPO) $(GOPATH)/src/$(WEBSITE_REPO)
	(cd $(GOPATH)/src/$(WEBSITE_REPO); \
	  git submodule init ext/terraform; git submodule update; \
	  ln -s ../../../ext/providers/gmailfilter/website/gmailfilter.erb content/source/layouts/gmailfilter.erb; \
	  ln -s ../../../../ext/providers/gmailfilter/website/docs source/docs/providers/gmailfilter \
	)
endif
	@$(MAKE) -C $(go env GOPATH)/src/$(WEBSITE_REPO) website-provider-test PROVIDER_PATH=$(shell pwd) PROVIDER_NAME=$(PKG_NAME)

.PHONY: website-scaffold
website-scaffold:
	go run tools/tfdocgen/cmd/gen-gmailfilter-docs/main.go website-scaffold

