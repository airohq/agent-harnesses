# One of Kagent's harness images with only the settings every toolset adds.
# HARNESS is the harness image, pinned by digest; the toolsets workflow
# passes it.
ARG HARNESS
FROM ${HARNESS}
# Claude Code's managed settings, which apply even though Kagent's Claude
# harness starts it with --setting-sources "": no attribution in commits and
# pull requests. Codex ignores the file.
COPY claude-managed-settings.json /etc/claude-code/managed-settings.json
# Let git over HTTPS authenticate through Substrate's egress gateway, on any
# host: after a 401, git sends the helper's placeholder, which the gateway
# replaces with the credential the agent's auth chooses for that host.
USER root
COPY --chmod=755 git-credential-egress /usr/local/bin/git-credential-egress
RUN git config --system credential.helper egress
USER 65532:65532
