APP?=app
APP_PORT?=9000

clean:
	rm -f ${APP}
build: clean
	go build -o ${APP}
run: build
	APP_PORT=${APP_PORT} ./${APP}
test:
	go test -v -race ./...