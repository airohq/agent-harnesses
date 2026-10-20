# One of Kagent's harness images with only the settings every toolset adds.
# HARNESS is the harness image, pinned by digest; the toolsets workflow
# passes it.
ARG HARNESS
FROM ${HARNESS}
# Claude Code's managed settings, which apply even though Kagent's Claude
# harness starts it with --setting-sources "": no attribution in commits and
# pull requests. Codex ignores the file.
COPY claude-managed-settings.json /etc/claude-code/managed-settings.json
