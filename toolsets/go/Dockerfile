# Adds the Go toolchain to one of Kagent's harness images, so an agent can
# build and test Go code. HARNESS is the harness image, pinned by digest; the
# toolsets workflow passes it.
ARG HARNESS
FROM golang:1.27-alpine@sha256:f92b6ef800e499660581efdabdf25d9d817a9d124eaf900924f0504e7e27e12d AS go

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
# Let git over HTTPS authenticate through Substrate's egress gateway, on any
# host: after a 401, git sends the helper's placeholder, which the gateway
# replaces with the credential the agent's auth chooses for that host.
COPY --chmod=755 git-credential-egress /usr/local/bin/git-credential-egress
RUN git config --system credential.helper egress
# Back to the harness's own user; the entrypoint stays the harness's.
USER 65532:65532
