# agent-harnesses

Kagent harnesses and harness images that Kagent doesn't ship itself.

## Harnesses

Each harness is a thin adapter for a model that Kagent doesn't run natively,
like Kagent's own Claude and Codex harnesses: it runs on Kagent's `byo`
runtime, takes a message in the model's own format, calls the model, and
replies with the model's response. Callers decide what to ask and how to read
the answer.

| Harness | Model | Image |
| --- | --- | --- |
| [jev](jev/) | TypeSafe's [Jev](https://docs.typesafe.ai), a decision model | `ghcr.io/airohq/agent-harnesses/jev` |

Each harness is its own Go module with its own CI workflow.

The harnesses hold no credentials. Kagent's egress gateway adds a model's API
key to the harness's calls, from the agent's ModelConfig.

## Toolsets

Kagent's harness images carry only what the harness needs, such as `git`. A
toolset adds the tools an agent needs to build and test code, on top of a
Kagent harness image. The image keeps the harness's entrypoint and user, so
it replaces `workload.image` in a Harness of the same kind (`claude: {}` or
`codex: {}`).

| Toolset | Adds | Images |
| --- | --- | --- |
| [go](toolsets/go/) | Go, a C compiler for cgo and `-race`, `make`, `curl` | `ghcr.io/airohq/agent-harnesses/claude-go`, `ghcr.io/airohq/agent-harnesses/codex-go` |

The [toolsets workflow](.github/workflows/toolsets.yml) lists the images: a
name, a toolset and the Kagent harness image it builds on, pinned by digest.
When Kagent upgrades, update those digests. A new combination is a new entry
there; a new toolset is a new directory under `toolsets/`.

## Publishing

A merge to `main` publishes the changed images, tagged `sha-<commit>`; the run
summary shows the digest. Kagent runs harness images by digest, so a
deployment pins the digest.
