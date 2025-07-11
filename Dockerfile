FROM golang:1.24.4-alpine3.22 AS builder

WORKDIR /go/src/app
COPY . .
RUN env CGO_ENABLED=0 go build -o /nyx

FROM scratch AS build-image
COPY --from=builder /nyx /nyx
COPY nyx.yaml /etc/nyx/nyx.yaml
ENTRYPOINT ["/nyx"]
CMD ["start"]