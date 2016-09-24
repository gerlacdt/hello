FROM golang:1.7.1

ADD . /go/src/github.com/gerlacdt/hello

RUN go install github.com/gerlacdt/hello

ENTRYPOINT /go/bin/hello

EXPOSE 3000
