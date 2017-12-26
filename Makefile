APP?=app
APP_PORT?=9000
GOOS?=linux
GOARCH=amd64

CONTAINER_IMAGE?=docker.io/neptuneg/${APP}
PROJECT?=github.com/NeptuneG/go-to-k8s

RELEASE?=0.0.1
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIMESTAMP?=$(shell date -u '+%Y%m%d-%H%M%S')

clean:
	rm -f ${APP}
build: clean
	CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} go build \
		-ldflags "-s -w \
		-X ${PROJECT}/version.Release=${RELEASE} \
		-X ${PROJECT}/version.Commit=${COMMIT} \
		-X ${PROJECT}/version.BuildTimestamp=${BUILD_TIMESTAMP}" \
		-o ${APP}
container: build
	docker build -t ${CONTAINER_IMAGE}:${RELEASE} .
run: container
	docker stop ${APP}:${RELEASE} || true && docker rm ${APP}:${RELEASE} || true
	docker run --name ${APP} -p ${APP_PORT}:${APP_PORT} --rm \
		-e "APP_PORT=${APP_PORT}" ${APP}:${RELEASE}
	APP_PORT=${APP_PORT} ./${APP}
push: container
	docker push ${CONTAINER_IMAGE}:${RELEASE}
test:
	go test -v -race ./...
