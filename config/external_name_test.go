package config

import (
	"context"
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

// attrFlags reports whether attrs has the given attribute, and if so, its
// required/optional/computed flags as recorded in schema.json.
func attrFlags(attrs map[string]any, name string) (found, required, optional, computed bool) {
	raw, ok := attrs[name]
	if !ok {
		return false, false, false, false
	}
	m, _ := raw.(map[string]any)
	req, _ := m["required"].(bool)
	opt, _ := m["optional"].(bool)
	comp, _ := m["computed"].(bool)
	return true, req, opt, comp
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

// TestExternalNameMatchesSchema replaces TestNameIdentifiedResourcesHaveNameAttribute:
// upjet's plugin-framework controller always calls GetExternalNameFn on the
// tfstate map (external_tfpluginfw.go setExternalName, ~line 962), regardless
// of what the resource's Terraform schema contains. So the class picked for
// each resource must only read attributes that actually exist in its schema.
func TestExternalNameMatchesSchema(t *testing.T) {
	schemas := loadResourceSchemas(t)

	for _, name := range nameFromStateResources {
		attrs, ok := schemas[name]
		if !ok {
			t.Errorf("%q classified nameFromState but is not in schema.json", name)
			continue
		}
		if _, hasName := attrs["name"]; !hasName {
			t.Errorf("%q classified nameFromState (reads tfstate[\"name\"]) but schema has no `name` attribute", name)
		}
		if _, hasID := attrs["id"]; hasID {
			t.Errorf("%q classified nameFromState but schema HAS an `id` attribute; move it to providerIdentified or role", name)
		}
	}

	for _, name := range providerIdentified {
		attrs, ok := schemas[name]
		if !ok {
			t.Errorf("%q classified providerIdentified but is not in schema.json", name)
			continue
		}
		if _, hasID := attrs["id"]; !hasID {
			t.Errorf("%q classified providerIdentified (config.IdentifierFromProvider reads tfstate[\"id\"]) but schema has no `id` attribute", name)
		}
	}

	for _, name := range singletonResources {
		attrs, ok := schemas[name]
		if !ok {
			t.Errorf("%q classified singleton but is not in schema.json", name)
			continue
		}
		if _, hasID := attrs["id"]; hasID {
			t.Errorf("%q classified singleton but schema HAS an `id` attribute", name)
		}
		if _, hasName := attrs["name"]; hasName {
			t.Errorf("%q classified singleton but schema HAS a `name` attribute", name)
		}
	}

	for _, name := range userResources {
		attrs, ok := schemas[name]
		if !ok {
			t.Errorf("%q classified user but is not in schema.json", name)
			continue
		}
		if _, hasUserID := attrs["user_id"]; !hasUserID {
			t.Errorf("%q classified user (reads tfstate[\"user_id\"]) but schema has no `user_id` attribute", name)
		}
	}

	for _, name := range roleResources {
		attrs, ok := schemas[name]
		if !ok {
			t.Errorf("%q classified role but is not in schema.json", name)
			continue
		}
		found, required, optional, computed := attrFlags(attrs, "id")
		if !found {
			t.Errorf("%q classified role but schema has no `id` attribute", name)
			continue
		}
		if !required && !optional {
			t.Errorf("%q id attribute is computed-only (required=%v optional=%v computed=%v); "+
				"role should use config.IdentifierFromProvider instead of config.ParameterAsIdentifier(\"id\")",
				name, required, optional, computed)
		}
	}

	all := len(nameFromStateResources) + len(providerIdentified) + len(singletonResources) + len(userResources) + len(roleResources)
	if all != 119 {
		t.Errorf("expected 119 total classified resources across all classes, got %d", all)
	}
}

func TestNoDuplicateExternalNameEntries(t *testing.T) {
	seen := map[string]bool{}
	for _, l := range [][]string{nameFromStateResources, providerIdentified, singletonResources, userResources, roleResources} {
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

// --- behavioural tests: one fake-state call per class ---

func TestNameFromStateGetExternalName(t *testing.T) {
	got, err := nameFromState.GetExternalNameFn(map[string]any{"name": "x"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != "x" {
		t.Errorf("got %q, want %q", got, "x")
	}

	if _, err := nameFromState.GetExternalNameFn(map[string]any{}); err == nil {
		t.Error("expected error when `name` is missing from tfstate, got nil")
	}
	if _, err := nameFromState.GetExternalNameFn(map[string]any{"name": ""}); err == nil {
		t.Error("expected error when `name` is empty in tfstate, got nil")
	}
}

func TestNameFromStateSetIdentifierArgument(t *testing.T) {
	base := map[string]any{}
	nameFromState.SetIdentifierArgumentFn(base, "x")
	if base["name"] != "x" {
		t.Errorf("base[\"name\"] = %v, want %q", base["name"], "x")
	}
}

func TestSingletonExternalName(t *testing.T) {
	got, err := singleton.GetExternalNameFn(map[string]any{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != defaultExternalName {
		t.Errorf("got %q, want %q", got, defaultExternalName)
	}

	// Ignores whatever is already in tfstate: the constant always wins.
	got, err = singleton.GetExternalNameFn(map[string]any{"id": "something-else"})
	if err != nil || got != defaultExternalName {
		t.Errorf("got (%q, %v), want (%q, nil)", got, err, defaultExternalName)
	}

	id, err := singleton.GetIDFn(context.Background(), "whatever", nil, nil)
	if err != nil || id != defaultExternalName {
		t.Errorf("GetIDFn got (%q, %v), want (%q, nil)", id, err, defaultExternalName)
	}
}

func TestSingletonSetIdentifierArgumentIsNop(t *testing.T) {
	base := map[string]any{}
	singleton.SetIdentifierArgumentFn(base, "whatever")
	if len(base) != 0 {
		t.Errorf("expected SetIdentifierArgumentFn to be a no-op, got %v", base)
	}
}

func TestUserExternalName(t *testing.T) {
	got, err := userExternalName.GetExternalNameFn(map[string]any{"user_id": "u"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != "u" {
		t.Errorf("got %q, want %q", got, "u")
	}

	if _, err := userExternalName.GetExternalNameFn(map[string]any{}); err == nil {
		t.Error("expected error when `user_id` is missing from tfstate, got nil")
	}
}

func TestUserSetIdentifierArgument(t *testing.T) {
	base := map[string]any{}
	userExternalName.SetIdentifierArgumentFn(base, "u")
	if base["user_id"] != "u" {
		t.Errorf("base[\"user_id\"] = %v, want %q", base["user_id"], "u")
	}
}

func TestRoleExternalName(t *testing.T) {
	role, ok := ExternalNameConfigs["sonatyperepo_role"]
	if !ok {
		t.Fatal("no external-name configuration for sonatyperepo_role")
	}
	base := map[string]any{}
	role.SetIdentifierArgumentFn(base, "x")
	if base["id"] != "x" {
		t.Errorf("base[\"id\"] = %v, want %q", base["id"], "x")
	}
	got, err := role.GetExternalNameFn(map[string]any{"id": "x"})
	if err != nil || got != "x" {
		t.Errorf("got (%q, %v), want (\"x\", nil)", got, err)
	}
}
