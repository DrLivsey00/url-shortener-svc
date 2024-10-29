FROM golang:1.23-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/github.com/DrLivsey00/url-shortener-svc
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/url-shortener-svc /go/src/github.com/DrLivsey00/url-shortener-svc


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/url-shortener-svc /usr/local/bin/url-shortener-svc
COPY config.yaml /usr/local/bin/config/config.yaml
COPY nginx.conf /usr/local/bin/config/nginx.conf
COPY entry.sh /usr/local/bin/entrypoint.sh

RUN chmod +x /usr/local/bin/entrypoint.sh
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]



