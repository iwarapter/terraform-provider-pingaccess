# Makefile

test:
	@go test ./... -v -coverprofile=coverage.out && go tool cover -func=coverage.out

build:
	@go build  -o terraform-provider-pingaccess .

deploy-local:
	@mkdir -p ~/.terraform.d/plugins
	@cp terraform-provider-pingaccess ~/.terraform.d/plugins/

func-init:
	@rm -rf func-tests/.terraform
	@rm -rf func-tests/crash.log
	@rm -rf func-tests/run.log
	@cd func-tests && terraform init

func-plan:
	@cd func-tests && terraform plan

func-apply:
	@cd func-tests && terraform apply -auto-approve