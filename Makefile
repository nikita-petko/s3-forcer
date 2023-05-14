IMAGE_TAG := $(shell git rev-parse HEAD 2> /dev/null)

ifeq ($(strip $(WITH_STDERR)),)
	TO_NUL := 2> /dev/null
else
	TO_NUL :=
endif

PROJECT_NAME ?= $(notdir $(CURDIR))

ifeq ($(OS),Windows_NT)
	SETC := SET
	ARCH := $(PROCESSOR_ARCHITECTURE)

	GOOS := windows
else 

	SETC := export 

	UNAME := $(shell uname)
	ARCH := $(shell uname -m)


	ifeq ($(UNAME),Linux)
		GOOS := linux
	else ifeq ($(UNAME),Darwin)
		GOOS := darwin
	endif

endif

# Convert the ARCH value to the corresponding GOARCH string
ifeq ($(ARCH),x86_64)
	GOARCH := amd64
else ifeq ($(ARCH),i386)
	GOARCH := 386
else ifeq ($(ARCH),armv7l)
	GOARCH := arm
else ifeq ($(ARCH),aarch64)
	GOARCH := arm64
else ifeq ($(ARCH),AMD64)
	GOARCH := amd64
else ifeq ($(ARCH),x86)
	GOARCH := 386
endif

ARCH_OUT = $(GOARCH)

ifeq ($(ARCH_OUT),386)
	ARCH_OUT = x86
else ifeq ($(ARCH_OUT),amd64)
	ARCH_OUT = x64
endif

ifeq ($(GOOS),windows)
	OUTFILE := $(PROJECT_NAME).exe
else
	OUTFILE := $(PROJECT_NAME)
endif

.DEFAULT_GOAL := all

all:
	$(info  make <cmd>)
	$(info )
	$(info commands:)
	$(info )
	$(info  vendor               - compile the vendor paths.)
	$(info  build-debug          - build a debug build. If GOARCH or GOOS were not specified "$(GOOS)/$(GOARCH)" will be used.)
	$(info  build-release        - build a release build that strips all symbols. If GOARCH or GOOS were not specified, "$(GOOS)/$(GOARCH)" will be used.)
	$(info  build-debug-<arch>   - helper for building a debug build for a specific arch, <arch> can be either x86, x64, arm or arm64. If GOOS was not specified, "$(GOOS)" will be used.)
	$(info  build-release-<arch> - helper for building a release build for a specific arch that strips all symbols, <arch> can be either x86, x64, arm or arm64. If GOOS was not specified, "$(GOOS)" will be used.)
	$(info )

vendor:
	$(info Go Vendor being applied to $(CURDIR)/src/vendor.)

ifeq ($(strip $(SKIP_VENDOR)),)
	@cd ./src && go mod tidy && go mod vendor
else
	$(info Go Vendor skipped because SKIP_VENDOR specified.)
endif

build-debug-x86: vendor
	$(eval GOARCH := 386)
	$(info Go Debug Build For is now set to: $(GOOS)/$(GOARCH) -> $(CURDIR)/bin/debug/$(GOOS)/x86/$(OUTFILE))
	@cd ./src && export CGO_ENABLED=1&& export GOARCH=$(GOARCH)&& export GOOS=$(GOOS)&& go build -mod=vendor -race -ldflags "-X main.commitSha=$(IMAGE_TAG) -X main.buildMode=debug -X main.applicationName=$(PROJECT_NAME)" -v -o ../bin/debug/$(GOOS)/x86/$(OUTFILE) ./main.go $(TO_NUL) || echo Go Build returned a non-zero status code

build-debug-x64: vendor
	$(eval GOARCH := amd64)
	$(info Go Debug Build For is now set to: $(GOOS)/$(GOARCH) -> $(CURDIR)/bin/debug/$(GOOS)/x64/$(OUTFILE))
	@cd ./src && export CGO_ENABLED=1&& export GOARCH=$(GOARCH)&& export GOOS=$(GOOS)&& go build -mod=vendor -race -ldflags "-X main.commitSha=$(IMAGE_TAG) -X main.buildMode=debug -X main.applicationName=$(PROJECT_NAME)" -v -o ../bin/debug/$(GOOS)/x64/$(OUTFILE) ./main.go $(TO_NUL) || echo Go Build returned a non-zero status code

build-debug-arm: vendor
	$(eval GOARCH := arm)
	$(info Go Debug Build For is now set to: $(GOOS)/$(GOARCH) -> $(CURDIR)/bin/debug/$(GOOS)/arm/$(OUTFILE))
	@cd ./src && export CGO_ENABLED=1&& export GOARCH=$(GOARCH)&& export GOOS=$(GOOS)&& go build -mod=vendor -race -ldflags "-X main.commitSha=$(IMAGE_TAG) -X main.buildMode=debug -X main.applicationName=$(PROJECT_NAME)" -v -o ../bin/debug/$(GOOS)/arm/$(OUTFILE) ./main.go $(TO_NUL) || echo Go Build returned a non-zero status code

