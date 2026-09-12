package config

import (
	"strings"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

// GroupKindOverrides fixes the API group and kind of resources whose
// Terraform name does not split into a sensible default. Upjet derives
// group = first token and kind = CamelCase(rest): blob_store_s3 would become
// "blob"/"StoreS3", and user/role have no group token at all.
//
// The three sonatyperepo_repository_ruby_gems_{group,hosted,proxy} resources
// are deprecated pre-rename aliases of sonatyperepo_repository_rubygems_*
// (both exist in the schema; see config/provider-metadata.yaml). Upjet's
// default Kind ("RubyGemsHosted" vs "RubygemsHosted") differs only in the
// casing of one letter, so the two resources generate the same lowercased
// output filename ("zz_rubygemshosted_types.go") and one silently clobbers
// the other on disk — `make generate` then fails angryjet/build with
// "undefined: RubyGemsHosted" (hit 2026-09-13, non-deterministic in which of
// group/hosted/proxy loses the race). Give the legacy trio distinct Kinds so
// both sides of the alias keep their own generated files.
func GroupKindOverrides() ujconfig.ResourceOption {
	return func(r *ujconfig.Resource) {
		switch {
		case strings.HasPrefix(r.Name, "sonatyperepo_blob_store_"):
			r.ShortGroup = "blobstore"
			r.Kind = camel(strings.TrimPrefix(r.Name, "sonatyperepo_blob_store_"))
		case r.Name == "sonatyperepo_user":
			r.ShortGroup, r.Kind = "security", "User"
		case r.Name == "sonatyperepo_role":
			r.ShortGroup, r.Kind = "security", "Role"
		case r.Name == "sonatyperepo_repository_ruby_gems_group":
			r.Kind = "RubygemsLegacyGroup"
		case r.Name == "sonatyperepo_repository_ruby_gems_hosted":
			r.Kind = "RubygemsLegacyHosted"
		case r.Name == "sonatyperepo_repository_ruby_gems_proxy":
			r.Kind = "RubygemsLegacyProxy"
		}
	}
}

// camel turns snake_case into CamelCase: "s3" -> "S3", "group" -> "Group".
func camel(s string) string {
	parts := strings.Split(s, "_")
	for i, p := range parts {
		if p != "" {
			parts[i] = strings.ToUpper(p[:1]) + p[1:]
		}
	}
	return strings.Join(parts, "")
}
