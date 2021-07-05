PKG_NAME     ?= gmailfilter
WEBSITE_REPO  = github.com/hashicorp/terraform-website
BUILD_LDFLAGS := "-s -w -X github.com/yamamoto-febc/terraform-provider-gmailfilter/gmailfilter.Revision=`git rev-parse --short HEAD`"

default: fmt goimports lint tflint docscheck

clean:
	rm -Rf $(CURDIR)/terraform-provider-gmailfilter

.PHONY: tools
tools:
	go install golang.org/x/tools/cmd/goimports@v0.1.4
	go install github.com/bflad/tfproviderdocs@v0.9.1
	go install github.com/bflad/tfproviderlint/cmd/tfproviderlintx@v0.27.0
	go install github.com/client9/misspell/cmd/misspell@v0.3.4
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin v1.24.0


.PHONY: build
build:
	GOOS=$${OS:-"`go env GOOS`"} GOARCH=$${ARCH:-"`go env GOARCH`"} CGO_ENABLED=0 go build -ldflags=$(BUILD_LDFLAGS)

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
        -AT001 -AT002 -AT003 -AT004 -AT005 -AT006 -AT007 -AT008 -AT009 \
        -R001 -R002 -R004 -R005 -R006 -R007 -R008 -R009 -R010 -R011 -R012 -R013 -R014 -R015 \
        -R016 -R017       -R019 \
        -S001 -S002 -S003 -S004 -S005 -S006 -S007 -S008 -S009 -S010 -S011 -S012 -S013 -S014 -S015 \
        -S016 -S017 -S018 -S019 -S020 -S021 -S022 -S023 -S024 -S025 -S026 -S027 -S028 -S029 -S030 \
        -S031 -S032 -S033 -S034 -S035 -S036 -S037 \
        -V001 -V002 -V003 -V004 -V005 -V006 -V007 -V008 -V009 -V010 \
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

