FROM golang:1.25.1 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./

RUN \
CGO_ENABLED=0 GOOS=linux \
go build -ldflags="-s -w" -o maki.cafe && \
strip maki.cafe

# create final image

FROM alpine:edge

WORKDIR /app

RUN apk add --no-cache -X http://dl-cdn.alpinelinux.org/alpine/edge/testing \
dart-sass

# COPY --from=build /etc/ssl/certs/ca-certificates.crt \
# /etc/ssl/certs/ca-certificates.crt

COPY --from=build /app/maki.cafe /app/maki.cafe

ENTRYPOINT ["/app/maki.cafe"]
