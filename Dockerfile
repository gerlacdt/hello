FROM golang:1.12.9-alpine3.10 as build

RUN apk add --no-cache git
WORKDIR /go/src/github.com/gerlacdt/hello
COPY . .
RUN go get -d -v ./...
RUN go build

# step 2
FROM alpine:3.10.2

COPY --from=build /go/src/github.com/gerlacdt/hello/hello /hello
CMD ["/hello"]
