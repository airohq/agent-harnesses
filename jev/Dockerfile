FROM --platform=$BUILDPLATFORM golang:1.27 AS build
ARG TARGETOS
ARG TARGETARCH
WORKDIR /src
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download
COPY *.go ./
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -trimpath -ldflags='-s -w' -o /jev-harness . \
    && mkdir /data

# Kagent's Claude harness also listens on port 80 as a non-root user.
FROM gcr.io/distroless/static:nonroot
COPY --from=build /jev-harness /jev-harness
# Kagent mounts the durable /data volume here and keeps the bootstrap bundle
# in it. Kagent's own runtime images give /data to the runtime user, too.
COPY --from=build --chown=65532:65532 /data /data
EXPOSE 80 8081
ENTRYPOINT ["/jev-harness"]
