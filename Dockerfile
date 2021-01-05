FROM golang:1.15-alpine3.12 AS build-env
ARG VERSION
RUN apk add build-base linux-headers
ENV BUILD_HOME=/BUILD
RUN mkdir -p ${BUILD_HOME}
ADD . ${BUILD_HOME}
WORKDIR ${BUILD_HOME}
RUN go mod download
RUN go build -o /bin/gondx \
    -ldflags "-X main.Version=$VERSION" \
    ./cmd/gondx

FROM alpine:3.12
RUN apk add --no-cache tini
COPY --from=build-env /bin/gondx /bin/gondx
ENTRYPOINT  ["/bin/gondx", "--config", "/config.yml" ]
CMD [ "--startup.sleep", "discord", "ndx-bot", "--update.database" ]