# One of Kagent's harness images with only the settings every toolset adds.
# HARNESS is the harness image, pinned by digest; the toolsets workflow
# passes it.
ARG HARNESS
FROM ${HARNESS}
# Claude Code's managed settings, which apply even though Kagent's Claude
# harness starts it with --setting-sources "": no attribution in commits and
# pull requests. Codex ignores the file.
COPY claude-managed-settings.json /etc/claude-code/managed-settings.json
# Let git over HTTPS to github.com authenticate through Substrate's egress
# gateway. The gateway only replaces a header the request already carries, so
# git sends this inert placeholder; when the agent opens the github network
# preset, the gateway swaps it for the repository's installation token. The
# real token never enters the sandbox (airohq/agentic-sdlc#81).
USER root
RUN git config --system http.https://github.com/.extraHeader \
      "Authorization: Basic eC1hY2Nlc3MtdG9rZW46cGxhY2Vob2xkZXI="
USER 65532:65532
