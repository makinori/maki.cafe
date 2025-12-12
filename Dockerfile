FROM docker.io/golang:1.25-alpine AS build

ARG SASS_VER=1.96.0
ARG SASS_FILE=dart-sass-${SASS_VER}-linux-x64-musl.tar.gz

WORKDIR /
ADD https://github.com/sass/dart-sass/releases/download/${SASS_VER}/${SASS_FILE} /
RUN tar xzvf ${SASS_FILE} && rm -f ${SASS_FILE}

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./

RUN \
GOEXPERIMENT=greenteagc \
CGO_ENABLED=0 GOOS=linux \
go build -ldflags="-s -w" -o maki.cafe

# create final image

FROM docker.io/alpine:latest

WORKDIR /app

# COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

ENV PATH=/dart-sass:$PATH
COPY --from=build /dart-sass /dart-sass

COPY --from=build /app/maki.cafe /app/maki.cafe

ENTRYPOINT ["/app/maki.cafe"]
