package main

import (
	"encoding/json"
	"testing"

	"github.com/kagent-dev/kagent/go/api/adk"
)

// The harness reads Kagent's config with its own struct, so this checks it
// against the JSON Kagent writes for a BYO harness.
func TestConfigReadsKagentBYOConfig(t *testing.T) {
	data, err := json.Marshal(adk.AgentConfig{
		Instruction: "Classify this issue.",
		Runtime:     &adk.RuntimeConfig{ControlPlaneURL: "http://kagent-controller.kagent.svc:8083"},
	})
	if err != nil {
		t.Fatal(err)
	}
	var cfg config
	if err := json.Unmarshal(data, &cfg); err != nil {
		t.Fatal(err)
	}
	if cfg.Instruction != "Classify this issue." {
		t.Errorf("instruction = %q", cfg.Instruction)
	}
	if cfg.Runtime.ControlPlaneURL != "http://kagent-controller.kagent.svc:8083" {
		t.Errorf("control plane URL = %q", cfg.Runtime.ControlPlaneURL)
	}
}
