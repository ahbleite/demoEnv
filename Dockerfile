FROM registry.hub.docker.com/library/golang:1.10-alpine3.7
RUN apk update && apk add --no-cache git
WORKDIR /source
ADD . /source
RUN go get -d -v
RUN cd /source && go build -o /go/bin/main
EXPOSE 9999
ENTRYPOINT ["/go/bin/main"]