build-debug-arm64: vendor
	$(eval GOARCH := arm64)
	$(info Go Debug Build For is now set to: $(GOOS)/$(GOARCH) -> $(CURDIR)/bin/debug/$(GOOS)/arm64/$(OUTFILE))
	@cd ./src && export CGO_ENABLED=1&& export GOARCH=$(GOARCH)&& export GOOS=$(GOOS)&& go build -mod=vendor -race -ldflags "-X main.commitSha=$(IMAGE_TAG) -X main.buildMode=debug -X main.applicationName=$(PROJECT_NAME)" -v -o ../bin/debug/$(GOOS)/arm64/$(OUTFILE) ./main.go $(TO_NUL) || echo Go Build returned a non-zero status code


build-release-x86: vendor
	$(eval GOARCH := 386)
	$(info Go Release Build For is now set to: $(GOOS)/$(GOARCH) -> $(CURDIR)/bin/release/$(GOOS)/x86/$(OUTFILE))
	@cd ./src && export CGO_ENABLED=1&& export GOARCH=$(GOARCH)&& export GOOS=$(GOOS)&& go build -mod=vendor -ldflags "-s -w -X main.commitSha=$(IMAGE_TAG) -X main.buildMode=release -X main.applicationName=$(PROJECT_NAME)" -v -o ../bin/release/$(GOOS)/x86/$(OUTFILE) ./main.go $(TO_NUL) || echo Go Build returned a non-zero status code

build-release-x64: vendor
	$(eval GOARCH := amd64)
	$(info Go Release Build For is now set to: $(GOOS)/$(GOARCH) -> $(CURDIR)/bin/release/$(GOOS)/x64/$(OUTFILE))
	@cd ./src && export CGO_ENABLED=1&& export GOARCH=$(GOARCH)&& export GOOS=$(GOOS)&& go build -mod=vendor -ldflags "-s -w -X main.commitSha=$(IMAGE_TAG) -X main.buildMode=release -X main.applicationName=$(PROJECT_NAME)" -v -o ../bin/release/$(GOOS)/x64/$(OUTFILE) ./main.go $(TO_NUL) || echo Go Build returned a non-zero status code

build-release-arm: vendor
	$(eval GOARCH := arm)
	$(info Go Release Build For is now set to: $(GOOS)/$(GOARCH) -> $(CURDIR)/bin/release/$(GOOS)/arm/$(OUTFILE))
	@cd ./src && export CGO_ENABLED=1&& export GOARCH=$(GOARCH)&& export GOOS=$(GOOS)&& go build -mod=vendor -ldflags "-s -w -X main.commitSha=$(IMAGE_TAG) -X main.buildMode=release -X main.applicationName=$(PROJECT_NAME)" -v -o ../bin/release/$(GOOS)/arm/$(OUTFILE) ./main.go $(TO_NUL) || echo Go Build returned a non-zero status code

build-release-arm64: vendor
	$(eval GOARCH := arm64)
	$(info Go Release Build For is now set to: $(GOOS)/$(GOARCH) -> $(CURDIR)/bin/release/$(GOOS)/arm64/$(OUTFILE))
	@cd ./src && export CGO_ENABLED=1&& export GOARCH=$(GOARCH)&& export GOOS=$(GOOS)&& go build -mod=vendor -ldflags "-s -w -X main.commitSha=$(IMAGE_TAG) -X main.buildMode=release -X main.applicationName=$(PROJECT_NAME)" -v -o ../bin/release/$(GOOS)/arm64/$(OUTFILE) ./main.go $(TO_NUL) || echo Go Build returned a non-zero status code

build-debug: vendor
	$(info Go Debug Build For is now set to: $(GOOS)/$(GOARCH) -> $(CURDIR)/bin/debug/$(GOOS)/$(ARCH_OUT)/$(OUTFILE))
	@cd ./src && export CGO_ENABLED=1&& go build -mod=vendor -race -ldflags "-X main.commitSha=$(IMAGE_TAG) -X main.buildMode=debug -X main.applicationName=$(PROJECT_NAME)" -v -o ../bin/debug/$(GOOS)/$(ARCH_OUT)/$(OUTFILE) ./main.go $(TO_NUL) || echo Go Build returned a non-zero status code

build-release: vendor
	$(info Go Release Build For is now set to: $(GOOS)/$(GOARCH) -> $(CURDIR)/bin/release/$(GOOS)/$(ARCH_OUT)/$(OUTFILE))
	@cd ./src && export CGO_ENABLED=1&& go build -mod=vendor -ldflags "-s -w -X main.commitSha=$(IMAGE_TAG) -X main.buildMode=release -X main.applicationName=$(PROJECT_NAME)" -v -o ../bin/release/$(GOOS)/$(ARCH_OUT)/$(OUTFILE) ./main.go $(TO_NUL) || echo Go Build returned a non-zero status code

build-debug-all: build-debug-x86 build-debug-x64 build-debug-arm build-debug-arm64
build-release-all: build-release-x86 build-release-x64 build-release-arm build-release-arm64

