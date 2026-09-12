package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	"terraform-provider-sonatyperepo/xpprovider"

	"github.com/tinyorbitvn/provider-nexus/internal/version"
)

const (
	resourcePrefix = "sonatyperepo"
	modulePath     = "github.com/tinyorbitvn/provider-nexus"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("nexus.tinyorbit.vn"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		// Every resource is served in-process by the embedded plugin-framework
		// provider; nothing falls back to the Terraform CLI.
		ujconfig.WithTerraformPluginFrameworkIncludeList(ExternalNameConfigured()),
		ujconfig.WithTerraformPluginFrameworkProvider(xpprovider.New(version.Version)),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
			GroupKindOverrides(),
		))

	pc.ConfigureResources()
	return pc
}

// GetProviderNamespaced returns the namespaced provider configuration
func GetProviderNamespaced() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("nexus.m.tinyorbit.vn"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		// Every resource is served in-process by the embedded plugin-framework
		// provider; nothing falls back to the Terraform CLI.
		ujconfig.WithTerraformPluginFrameworkIncludeList(ExternalNameConfigured()),
		ujconfig.WithTerraformPluginFrameworkProvider(xpprovider.New(version.Version)),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
			GroupKindOverrides(),
		),
		ujconfig.WithExampleManifestConfiguration(ujconfig.ExampleManifestConfiguration{
			ManagedResourceNamespace: "crossplane-system",
		}))

	pc.ConfigureResources()
	return pc
}
