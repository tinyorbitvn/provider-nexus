package config

import (
	"context"
	"errors"

	"github.com/crossplane/upjet/v2/pkg/config"
)

// Facts (config/schema.json, resource_schemas, verified 2026-09-13): of the
// 119 sonatyperepo_* resources, 25 have a top-level `id` attribute and 94 do
// not. Upjet's plugin-framework controller
// (external_tfpluginfw.go:setExternalName, ~line 962) always calls
// cfg.ExternalName.GetExternalNameFn(stateValueMap) on every
// Observe/Create/Update, no matter what the resource's Terraform schema
// contains — there is no schema check on that path (the schema check only
// gates whether "id" gets written *into* tfstate/params, in
// getFrameworkExtendedParameters and Connect, ~lines 156/229). The upjet
// default GetExternalNameFn (IDAsExternalName, pkg/config/resource.go) reads
// tfstate["id"] and returns the literal error "cannot find id in tfstate"
// when that key is absent — exactly the failure seen against the real Nexus
// (Task 9 round 1) for every one of the 94 id-less kinds that used to be
// classified NameAsIdentifier/IdentifierFromProvider (both of which inherit
// GetExternalNameFn: IDAsExternalName). The five classes below replace that
// with a GetExternalNameFn that only reads attributes the schema guarantees.

// nameFromStateResources are the 84 resources with a `name` attribute and no
// `id` attribute. Terraform import for these uses NAME as the id
// (`terraform import <resource>.x <NAME>`), but because there is no `id` in
// the schema, the external name must be read back from tfstate["name"], not
// tfstate["id"]. This is why config.ParameterAsIdentifier("name") is NOT
// enough here: it only changes SetIdentifierArgumentFn/OmittedFields, and
// keeps the inherited GetExternalNameFn: IDAsExternalName from
// NameAsIdentifier (pkg/config/externalname.go), which still reads
// tfstate["id"] and would still fail with "cannot find id in tfstate". The
// nameFromState ExternalName below is therefore a hand-written literal with
// its own GetExternalNameFn.
var nameFromStateResources = []string{
	"sonatyperepo_blob_store_acs",
	"sonatyperepo_blob_store_file",
	"sonatyperepo_blob_store_gcs",
	"sonatyperepo_blob_store_group",
	"sonatyperepo_blob_store_s3",
	"sonatyperepo_cleanup_policy",
	"sonatyperepo_content_selector",
	"sonatyperepo_privilege_application",
	"sonatyperepo_privilege_repository_admin",
	"sonatyperepo_privilege_repository_content_selector",
	"sonatyperepo_privilege_repository_view",
	"sonatyperepo_privilege_script",
	"sonatyperepo_privilege_wildcard",
	"sonatyperepo_repository_alpine_group",
	"sonatyperepo_repository_alpine_hosted",
	"sonatyperepo_repository_alpine_proxy",
	"sonatyperepo_repository_ansiblegalaxy_group",
	"sonatyperepo_repository_ansiblegalaxy_hosted",
	"sonatyperepo_repository_ansiblegalaxy_proxy",
	"sonatyperepo_repository_apt_hosted",
	"sonatyperepo_repository_apt_proxy",
	"sonatyperepo_repository_cargo_group",
	"sonatyperepo_repository_cargo_hosted",
	"sonatyperepo_repository_cargo_proxy",
	"sonatyperepo_repository_cocoapods_proxy",
	"sonatyperepo_repository_composer_proxy",
	"sonatyperepo_repository_conan_group",
	"sonatyperepo_repository_conan_hosted",
	"sonatyperepo_repository_conan_proxy",
	"sonatyperepo_repository_conda_proxy",
	"sonatyperepo_repository_docker_group",
	"sonatyperepo_repository_docker_hosted",
	"sonatyperepo_repository_docker_proxy",
	"sonatyperepo_repository_gitlfs_hosted",
	"sonatyperepo_repository_go_group",
	"sonatyperepo_repository_go_hosted",
	"sonatyperepo_repository_go_proxy",
	"sonatyperepo_repository_helm_group",
	"sonatyperepo_repository_helm_hosted",
	"sonatyperepo_repository_helm_proxy",
	"sonatyperepo_repository_huggingface_proxy",
	"sonatyperepo_repository_maven2_group",
	"sonatyperepo_repository_maven2_hosted",
	"sonatyperepo_repository_maven2_proxy",
	"sonatyperepo_repository_maven_group",
	"sonatyperepo_repository_maven_hosted",
	"sonatyperepo_repository_maven_proxy",
	"sonatyperepo_repository_npm_group",
	"sonatyperepo_repository_npm_hosted",
	"sonatyperepo_repository_npm_proxy",
	"sonatyperepo_repository_nuget_group",
	"sonatyperepo_repository_nuget_hosted",
	"sonatyperepo_repository_nuget_proxy",
	"sonatyperepo_repository_oci_group",
	"sonatyperepo_repository_oci_hosted",
	"sonatyperepo_repository_oci_proxy",
	"sonatyperepo_repository_p2_proxy",
	"sonatyperepo_repository_pub_group",
	"sonatyperepo_repository_pub_hosted",
	"sonatyperepo_repository_pub_proxy",
	"sonatyperepo_repository_pypi_group",
	"sonatyperepo_repository_pypi_hosted",
	"sonatyperepo_repository_pypi_proxy",
	"sonatyperepo_repository_r_group",
	"sonatyperepo_repository_r_hosted",
	"sonatyperepo_repository_r_proxy",
	"sonatyperepo_repository_raw_group",
	"sonatyperepo_repository_raw_hosted",
	"sonatyperepo_repository_raw_proxy",
	"sonatyperepo_repository_ruby_gems_group",
	"sonatyperepo_repository_ruby_gems_hosted",
	"sonatyperepo_repository_ruby_gems_proxy",
	"sonatyperepo_repository_rubygems_group",
	"sonatyperepo_repository_rubygems_hosted",
	"sonatyperepo_repository_rubygems_proxy",
	"sonatyperepo_repository_swift_group",
	"sonatyperepo_repository_swift_proxy",
	"sonatyperepo_repository_terraform_group",
	"sonatyperepo_repository_terraform_hosted",
	"sonatyperepo_repository_terraform_proxy",
	"sonatyperepo_repository_yum_group",
	"sonatyperepo_repository_yum_hosted",
	"sonatyperepo_repository_yum_proxy",
	"sonatyperepo_routing_rule",
}

