# FROM golang:1.15.6-buster as BUILD-ENV
FROM golang:1.15-alpine AS build-env
ARG VERSION
RUN apk add build-base linux-headers

ENV BUILD_HOME=/BUILD

RUN mkdir -p ${BUILD_HOME}

ADD . ${BUILD_HOME}
WORKDIR ${BUILD_HOME}

RUN go mod download

# Build temporal binary
RUN go build -o /bin/temporal \
    -ldflags "-X main.Version=$VERSION" \
    ./cmd/gondx

RUN go build -o /bin/gondx ./cmd/gondx

COPY entrypoint.sh /bin/entrypoint.sh

ENTRYPOINT [ "/bin/entrypoint.sh"]