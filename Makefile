# Makefile

sweep:
	@TF_ACC=1 go test ./... -v -sweep=true

pa-init:
	@docker run --rm -d --name pingaccess --publish 9000:9000 pingidentity/pingaccess:5.2.2-edge

test:
	@rm -f pingaccess/terraform.log
	@TF_LOG=TRACE TF_LOG_PATH=./terraform.log TF_ACC=1 go test ./... -v

test-and-report:
	@rm -f pingaccess/terraform.log coverage.out report.json
	@TF_LOG=TRACE TF_LOG_PATH=./terraform.log TF_ACC=1 go test ./... -v -coverprofile=coverage.out -json > report.json && go tool cover -func=coverage.out

build:
	@go build -o terraform-provider-pingaccess .

release:
	@rm -rf build/*
	GOOS=darwin GOARCH=amd64 go build -o build/darwin_amd64/terraform-provider-pingaccess . && zip -j build/darwin_amd64.zip build/darwin_amd64/terraform-provider-pingaccess
	# GOOS=freebsd GOARCH=386 go build -o build/freebsd_386/terraform-provider-pingaccess .
	# GOOS=freebsd GOARCH=amd64 go build -o build/freebsd_amd64/terraform-provider-pingaccess .
	# GOOS=freebsd GOARCH=arm go build -o build/freebsd_arm/terraform-provider-pingaccess .
	# GOOS=linux GOARCH=386 go build -o build/linux_386/terraform-provider-pingaccess .
	GOOS=linux GOARCH=amd64 go build -o build/linux_amd64/terraform-provider-pingaccess . && zip -j build/linux_amd64.zip build/linux_amd64/terraform-provider-pingaccess
	# GOOS=linux GOARCH=arm go build -o build/linux_arm/terraform-provider-pingaccess .
	# GOOS=openbsd GOARCH=386 go build -o build/openbsd_386/terraform-provider-pingaccess .
	# GOOS=openbsd GOARCH=amd64 go build -o build/openbsd_amd64/terraform-provider-pingaccess .
	# GOOS=solaris GOARCH=amd64 go build -o build/solaris_amd64/terraform-provider-pingaccess .
	# GOOS=windows GOARCH=386 go build -o build/windows_386/terraform-provider-pingaccess .
	GOOS=windows GOARCH=amd64 go build -o build/windows_amd64/terraform-provider-pingaccess . && zip -j build/windows_amd64.zip build/windows_amd64/terraform-provider-pingaccess
	
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
	@cd func-tests && TF_LOG=TRACE TF_LOG_PATH=./terraform.log terraform apply -auto-approve

func-destroy:
	@cd func-tests && terraform destroy -auto-approve

func-cli-gen:
	@cd ../pingaccess-sdk-go-gen-cli/ && make generate

sonar:
	@sonar-scanner \
		-Dsonar.projectKey=github.com.iwarapter.terraform-provider-pingaccess \
		-Dsonar.sources=. \
		-Dsonar.go.coverage.reportPaths=coverage.out \
		-Dsonar.go.tests.reportPaths=report.json \
		-Dsonar.exclusions=vendor/**/* \
		-Dsonar.tests="." \
		-Dsonar.test.inclusions="**/*_test.go"
		
.PHONY: test build deploy-local