// providerIdentified lists the 24 resources whose Terraform schema HAS an
// `id` attribute assigned by Nexus itself (capability and task IDs,
// certificate IDs, the fixed SECURITY_REALMS/LDAP tokens): unchanged from
// before this fix, since these already work against the real Nexus
// (measured for security_realms and capability_rut_auth, Task 9 round 1).
// config.IdentifierFromProvider's GetExternalNameFn (IDAsExternalName) reads
// tfstate["id"], which the schema guarantees exists here.
var providerIdentified = []string{
	"sonatyperepo_capability_audit",
	"sonatyperepo_capability_base_url",
	"sonatyperepo_capability_custom_s3_regions",
	"sonatyperepo_capability_default_role",
	"sonatyperepo_capability_firewall_audit_and_quarantine",
	"sonatyperepo_capability_healthcheck",
	"sonatyperepo_capability_outreach_management",
	"sonatyperepo_capability_rut_auth",
	"sonatyperepo_capability_storage_settings",
	"sonatyperepo_capability_ui_branding",
	"sonatyperepo_capability_ui_settings",
	"sonatyperepo_capability_webhook_global",
	"sonatyperepo_capability_webhook_repository",
	"sonatyperepo_security_realms",
	"sonatyperepo_security_ssl_truststore",
	"sonatyperepo_system_config_ldap_connection",
	"sonatyperepo_task_blobstore_compact",
	"sonatyperepo_task_license_expiration_notification",
	"sonatyperepo_task_malware_remediator",
	"sonatyperepo_task_repair_create_browse_nodes",
	"sonatyperepo_task_repository_docker_gc",
	"sonatyperepo_task_repository_docker_upload_purge",
	"sonatyperepo_task_repository_maven_remove_snapshots",
	"sonatyperepo_task_repository_purge_unused",
}

// singletonResources are the 9 resources with NEITHER `id` NOR `name`: one
// object per Nexus server (there is nothing to key on, and Nexus itself
// exposes no distinguishing identifier for them over this API). The external
// name is therefore a fixed constant, ignoring tfstate entirely — using
// tfstate["id"] here (the previous providerIdentified behaviour) was exactly
// the "cannot find id in tfstate" bug for these 9 kinds.
var singletonResources = []string{
	"sonatyperepo_security_oauth2",
	"sonatyperepo_security_saml",
	"sonatyperepo_security_ssrf_protection",
	"sonatyperepo_security_user_tokens",
	"sonatyperepo_system_anonymous_access",
	"sonatyperepo_system_config_http",
	"sonatyperepo_system_config_mail",
	"sonatyperepo_system_config_product_license",
	"sonatyperepo_system_iq_connection",
}

