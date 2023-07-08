SERVICE_NAME=asynqmon-multi-service
VERSION?= $(shell git describe --match 'v[0-9]*' --tags --always)
DOCKER_IMAGE_NAME=krobus00/${SERVICE_NAME}
CONFIG?=./config.yml
NAMESPACE?=default
PACKAGE_NAME=github.com/krobus00/${SERVICE_NAME}

build_args=-ldflags "-s -w -X $(PACKAGE_NAME)/internal/config.serviceVersion=$(VERSION) -X $(PACKAGE_NAME)/internal/config.serviceName=$(SERVICE_NAME)" -o ./bin/$(SERVICE_NAME) main.go
launch_args=
test_args=-coverprofile cover.out && go tool cover -func cover.out
cover_args=-cover -coverprofile=cover.out `go list ./...` && go tool cover -html=cover.out
air_args=--log.main_only=true --build.send_interrupt=true --build.rerun=true --build.delay '500'

# make tidy
tidy:
	go mod tidy

# make clean-up-mock
clean-up-mock:
	rm -rf ./internal/model/mock


# make generate
generate: clean-up-mock
	go generate ./...


# make lint
lint:
	@golangci-lint run

# make run dev server
# make run server
run:
ifeq (dev server, $(filter dev server,$(MAKECMDGOALS)))
	$(eval launch_args=server $(launch_args))
	air $(air_args) --build.cmd 'go build $(build_args)' --build.bin "./bin/$(SERVICE_NAME) $(launch_args)"
else ifeq (server, $(filter server,$(MAKECMDGOALS)))
	$(eval launch_args=server $(launch_args))
	$(shell if test -s ./bin/$(SERVICE_NAME); then ./bin/$(SERVICE_NAME) $(launch_args); else echo binary not found; fi)
endif

# make build
build:
	# build binary file
	go build -ldflags "-s -w -X $(PACKAGE_NAME)/internal/config.serviceVersion=$(VERSION) -X $(PACKAGE_NAME)/internal/config.serviceName=$(SERVICE_NAME)" -o ./bin/$(SERVICE_NAME) ./main.go
ifeq (, $(shell which upx))
	$(warning "upx not installed")
else
	# compress binary file if upx command exist
	upx --best --lzma ./bin/$(SERVICE_NAME)
endif

# make image VERSION="vx.x.x"
image:
	docker build -t ${DOCKER_IMAGE_NAME}:${VERSION} . -f ./deployments/Dockerfile

# make push-image VERSION="vx.x.x"
push-image:
	docker push ${DOCKER_IMAGE_NAME}:${VERSION}

# make docker-build-push VERSION="vx.x.x"
docker-build-push: image push-image

# make deploy VERSION="vx.x.x"
# make deploy VERSION="vx.x.x" NAMESPACE="staging"
# make deploy VERSION="vx.x.x" NAMESPACE="staging" CONFIG="./config-staging.yml"
deploy:
	@helm upgrade --install ${SERVICE_NAME} ./deployments/helm/${SERVICE_NAME} --set-file appConfig="${CONFIG}" --set app.container.version="${VERSION}" -n ${NAMESPACE}

# make coverage
coverage:
	@echo "total code coverage : "
	@go tool cover -func cover.out | grep total | awk '{print substr($$3, 1, length($$3)-1)}'

# make test
test:
ifeq (, $(shell which richgo))
	go test ./... $(test_args)
else
	richgo test ./... $(test_args)
endif

# make cover
cover:
ifeq (, $(shell which richgo))
	go test $(cover_args)
else
	richgo test $(cover_args)
endif

# make changelog VERSION=vx.x.x
changelog: tidy generate
	git-chglog -o CHANGELOG.md --next-tag $(VERSION)

%:
	@: