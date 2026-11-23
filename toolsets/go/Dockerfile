# Adds the Go toolchain to one of Kagent's harness images, so an agent can
# build and test Go code. HARNESS is the harness image, pinned by digest; the
# toolsets workflow passes it.
ARG HARNESS
FROM golang:1.27-alpine@sha256:70bd7ddea8b6cbe5d6943cde279d6b67feb10cc924fb9100d1d8757c1575b280 AS go

FROM ${HARNESS}
USER root
COPY --from=go /usr/local/go /usr/local/go
# build-base gives cgo a C compiler, which `go test -race` needs.
RUN apk add --no-cache build-base curl make
# GOTOOLCHAIN=local keeps go from downloading another toolchain when a go.mod
# asks for a newer one: the sandbox may not reach go.dev.
ENV PATH=/usr/local/go/bin:/data/go/bin:$PATH \
    GOPATH=/data/go \
    GOTOOLCHAIN=local
# Claude Code's managed settings, which apply even though Kagent's Claude
# harness starts it with --setting-sources "": no attribution in commits and
# pull requests. Codex ignores the file.
COPY claude-managed-settings.json /etc/claude-code/managed-settings.json
# Back to the harness's own user; the entrypoint stays the harness's.
USER 65532:65532
