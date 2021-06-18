# image-clone-controller
 Kubernetes image clone controller

## Description

* It copies public docker images of Deployments & DaemonSet into backup-registry.
* It also updates the Deployments/DaemonSets to use image from backup-registry.
* `Dockerfile` is included in `cmd/imgcc/main` directory.
* `deployment.yaml` in the `setup/makebuild` directory is using `docker.io/sankalprangare/image-clone-controller` as image, and it also contains `RBAC` manifests required by the controller.
* `Makefile` in the `setup/makebuild` directory contains commands for publishing docker image, deploying and uninstalling the controller.

## How to run
Note: This controller will get deployed in `image-clone-controller` namespace. 
* Navigate to `setup/makebuild` directory and export `REGISTRY`, `REGISTRY_USER` and `REGISTRY_PASSWORD` as env variables like example below:
    * export REGISTRY=docker.io/sankalprangare
    * export REGISTRY_USER=abc
    * export REGISTRY_PASSWORD=abc123
    
* After the previous step, run `make deploy` command from `Makefile`, this will deploy the controller. <br>
  Or you can also use `kubectl` directly, given that you replace env variables in the `deployment.yaml`.
  
* Controller will update the images of respective objects after sometime.
  
* To uninstall the controller run `make clean`.

## Demo
[![asciicast](https://asciinema.org/a/397344.svg)](https://asciinema.org/a/397344)

## Enhancement scope
* This simple implementation can be further improved to production-grade implementation.
* Enhancement can be made to support private-registry.   
* Unit-tests can be improved to cover more cases.
* Predicates can be added to improve filtering.



## References
* [Controller-runtime](https://github.com/kubernetes-sigs/controller-runtime)
* [go-containerregistry](https://github.com/google/go-containerregistry)


