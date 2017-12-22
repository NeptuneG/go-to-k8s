APP?=app
APP_PORT?=9000

PROJECT?=github.com/NeptuneG/go-to-k8s

RELEASE?=0.0.1
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIMESTAMP?=$(shell date -u '+%Y%m%d-%H%M%S')

clean:
	rm -f ${APP}
build: clean
	go build \
		-ldflags "-s -w \
		-X ${PROJECT}/version.Release=${RELEASE} \
		-X ${PROJECT}/version.Commit=${COMMIT} \
		-X ${PROJECT}/version.BuildTimestamp=${BUILD_TIMESTAMP}" \
		-o ${APP}
run: build
	APP_PORT=${APP_PORT} ./${APP}
test:
	go test -v -race ./...