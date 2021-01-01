FROM golang:1.15-alpine
ADD gondx-arm /gondx-arm
ENTRYPOINT [ "/gondx-arm"]