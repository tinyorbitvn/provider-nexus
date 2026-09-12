package config

import (
	"strings"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

// GroupKindOverrides fixes the API group and kind of resources whose
// Terraform name does not split into a sensible default. Upjet derives
// group = first token and kind = CamelCase(rest): blob_store_s3 would become
// "blob"/"StoreS3", and user/role have no group token at all.
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
