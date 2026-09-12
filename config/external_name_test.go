package config

import (
	"encoding/json"
	"os"
	"testing"
)

const schemaKey = "registry.terraform.io/sonatype-nexus-community/sonatyperepo"

func loadResourceSchemas(t *testing.T) map[string]map[string]any {
	t.Helper()
	raw, err := os.ReadFile("schema.json")
	if err != nil {
		t.Fatalf("read schema.json: %v", err)
	}
	var s struct {
		ProviderSchemas map[string]struct {
			ResourceSchemas map[string]struct {
				Block struct {
					Attributes map[string]any `json:"attributes"`
				} `json:"block"`
			} `json:"resource_schemas"`
		} `json:"provider_schemas"`
	}
	if err := json.Unmarshal(raw, &s); err != nil {
		t.Fatalf("parse schema.json: %v", err)
	}
	out := map[string]map[string]any{}
	for name, rs := range s.ProviderSchemas[schemaKey].ResourceSchemas {
		out[name] = rs.Block.Attributes
	}
	return out
}

// Every Terraform resource must have exactly one external-name configuration,
// and every configured name must exist in the provider schema.
func TestExternalNameConfigsCoverSchema(t *testing.T) {
	schemas := loadResourceSchemas(t)
	if len(schemas) != 119 {
		t.Fatalf("expected 119 resources in schema.json, got %d", len(schemas))
	}
	for name := range schemas {
		if _, ok := ExternalNameConfigs[name]; !ok {
			t.Errorf("resource %q has no external-name configuration", name)
		}
	}
	for name := range ExternalNameConfigs {
		if _, ok := schemas[name]; !ok {
			t.Errorf("external-name configured for unknown resource %q", name)
		}
	}
}

// NameAsIdentifier only works when the resource has a "name" argument.
func TestNameIdentifiedResourcesHaveNameAttribute(t *testing.T) {
	schemas := loadResourceSchemas(t)
	for _, name := range nameIdentified {
		if _, ok := schemas[name]["name"]; !ok {
			t.Errorf("%q is configured NameAsIdentifier but has no `name` attribute; move it to providerIdentified", name)
		}
	}
}

func TestNoDuplicateExternalNameEntries(t *testing.T) {
	seen := map[string]bool{}
	for _, l := range [][]string{nameIdentified, providerIdentified} {
		for _, n := range l {
			if seen[n] {
				t.Errorf("%q listed twice", n)
			}
			seen[n] = true
		}
	}
	if len(seen) != 119 {
		t.Errorf("expected 119 unique entries, got %d", len(seen))
	}
}
