# Makefile
VERSION ?= local
NAME=terraform-provider-pingaccess_v${VERSION}

sweep:
	@TF_ACC=1 go test ./... -v -sweep=true

pa-init:
	@docker run --rm -d --name pingaccess --publish 9000:9000 pingidentity/pingaccess:5.2.2-edge

test:
	@rm -f pingaccess/terraform.log
	@TF_LOG=TRACE TF_LOG_PATH=./terraform.log TF_ACC=1 go test -mod=vendor ./... -v

test-and-report:
	@rm -f pingaccess/terraform.log coverage.out report.json
	@TF_LOG=TRACE TF_LOG_PATH=./terraform.log TF_ACC=1 go test -mod=vendor ./... -v -coverprofile=coverage.out -json > report.json && go tool cover -func=coverage.out

build:
	@go build -mod=vendor -o ${NAME} -gcflags "all=-trimpath=$GOPATH" .

release:
	@rm -rf build/*
	GOOS=darwin GOARCH=amd64 go build -o -mod=vendor build/darwin_amd64/${NAME} -gcflags "all=-trimpath=$GOPATH" . && zip -j build/darwin_amd64.zip build/darwin_amd64/${NAME}
	# GOOS=freebsd GOARCH=386 go build -o -mod=vendor build/freebsd_386/${NAME} -gcflags "all=-trimpath=$GOPATH" .
	# GOOS=freebsd GOARCH=amd64 go build -o -mod=vendor build/freebsd_amd64/${NAME} -gcflags "all=-trimpath=$GOPATH" .
	# GOOS=freebsd GOARCH=arm go build -o -mod=vendor build/freebsd_arm/${NAME} -gcflags "all=-trimpath=$GOPATH" .
	# GOOS=linux GOARCH=386 go build -o -mod=vendor build/linux_386/${NAME} -gcflags "all=-trimpath=$GOPATH" .
	GOOS=linux GOARCH=amd64 go build -o -mod=vendor build/linux_amd64/${NAME} -gcflags "all=-trimpath=$GOPATH" . && zip -j build/linux_amd64.zip build/linux_amd64/${NAME}
	# GOOS=linux GOARCH=arm go build -o -mod=vendor build/linux_arm/${NAME} -gcflags "all=-trimpath=$GOPATH" .
	# GOOS=openbsd GOARCH=386 go build -o -mod=vendor build/openbsd_386/${NAME} -gcflags "all=-trimpath=$GOPATH" .
	# GOOS=openbsd GOARCH=amd64 go build -o -mod=vendor build/openbsd_amd64/${NAME} -gcflags "all=-trimpath=$GOPATH" .
	# GOOS=solaris GOARCH=amd64 go build -o -mod=vendor build/solaris_amd64/${NAME} -gcflags "all=-trimpath=$GOPATH" .
	# GOOS=windows GOARCH=386 go build -o -mod=vendor build/windows_386/${NAME} -gcflags "all=-trimpath=$GOPATH" .
	GOOS=windows GOARCH=amd64 go build -o -mod=vendor build/windows_amd64/${NAME} -gcflags "all=-trimpath=$GOPATH" . && zip -j build/windows_amd64.zip build/windows_amd64/${NAME}
	
deploy-local:
	@mkdir -p ~/.terraform.d/plugins
	@cp ${NAME} ~/.terraform.d/plugins/

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
