# agent-harnesses

Kagent harnesses for models that Kagent doesn't run natively. Each harness is
a thin adapter, like Kagent's own Claude and Codex harnesses: it runs on
Kagent's `byo` runtime, takes a message in the model's own format, calls the
model, and replies with the model's response. Callers decide what to ask and
how to read the answer.

| Harness | Model | Image |
| --- | --- | --- |
| [jev](jev/) | TypeSafe's [Jev](https://docs.typesafe.ai), a decision model | `ghcr.io/airohq/agent-harnesses/jev` |

Each harness is its own Go module with its own CI workflow. A merge to `main`
publishes its image, tagged `sha-<commit>`; the run summary shows the digest.
Kagent runs harness images by digest, so a deployment pins the digest.

The harnesses hold no credentials. Kagent's egress gateway adds a model's API
key to the harness's calls, from the agent's ModelConfig.
