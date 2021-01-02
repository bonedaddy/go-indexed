# FROM golang:1.15.6-buster as BUILD-ENV
FROM golang:1.15-alpine AS build-env
ARG VERSION
RUN apk add build-base linux-headers

ENV BUILD_HOME=/BUILD

RUN mkdir -p ${BUILD_HOME}

ADD . ${BUILD_HOME}
WORKDIR ${BUILD_HOME}

RUN go mod download

# Build gondx binary
RUN go build -o /bin/gondx \
    -ldflags "-X main.Version=$VERSION" \
    ./cmd/gondx

FROM alpine

COPY --from=build-env /bin/gondx /bin/gondx

COPY entrypoint.sh /bin/entrypoint.sh

ENTRYPOINT [ "/bin/entrypoint.sh"]