FROM golang:1.15-alpine
ADD gondx /gondx
ENTRYPOINT [ "/gondx"]