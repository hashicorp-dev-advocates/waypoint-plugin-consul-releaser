PLUGIN_NAME=consul-release-controller
DOCKER_REGISTRY=nicholasjackson/waypoint-custom-odr
DOCKER_TAG=0.1.0

ifndef OS
	OS := $(shell ./print_os)
export OS
endif

ARCH := "amd64"

ifeq ($(PLATFORM),aarch64)
	ARCH = "arm64"
endif

.PHONY: all

test:
	echo "Testing... ${ARCH} ${OS}"

all: protos build

# Generate the Go code from Protocol Buffer definitions
protos:
	@echo ""
	@echo "Build Protos"
	@echo ""

	protoc -I . --go-grpc_out==plugins=grpc:. --go_out=paths=source_relative:. ./builder/output.proto

# Builds the plugin on your local machine
build:
	@echo ""
	@echo "Compile Plugin"
	@echo ""

	rm -rf ./bin

	GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -o ./bin/linux_arm64/waypoint-plugin-${PLUGIN_NAME} ./main.go 
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./bin/linux_amd64/waypoint-plugin-${PLUGIN_NAME} ./main.go 

# Install the plugin locally
install: build
	@echo ""
	@echo "Installing Plugin"
	@echo ""

	mkdir -p ${HOME}/.config/waypoint/plugins
	cp ./bin/${OS}_${ARCH}/waypoint-plugin-${PLUGIN_NAME}* ${HOME}/.config/waypoint/plugins/   

# Build the plugin using a Docker container
build-docker:
	rm -rf ./releases
	DOCKER_BUILDKIT=1 docker build --output releases --progress=plain .

build-odr-dev:
	DOCKER_BUILDKIT=1 docker build --progress=plain -f Dockerfile.odr -t ${DOCKER_REGISTRY}:0.1.0 .

build-odr-multi-arch:
	docker run --rm --privileged multiarch/qemu-user-static --reset -p yes
	docker buildx create --name multi || true
	docker buildx use multi
	docker buildx inspect --bootstrap
	docker buildx build --platform linux/arm64,linux/amd64 \
		-t ${DOCKER_REGISTRY}:${DOCKER_TAG} \
    -f ./Dockerfile.odr \
    .  \
		--push
	docker buildx rm multi

zip:
	zip -j ./bin/waypoint-plugin-${PLUGIN_NAME}_linux_arm64.zip ./bin/linux_arm64/waypoint-plugin-${PLUGIN_NAME}
	zip -j ./bin/waypoint-plugin-${PLUGIN_NAME}_linux_amd64.zip ./bin/linux_amd64/waypoint-plugin-${PLUGIN_NAME}

dev: server bootstrap local-project

server:
	waypoint server run -accept-tos 2>&1 > waypoint.log &

bootstrap:
	waypoint server bootstrap -server-addr=127.0.0.1:9701 -server-tls-skip-verify > waypoint.token

local-project:
	waypoint project apply test

build-app:
	cd example_app && waypoint build -push=false --plain

nuke:
	@-pkill -9 -f waypoint
	@-rm -rf .terraform
	@-rm -rf .waypoint waypoint.token waypoint.log waypoint-restore.db.lock data.db || true