// userResources: sonatyperepo_user has neither `id` nor `name`, but it does
// have `user_id` (required), which is Nexus's actual unique key for a user.
// The Terraform import ID for this resource is the composite `<user_id>,
// <source>`, but that composite is a Terraform-import-only convention; upjet
// never calls terraform import, so only user_id matters here.
var userResources = []string{
	"sonatyperepo_user",
}

// roleResources: sonatyperepo_role is the one resource that has BOTH `id`
// and `name`. Its `id` attribute is `required` in the schema (user-settable,
// not computed-only, verified 2026-09-13), i.e. Nexus lets the caller choose
// the role's id and treats it as the unique key; `name` is just a display
// label. So the identifier is `id`, not `name`. Unlike nameFromState above,
// config.ParameterAsIdentifier("id") is exactly right here: it sets
// base["id"] = externalName and keeps the inherited GetExternalNameFn
// (IDAsExternalName), which reads tfstate["id"] — and the schema guarantees
// "id" exists for this resource, so that default is correct, not a bug.
var roleResources = []string{
	"sonatyperepo_role",
}

// nameFromState is a hand-written config.ExternalName (see the
// nameFromStateResources comment for why config.ParameterAsIdentifier is not
// sufficient): the external name is the Nexus `name`, read back from
// tfstate["name"] instead of the upjet default tfstate["id"].
var nameFromState = config.ExternalName{
	SetIdentifierArgumentFn: func(base map[string]any, externalName string) {
		base["name"] = externalName
	},
	GetExternalNameFn: func(tfstate map[string]any) (string, error) {
		name, ok := tfstate["name"].(string)
		if !ok || name == "" {
			return "", errors.New("cannot find name in tfstate")
		}
		return name, nil
	},
	GetIDFn: config.ExternalNameAsID,
	OmittedFields: []string{
		"name",
	},
}

// defaultExternalName is the fixed external name returned for every
// singletonResources kind, regardless of tfstate.
const defaultExternalName = "default"

// singleton is a hand-written config.ExternalName for the 9 one-per-server
// resources: the external name is always the constant defaultExternalName,
// independent of tfstate or the previously-set external name.
// DisableNameInitializer is deliberately left false (the zero value) so
// metadata.name seeds the crossplane.io/external-name annotation on first
// Observe instead of leaving it empty; GetExternalNameFn then immediately
// normalizes it to defaultExternalName.
var singleton = config.ExternalName{
	SetIdentifierArgumentFn: config.NopSetIdentifierArgument,
	GetExternalNameFn: func(_ map[string]any) (string, error) {
		return defaultExternalName, nil
	},
	GetIDFn: func(_ context.Context, _ string, _ map[string]any, _ map[string]any) (string, error) {
		return defaultExternalName, nil
	},
}

// userExternalName is a hand-written config.ExternalName for
// sonatyperepo_user: the external name is Nexus's user_id, read back from
// tfstate["user_id"] instead of the upjet default tfstate["id"]. `source`
// stays a normal read-only field; the Terraform import ID `uid,SOURCE` is a
// terraform-import-only convention and irrelevant to upjet.
var userExternalName = config.ExternalName{
	SetIdentifierArgumentFn: func(base map[string]any, externalName string) {
		base["user_id"] = externalName
	},
	GetExternalNameFn: func(tfstate map[string]any) (string, error) {
		userID, ok := tfstate["user_id"].(string)
		if !ok || userID == "" {
			return "", errors.New("cannot find user_id in tfstate")
		}
		return userID, nil
	},
	GetIDFn: config.ExternalNameAsID,
	OmittedFields: []string{
		"user_id",
	},
}

// ExternalNameConfigs contains all external name configurations for this
// provider. Both scopes (cluster and namespaced) share it.
var ExternalNameConfigs = map[string]config.ExternalName{}

func init() {
	for _, r := range nameFromStateResources {
		ExternalNameConfigs[r] = nameFromState
	}
	for _, r := range providerIdentified {
		ExternalNameConfigs[r] = config.IdentifierFromProvider
	}
	for _, r := range singletonResources {
		ExternalNameConfigs[r] = singleton
	}
	for _, r := range userResources {
		ExternalNameConfigs[r] = userExternalName
	}
	for _, r := range roleResources {
		ExternalNameConfigs[r] = config.ParameterAsIdentifier("id")
	}
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
