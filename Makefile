PHONY=build run docker-build docker-run docker-push

build:
	go build

run:
	NOMAD_PORT_http=8080 ./hello

docker-build:
	docker build . -t gerlacdt/helloapp

docker-run:
	docker run -it  -e NOMAD_PORT_http=8080 -p 8080:8080 gerlacdt/helloapp:latest

docker-push:
	docker push gerlacdt/helloapp:latest
