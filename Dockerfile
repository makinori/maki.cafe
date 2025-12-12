# compile sass

ARG SASS_VERSION=1.96.0

FROM docker.io/bufbuild/buf:latest AS buf
FROM ghcr.io/dart-musl/dart:latest AS dart

COPY --from=buf /usr/local/bin/buf /usr/local/bin/

RUN \
git clone https://github.com/sass/dart-sass.git /dart-sass && \
cd /dart-sass && \
git checkout ${SASS_VERSION} && \
dart pub get && \
dart run grinder protobuf && \
dart compile exe bin/sass.dart -o sass

# compile site

FROM docker.io/golang:alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./

RUN \
GOEXPERIMENT=greenteagc \
CGO_ENABLED=0 GOOS=linux \
go build -ldflags="-s -w" -o maki.cafe

# create image

FROM scratch

WORKDIR /
ENV PATH=/:$PATH

# should only be one file
COPY --from=dart /lib/ld-musl-*.so.1 /lib/
COPY --from=dart /dart-sass/sass /sass

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=build /app/maki.cafe /maki.cafe

CMD ["/maki.cafe"]
