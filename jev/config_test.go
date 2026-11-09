package main

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"google.golang.org/protobuf/proto"

	kagentconfig "github.com/kagent-dev/kagent/go/adk/pkg/config"
	contract "github.com/kagent-dev/kagent/go/api/bootstrap"
	configv1 "github.com/kagent-dev/kagent/go/api/gen/kagent/config/v1"
)

// Checks that the harness finds its settings in the bootstrap bundle Kagent
// builds for a BYO harness with an OpenAI ModelConfig, once the bundle is on
// the durable volume.
func TestLoadBundleReadsKagentBYOConfig(t *testing.T) {
	config, err := kagentconfig.Marshal(&configv1.ADKConfig{
		Runtime: &configv1.ADKRuntimeConfig{ControlPlaneUrl: "http://kagent-controller.kagent.svc.cluster.local:8083"},
		Agent: &configv1.ADKAgentConfig{
			Name:        "judge",
			Instruction: "Classify this issue.",
			Model: &configv1.ADKModelConfig{
				Model:    "jev-latest",
				Provider: &configv1.ADKModelConfig_Openai{Openai: &configv1.ADKOpenAI{BaseUrl: "http://llm-gateway/typesafe"}},
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	revision := "0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef"
	bundle, err := contract.NewBundle(revision, config, []byte(`{"name":"judge"}`))
	if err != nil {
		t.Fatal(err)
	}
	dir := t.TempDir()
	raw, err := proto.Marshal(bundle)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "bundle.pb"), raw, 0o600); err != nil {
		t.Fatal(err)
	}
	descriptor, err := json.Marshal(contract.Descriptor{
		FormatVersion:   contract.FormatVersion,
		ControlPlaneURL: "http://kagent-controller.kagent.svc.cluster.local:8083",
		Revision:        revision,
		SHA256:          bundle.Sha256,
	})
	if err != nil {
		t.Fatal(err)
	}
	getenv := func(key string) string {
		if key == contract.Environment {
			return string(descriptor)
		}
		return ""
	}

	cfg, card, err := loadBundle(context.Background(), getenv, dir)
	if err != nil {
		t.Fatal(err)
	}
	if card.Name != "judge" {
		t.Errorf("card name = %q", card.Name)
	}
	agent := cfg.GetAgent()
	if got := agent.GetInstruction(); got != "Classify this issue." {
		t.Errorf("instruction = %q", got)
	}
	if got := agent.GetModel().GetModel(); got != "jev-latest" {
		t.Errorf("model = %q", got)
	}
	if got := agent.GetModel().GetOpenai().GetBaseUrl(); got != "http://llm-gateway/typesafe" {
		t.Errorf("base URL = %q", got)
	}
	if got := cfg.GetRuntime().GetControlPlaneUrl(); got != "http://kagent-controller.kagent.svc.cluster.local:8083" {
		t.Errorf("control plane URL = %q", got)
	}
}
