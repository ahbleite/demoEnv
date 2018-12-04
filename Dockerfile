FROM golang:1.10-alpine3.7 as build-env
RUN apk update && apk add --no-cache git
WORKDIR /source
ADD . /source
RUN go get -d -v
RUN cd /source && go build -o /go/bin/main

FROM nginx:1.15.7-alpine
COPY --from=build-env /go/bin/main /go/bin/main
EXPOSE 9999
ENTRYPOINT ["/go/bin/main"]