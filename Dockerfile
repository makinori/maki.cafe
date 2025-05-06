FROM golang:1.24.2 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./

RUN \
CGO_ENABLED=0 GOOS=linux \
go build -ldflags="-s -w" -o maki.cafe && \
strip maki.cafe

FROM scratch

WORKDIR /app

COPY --from=build /etc/ssl/certs/ca-certificates.crt \
/etc/ssl/certs/ca-certificates.crt

COPY --from=build /app/maki.cafe /maki.cafe

ENTRYPOINT ["/maki.cafe"]