# Go parameters
SHELL:=/bin/bash
DOCKER_REPO=sankalprangare
DOCKER_API_IMG=$(DOCKER_REPO)/image-clone-controller
DOCKER_API_IMG_TAG=latest

build:
	go build ../../cmd/api/main/main.go

# Run locally
run:
	go run ../../cmd/api/main/main.go

# Run unit tests
test:
	cd ../.. && go test ./...

# Build docker image and push to repo
build-docker:
	docker build -t $(DOCKER_API_IMG):$(DOCKER_API_IMG_TAG) ../../ -f ../../cmd/imgcc/main/Dockerfile
	docker push $(DOCKER_API_IMG):$(DOCKER_API_IMG_TAG)

# Run docker container
run-docker:
	docker run -p 10000:10000 $(DOCKER_API_IMG):$(DOCKER_API_IMG_TAG)

# deploy app to kubernetes-cluster
deploy:
	kubectl create -f deployment.yaml

# remove app from kubernetes-cluster
clean:
	kubectl delete -f deployment.yaml