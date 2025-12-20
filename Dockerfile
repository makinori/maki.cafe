# compile site

FROM docker.io/golang:alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY --exclude=.git ./ ./
COPY .git/refs/heads/main ./HEAD

# ARG GIT_COMMIT # cant run git in quadlet build

RUN \
GOEXPERIMENT=greenteagc \
CGO_ENABLED=0 GOOS=linux \
go build -ldflags="-s -w \
-X 'maki.cafe/src/config.GitCommit=$(cat HEAD | head -c 8)'\
" -o maki.cafe

# create image

FROM ghcr.io/makinori/foxlib:base

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=build /app/maki.cafe /

CMD ["/maki.cafe"]
