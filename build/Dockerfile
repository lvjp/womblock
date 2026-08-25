FROM golang:1.26.3-trixie@sha256:f34e7161a14638b812ce491bd89c81718f309cac6ec0ffe016e5fbcb4bdc8c06 AS build

WORKDIR /opt/wtf-go/src

RUN go env
RUN [ $(go env GOCACHE) = /root/.cache/go-build ]
RUN [ $(go env GOMODCACHE) = /go/pkg/mod ]

ARG COVER
RUN --mount=type=bind,target=.,readonly \
    --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg/mod \
    go mod download -x && CGO_ENABLED=0 go build ${COVER:+-cover} -o /opt/wtf-go/wtf-go .

FROM debian:13.6-slim@sha256:d7e12182ce18b85b93007c1dedf31f2d29e01ccf3182cc4017c709b6259bc132

COPY --from=build /opt/wtf-go/wtf-go /usr/bin
ENTRYPOINT ["/usr/bin/wtf-go"]
CMD ["serve"]
