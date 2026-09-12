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
		// Deprecated pre-rename aliases of sonatyperepo_repository_rubygems_*
		// (both exist in the schema): the default Kind differs from the
		// canonical resource's only by the casing of one letter
		// ("RubyGemsHosted" vs "RubygemsHosted"), so both lowercase to the
		// same generated filename and one clobbers the other on disk during
		// `make generate` (hit 2026-09-13). Only Kind is overridden here;
		// ShortGroup keeps Upjet's default ("repository"), which is why the
		// expectation below is "default" — same as the untouched cases.
		"sonatyperepo_repository_ruby_gems_group":  {"default", "RubygemsLegacyGroup"},
		"sonatyperepo_repository_ruby_gems_hosted": {"default", "RubygemsLegacyHosted"},
		"sonatyperepo_repository_ruby_gems_proxy":  {"default", "RubygemsLegacyProxy"},
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
