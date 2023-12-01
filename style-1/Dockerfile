FROM alpine:latest AS timezone_build
RUN apk --no-cache add tzdata ca-certificates  


FROM golang:1.21.4-alpine3.18 AS builder

RUN apk --no-cache add tzdata ca-certificates

ADD . /go/server

WORKDIR /go/server

RUN mkdir deploy
RUN go clean --modcache
RUN go mod tidy

RUN CGO_ENABLED=0 go build -o go_app adapter/cli/main.go 
RUN mv go_app ./deploy/go_app

FROM scratch AS production

COPY --from=timezone_build /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=timezone_build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder /go/server/deploy /server/

WORKDIR /server

ENTRYPOINT  ["./go_app", "serve"]