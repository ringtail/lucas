
# Image URL to use all building/pushing image targets
IMG ?= registry.cn-hangzhou.aliyuncs.com/ringtail/lucas
# ETCD servers
ENDPOINTS ?= 0.0.0.0:2379
# FOLDER containing the certifcates 
CERTFOLDER ?= /etc/kubernetes/pki/etcd/

# Build binary
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o bin/lucas lucas.go

docker-all: docker-build docker-push

# Build the docker image
docker-build:
	docker build . -t ${IMG}

# Push the docker image
docker-push:
	docker push ${IMG}

# Run the container 
docker-run:
	docker run -d -p 8088:8080  -v ${CERTFOLDER}:/etc/kubernetes/pki/etcd/ -e CA_FILE=/etc/kubernetes/pki/etcd/ca.pem -e CERT_FILE=/etc/kubernetes/pki/etcd/etcd-client.pem -e KEY_FILE=/etc/kubernetes/pki/etcd/etcd-client-key.pem -e ENDPOINTS=${ENDPOINTS} ${IMG}
