# Makefile
VERSION ?= 0.0.1-ci
NAME=terraform-provider-pingaccess_v${VERSION}
PINGACCESS_VERSION=6.1.3-edge
BASE_DOCKER_TAG=pingidentity/pingaccess:${PINGACCESS_VERSION}
OS_NAME := $(shell uname -s | tr A-Z a-z)

pa-init:
	@docker run --rm -d --hostname pingaccess --name pingaccess -e PING_IDENTITY_DEVOPS_KEY=$(PING_IDENTITY_DEVOPS_KEY) -e PING_IDENTITY_DEVOPS_USER=$(PING_IDENTITY_DEVOPS_USER) -e PING_IDENTITY_ACCEPT_EULA=YES --publish 9000:9000 ${BASE_DOCKER_TAG}

checks:
	@go fmt ./...
	@staticcheck ./...
	@gosec ./...
	@goimports -w internal

sweep:
	@echo "WARNING: This will destroy infrastructure. Use only in development accounts."
	go test ./... -v -sweep=all -timeout 60m

test-proto:
	@TF_ACC=1 go test -mod=vendor ./internal/protocolprovider -v -trimpath

test-sdkv2:
	@TF_ACC=1 go test -mod=vendor ./internal/sdkv2provider -v -trimpath

test:
	@TF_ACC=1 go test -mod=vendor ./... -v -trimpath -coverprofile=coverage.out && go tool cover -func=coverage.out

unit-test:
	@go test -mod=vendor ./... -v -trimpath

test-and-report:
	@TF_ACC=1 go test -mod=vendor ./... -v -trimpath -coverprofile=coverage.out -json | tee report.json

build:
	@go build -mod=vendor -o ${NAME} -trimpath .

deploy-local:
	@mkdir -p ~/.terraform.d/plugins/registry.terraform.io/iwarapter/pingaccess/${VERSION}/${OS_NAME}_amd64
	@cp ${NAME} ~/.terraform.d/plugins/registry.terraform.io/iwarapter/pingaccess/${VERSION}/${OS_NAME}_amd64

func-init:
	@rm -rf func-tests/.terraform
	@rm -rf func-tests/crash.log
	@rm -rf func-tests/run.log
	@cd func-tests && terraform init

func-plan:
	@cd func-tests && TF_LOG=TRACE TF_LOG_PATH=./terraform.log terraform plan

func-apply:
	@cd func-tests && TF_LOG=TRACE TF_LOG_PATH=./terraform.log terraform apply -auto-approve

func-destroy:
	@cd func-tests && terraform destroy -auto-approve

.PHONY: test build deploy-local
