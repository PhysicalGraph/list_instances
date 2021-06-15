# builder image
FROM golang:1.13-alpine3.11 as builder
RuN mkdir -p $GOPATH/src/github.com/PhysicalGraph/list_instances/
RUN set -ex \
	&& apk add git
COPY * $GOPATH/src/github.com/PhysicalGraph/list_instances/
WORKDIR $GOPATH/src/github.com/PhysicalGraph/list_instances/
RUN go get
RUN CGO_ENABLED=0 GOOS=linux go build -a -o list_instances .

# generate clean, final image for end users
FROM alpine:3.11.3
COPY --from=builder /go/src/github.com/PhysicalGraph/list_instances/list_instances .
RUN set -ex \
	&& apk add --no-cache ca-certificates

CMD /list_instances
