# SPDX-License-Identifier: Apache-2.0
#
# -------------------------------------------------------------
# This makefile defines the following targets
#
#   - all (default) - formats the code, runs liners, downloads vendor libs, and builds executable
#   - fmt - formats the code
#   - vendor - download all third party libraries and puts them inside vendor directory
#   - clean-vendor - removes third party libraries from vendor directory
#   - build - builds the controller
#   - build-images - builds docker image locally for running the components using docker
#   - push-images - pushes the local docker image to docker registry
#   - clean - cleans the build directories
#   - clean-all - superset of 'clean' that also removes vendor dir
#   - lint - runs code analysis tools

COMPONENT := spot-instance-controller
IMAGE_TAG ?= latest
IMAGE := ${REGISTRY}/${COMPONENT}:${IMAGE_TAG}

.PHONY: all				##formats the code, runs liners, downloads vendor libs, and builds executable
all: vendor fmt build ## TODO: enable linting

.PHONY: fmt				##formats the code
fmt:
	@gci -w ./cmd/ ./pkg/
	@go fmt ./cmd/... ./pkg/...
	@gofumpt -w ./cmd/ ./pkg/

.PHONY: vendor			##download all third party libraries and puts them inside vendor directory
vendor:
	@go mod vendor

.PHONY: clean-vendor			##removes third party libraries from vendor directory
clean-vendor:
	-@rm -rf vendor

.PHONY: build			##builds the controller
build:
	@go build -o bin/${COMPONENT} cmd/manager/main.go

.PHONY: build-images			##builds docker image locally for running the components using docker
build-images: all
	docker build -t ${IMAGE} --build-arg COMPONENT=${COMPONENT} -f build/Dockerfile .

.PHONY: push-images			##pushes the local docker image to docker registry
push-images: build-images
	@docker push ${IMAGE}

.PHONY: clean			##cleans the build directories
clean:
	@rm -rf bin

.PHONY: clean-all			##superset of 'clean' that also removes vendor dir
clean-all: clean-vendor clean

.PHONY: lint				##runs code analysis tools
lint:
	go vet ./cmd/... ./pkg/...
	golangci-lint run ./cmd/... ./pkg/...

.PHONY: help				##show this help message
help:
	@echo "usage: make [target]\n"; echo "options:"; \fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//' | sed 's/.PHONY:*//' | sed -e 's/^/  /'; echo "";

## for CRD generation

##@ Development

.PHONY: manifests
manifests: controller-gen ## Generate WebhookConfiguration, ClusterRole and CustomResourceDefinition objects.
	$(CONTROLLER_GEN) rbac:roleName=manager-role crd webhook paths="./..." output:crd:artifacts:config=config/crd/bases

.PHONY: generate
generate: controller-gen ## Generate code containing DeepCopy, DeepCopyInto, and DeepCopyObject method implementations.
	$(CONTROLLER_GEN) object:headerFile="hack/boilerplate.go.txt" paths="./..."

## Location to install dependencies to
LOCALBIN ?= $(shell pwd)/bin
$(LOCALBIN):
	mkdir -p $(LOCALBIN)

## Tool Binaries
CONTROLLER_GEN ?= $(LOCALBIN)/controller-gen
CONTROLLER_TOOLS_VERSION ?= v0.11.1

.PHONY: controller-gen
controller-gen: $(CONTROLLER_GEN) ## Download controller-gen locally if necessary. If wrong version is installed, it will be overwritten.
$(CONTROLLER_GEN): $(LOCALBIN)
	test -s $(LOCALBIN)/controller-gen && $(LOCALBIN)/controller-gen --version | grep -q $(CONTROLLER_TOOLS_VERSION) || \
	GOBIN=$(LOCALBIN) go install sigs.k8s.io/controller-tools/cmd/controller-gen@$(CONTROLLER_TOOLS_VERSION)