BUILD_NUMBER := latest
PROJECT_NAME := stupid-httpd-linux-amd64
DOCKER_REGISTRY := jodydadescott
DOCKER_IMAGE_NAME?=$(PROJECT_NAME)
DOCKER_IMAGE_TAG?=$(BUILD_NUMBER)

build:
	cd ../../build && $(MAKE) linux-amd64
	cp ../../stupid-httpd-linux-amd64 stupid-httpd
	docker build -t $(DOCKER_REGISTRY)/$(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG) .
	touch build

push:
	@ $(MAKE) build
	docker push $(DOCKER_REGISTRY)/$(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG)

clean:
	$(RM) -f build stupid-httpd
