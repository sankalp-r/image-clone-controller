# Go parameters
SHELL:=/bin/bash
DOCKER_REPO=sankalprangare
DOCKER_IMG=$(DOCKER_REPO)/image-clone-controller
DOCKER_IMG_TAG=latest

# Build docker image and push to repo
build-docker:
	docker build -t $(DOCKER_IMG):$(DOCKER_IMG_TAG) ../../ -f ../../cmd/imgcc/main/Dockerfile
	docker push $(DOCKER_IMG):$(DOCKER_IMG_TAG)


# deploy controller to kubernetes-cluster
deploy:
	cat deployment.yaml | envsubst | kubectl apply -f -

# remove controller from kubernetes-cluster
clean:
	cat deployment.yaml | envsubst | kubectl delete -f -