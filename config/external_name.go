package config

import "github.com/crossplane/upjet/v2/pkg/config"

// nameIdentified lists resources whose Terraform ID is the object's own
// `name` (docs: `terraform import <resource>.x <NAME>`). The Kubernetes
// external name is therefore the Nexus name, and an existing object is
// adopted by setting the crossplane.io/external-name annotation.
var nameIdentified = []string{
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
	"sonatyperepo_role",
	"sonatyperepo_routing_rule",
}

// providerIdentified lists resources whose ID is assigned by Nexus
// (capability and task IDs, certificate IDs), singletons whose ID is a fixed
// token returned by the provider (SECURITY_REALMS, ANONYMOUS_ACCESS,
// SSRF_PROTECTION, OAUTH2, SAML, CONFIG_MAIL, ...), and user, whose ID is the
// composite `<user_id>,<source>` (e.g. admin,DEFAULT).
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
	"sonatyperepo_security_oauth2",
	"sonatyperepo_security_realms",
	"sonatyperepo_security_saml",
	"sonatyperepo_security_ssl_truststore",
	"sonatyperepo_security_ssrf_protection",
	"sonatyperepo_security_user_tokens",
	"sonatyperepo_system_anonymous_access",
	"sonatyperepo_system_config_http",
	"sonatyperepo_system_config_ldap_connection",
	"sonatyperepo_system_config_mail",
	"sonatyperepo_system_config_product_license",
	"sonatyperepo_system_iq_connection",
	"sonatyperepo_task_blobstore_compact",
	"sonatyperepo_task_license_expiration_notification",
	"sonatyperepo_task_malware_remediator",
	"sonatyperepo_task_repair_create_browse_nodes",
	"sonatyperepo_task_repository_docker_gc",
	"sonatyperepo_task_repository_docker_upload_purge",
	"sonatyperepo_task_repository_maven_remove_snapshots",
	"sonatyperepo_task_repository_purge_unused",
	"sonatyperepo_user",
}

// ExternalNameConfigs contains all external name configurations for this
// provider. Both scopes (cluster and namespaced) share it.
var ExternalNameConfigs = map[string]config.ExternalName{}

func init() {
	for _, r := range nameIdentified {
		ExternalNameConfigs[r] = config.NameAsIdentifier
	}
	for _, r := range providerIdentified {
		ExternalNameConfigs[r] = config.IdentifierFromProvider
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
