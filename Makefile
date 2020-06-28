# Makefile
VERSION ?= 0.0.0
NAME=terraform-provider-pingaccess_v${VERSION}
BASE_DOCKER_TAG=pingidentity/pingaccess:6.0.2-edge

pa-init:
	@docker run --rm -d --hostname pingaccess --name pingaccess -e PING_IDENTITY_DEVOPS_KEY=$(PING_IDENTITY_DEVOPS_KEY) -e PING_IDENTITY_DEVOPS_USER=$(PING_IDENTITY_DEVOPS_USER) -e PING_IDENTITY_ACCEPT_EULA=YES --publish 9000:9000 ${BASE_DOCKER_TAG}

test:
	@rm -f pingaccess/terraform.log
	@TF_LOG=TRACE TF_LOG_PATH=./terraform.log TF_ACC=1 go test -mod=vendor ./... -v -trimpath -coverprofile=coverage.out && go tool cover -func=coverage.out

unit-test:
	@go test -mod=vendor ./... -v -trimpath

test-and-report:
	TF_LOG=TRACE TF_LOG_PATH=./terraform.log TF_ACC=1 go test -mod=vendor ./... -v -trimpath -coverprofile=coverage.out -json | tee report.json

build:
	@go build -mod=vendor -o ${NAME} -trimpath .

deploy-local:
	@mkdir -p ~/.terraform.d/plugins
	@cp ${NAME} ~/.terraform.d/plugins/

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
