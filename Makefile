
# Image URL to use all building/pushing image targets
IMG ?= raushan2016/etcdexplore

# Build manager binary
manager:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o bin/etcdexplore lucas.go

# Run against the configured Kubernetes cluster in ~/.kube/config
run:
	go run ./main.go

docker-all: docker-build docker-push

# Build the docker image
docker-build:
	docker build . -t ${IMG}

# Push the docker image
docker-push:
	docker push ${IMG}
