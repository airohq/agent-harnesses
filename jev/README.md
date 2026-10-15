# Jev harness

Runs TypeSafe's [Jev](https://docs.typesafe.ai), a decision model, as a
Kagent agent. Each A2A message is a
[System One request](https://docs.typesafe.ai/api), and the reply is System
One's response, unchanged:

```json
{"state": {"title": "App crashes when I press Save"}, "questions": {"kind": {"type": "choice", "criteria": {"bug": "A defect.", "feature": "A request."}}}}
```

```json
{"model": "jev-1.13.0", "answers": {"kind": {"type": "choice", "choice": "bug", "confidence": 1.0, "probabilities": {"bug": 1.0, "feature": 0.0}}}, "usage": {"input_tokens": 337, "output_tokens": 38}}
```

The harness adds two things to the request:

- `model`, from the agent's ModelConfig, unless the request sets one.
- `instructions`, the agent's system prompt, for each question without its
  own. So an agent's system prompt can hold the question, and messages carry
  only the state and the answers.

A message that isn't a System One request, or an error from the API, fails
the task with the reason. Rate-limited (429) and overloaded (529) calls are
retried with backoff. The harness keeps no conversation: each message is
evaluated on its own.

## Kagent setup

An `OpenAI` ModelConfig with TypeSafe's base URL gives the harness its model,
and makes Kagent's egress gateway add the API key to calls to that host, so
the sandbox only holds a placeholder:

```yaml
apiVersion: api.kagent.dev/v1alpha3
kind: ModelConfig
metadata:
  name: jev
spec:
  provider: OpenAI
  model: jev-latest
  apiKeySecret: typesafe-jev        # a Secret with the key
  apiKeySecretKey: TYPESAFE_API_KEY
  openAI:
    baseUrl: https://api.typesafe.ai
---
apiVersion: api.kagent.dev/v1alpha3
kind: Harness
metadata:
  name: jev
spec:
  byo: {}
  workload:
    image: ghcr.io/airohq/agent-harnesses/jev@sha256:<digest>
    command: ["/jev-harness"]
  substrate:
    workerPoolRef: {name: kagent-default}
    snapshotPolicy: {location: <snapshot location>}
```

An AgentTemplate that uses the harness names the ModelConfig with
`modelConfig`, and its `systemPrompt` is the default question.

## Develop

The module pins Kagent to v1.0.0-alpha13 and needs Go 1.27.

```sh
go test ./...                                     # with a fake API
TYPESAFE_API_KEY=... go test -run TestLive ./...  # against System One
```

Outside Kagent, `TYPESAFE_API_KEY`, `JEV_BASE_URL` and `JEV_MODEL` set the key,
URL and model.
