package config

import (
	"testing"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

func TestGroupKindOverrides(t *testing.T) {
	cases := map[string]struct{ group, kind string }{
		"sonatyperepo_blob_store_s3":    {"blobstore", "S3"},
		"sonatyperepo_blob_store_file":  {"blobstore", "File"},
		"sonatyperepo_blob_store_group": {"blobstore", "Group"},
		"sonatyperepo_user":             {"security", "User"},
		"sonatyperepo_role":             {"security", "Role"},
	}
	for name, want := range cases {
		r := &ujconfig.Resource{Name: name, ShortGroup: "default", Kind: "Default"}
		GroupKindOverrides()(r)
		if r.ShortGroup != want.group || r.Kind != want.kind {
			t.Errorf("%s: got %s/%s, want %s/%s", name, r.ShortGroup, r.Kind, want.group, want.kind)
		}
	}
	// Resources with a good Upjet default must be left untouched.
	r := &ujconfig.Resource{Name: "sonatyperepo_repository_docker_hosted", ShortGroup: "repository", Kind: "DockerHosted"}
	GroupKindOverrides()(r)
	if r.ShortGroup != "repository" || r.Kind != "DockerHosted" {
		t.Errorf("default naming was altered: %s/%s", r.ShortGroup, r.Kind)
	}
}
