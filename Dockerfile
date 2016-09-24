FROM golang:1.7.1

# copy project-src in GOPATH of golang
# docker image, so it can be build
ADD . /go/src/github.com/gerlacdt/hello
WORKDIR /go/src/github.com/gerlacdt/hello

# build binary
RUN go build -o main

EXPOSE 3000

CMD ["./main"]
