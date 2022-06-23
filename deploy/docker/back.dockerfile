ARG GO_VERSION=1.18.3-buster
ARG DEBIAN_VERSION=11.2-slim

FROM docker.io/golang:${GO_VERSION} as builder

COPY . /build
WORKDIR /build

RUN mkdir -p ./bin/
RUN go build -o ./bin/ ./cmd/chat-server

FROM docker.io/debian:${DEBIAN_VERSION}

COPY --from=builder /build/bin/chat-server /

ENTRYPOINT [ "/chat-server" ]
CMD []