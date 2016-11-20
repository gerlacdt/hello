FROM golang:1.7.1

# copy project-src in GOPATH of golang
# docker image, so it can be build
ADD . /go/src/github.com/gerlacdt/hello
WORKDIR /go/src/github.com/gerlacdt/hello

# build binary
RUN go get && go build -o main

EXPOSE 8080

CMD ["./main"]
