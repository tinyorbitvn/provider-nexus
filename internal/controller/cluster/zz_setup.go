// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	acs "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/blobstore/acs"
	file "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/blobstore/file"
	gcs "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/blobstore/gcs"
	group "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/blobstore/group"
	s3 "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/blobstore/s3"
	audit "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/capability/audit"
	baseurl "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/capability/baseurl"
	customs3regions "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/capability/customs3regions"
	defaultrole "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/capability/defaultrole"
	firewallauditandquarantine "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/capability/firewallauditandquarantine"
	healthcheck "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/capability/healthcheck"
	outreachmanagement "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/capability/outreachmanagement"
	rutauth "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/capability/rutauth"
	storagesettings "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/capability/storagesettings"
	uibranding "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/capability/uibranding"
	uisettings "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/capability/uisettings"
	webhookglobal "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/capability/webhookglobal"
	webhookrepository "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/capability/webhookrepository"
	policy "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/cleanup/policy"
	selector "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/content/selector"
	application "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/privilege/application"
	repositoryadmin "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/privilege/repositoryadmin"
	repositorycontentselector "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/privilege/repositorycontentselector"
	repositoryview "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/privilege/repositoryview"
	script "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/privilege/script"
	wildcard "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/privilege/wildcard"
	providerconfig "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/providerconfig"
	alpinegroup "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/repository/alpinegroup"
	alpinehosted "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/repository/alpinehosted"
	alpineproxy "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/repository/alpineproxy"
	ansiblegalaxygroup "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/repository/ansiblegalaxygroup"
	ansiblegalaxyhosted "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/repository/ansiblegalaxyhosted"
	ansiblegalaxyproxy "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/repository/ansiblegalaxyproxy"
	apthosted "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/repository/apthosted"
	aptproxy "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/repository/aptproxy"
	cargogroup "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/repository/cargogroup"
	cargohosted "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/repository/cargohosted"
	cargoproxy "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/repository/cargoproxy"
	cocoapodsproxy "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/repository/cocoapodsproxy"
	composerproxy "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/repository/composerproxy"
	conangroup "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/repository/conangroup"
	conanhosted "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/repository/conanhosted"
	conanproxy "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/repository/conanproxy"
	condaproxy "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/repository/condaproxy"
	dockergroup "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/repository/dockergroup"
	dockerhosted "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/repository/dockerhosted"
	dockerproxy "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/repository/dockerproxy"
	gitlfshosted "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/repository/gitlfshosted"
	gogroup "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/repository/gogroup"
	gohosted "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/repository/gohosted"
	goproxy "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/repository/goproxy"
	helmgroup "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/repository/helmgroup"
	helmhosted "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/repository/helmhosted"
	helmproxy "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/repository/helmproxy"
	huggingfaceproxy "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/repository/huggingfaceproxy"
	maven2group "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/repository/maven2group"
	maven2hosted "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/repository/maven2hosted"
	maven2proxy "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/repository/maven2proxy"
	mavengroup "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/repository/mavengroup"
	mavenhosted "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/repository/mavenhosted"
	mavenproxy "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/repository/mavenproxy"
	npmgroup "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/repository/npmgroup"
	npmhosted "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/repository/npmhosted"
	npmproxy "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/repository/npmproxy"
	nugetgroup "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/repository/nugetgroup"
	nugethosted "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/repository/nugethosted"
	nugetproxy "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/repository/nugetproxy"
	ocigroup "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/repository/ocigroup"
	ocihosted "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/repository/ocihosted"
	ociproxy "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/repository/ociproxy"
	p2proxy "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/repository/p2proxy"
	pubgroup "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/repository/pubgroup"
	pubhosted "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/repository/pubhosted"
	pubproxy "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/repository/pubproxy"
	pypigroup "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/repository/pypigroup"
	pypihosted "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/repository/pypihosted"
	pypiproxy "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/repository/pypiproxy"
	rawgroup "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/repository/rawgroup"
	rawhosted "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/repository/rawhosted"
	rawproxy "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/repository/rawproxy"
	rgroup "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/repository/rgroup"
	rhosted "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/repository/rhosted"
	rproxy "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/repository/rproxy"
	rubygemsgroup "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/repository/rubygemsgroup"
	rubygemshosted "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/repository/rubygemshosted"
	rubygemslegacygroup "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/repository/rubygemslegacygroup"
	rubygemslegacyhosted "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/repository/rubygemslegacyhosted"
	rubygemslegacyproxy "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/repository/rubygemslegacyproxy"
	rubygemsproxy "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/repository/rubygemsproxy"
	swiftgroup "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/repository/swiftgroup"
	swiftproxy "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/repository/swiftproxy"
	terraformgroup "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/repository/terraformgroup"
	terraformhosted "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/repository/terraformhosted"
	terraformproxy "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/repository/terraformproxy"
	yumgroup "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/repository/yumgroup"
	yumhosted "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/repository/yumhosted"
	yumproxy "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/repository/yumproxy"
	rule "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/routing/rule"
	oauth2 "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/security/oauth2"
	realms "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/security/realms"
	role "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/security/role"
	saml "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/security/saml"
	ssltruststore "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/security/ssltruststore"
	ssrfprotection "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/security/ssrfprotection"
	user "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/security/user"
	usertokens "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/security/usertokens"
	anonymousaccess "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/system/anonymousaccess"
	confighttp "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/system/confighttp"
	configldapconnection "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/system/configldapconnection"
	configmail "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/system/configmail"
	configproductlicense "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/system/configproductlicense"
	iqconnection "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/system/iqconnection"
	blobstorecompact "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/task/blobstorecompact"
	licenseexpirationnotification "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/task/licenseexpirationnotification"
	malwareremediator "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/task/malwareremediator"
	repaircreatebrowsenodes "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/task/repaircreatebrowsenodes"
	repositorydockergc "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/task/repositorydockergc"
	repositorydockeruploadpurge "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/task/repositorydockeruploadpurge"
	repositorymavenremovesnapshots "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/task/repositorymavenremovesnapshots"
	repositorypurgeunused "github.com/tinyorbitvn/provider-nexus/internal/controller/cluster/task/repositorypurgeunused"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		acs.Setup,
		file.Setup,
		gcs.Setup,
		group.Setup,
		s3.Setup,
		audit.Setup,
		baseurl.Setup,
		customs3regions.Setup,
		defaultrole.Setup,
		firewallauditandquarantine.Setup,
		healthcheck.Setup,
		outreachmanagement.Setup,
		rutauth.Setup,
		storagesettings.Setup,
		uibranding.Setup,
		uisettings.Setup,
		webhookglobal.Setup,
		webhookrepository.Setup,
		policy.Setup,
		selector.Setup,
		application.Setup,
		repositoryadmin.Setup,
		repositorycontentselector.Setup,
		repositoryview.Setup,
		script.Setup,
		wildcard.Setup,
		providerconfig.Setup,
		alpinegroup.Setup,
		alpinehosted.Setup,
		alpineproxy.Setup,
		ansiblegalaxygroup.Setup,
		ansiblegalaxyhosted.Setup,
		ansiblegalaxyproxy.Setup,
		apthosted.Setup,
		aptproxy.Setup,
		cargogroup.Setup,
		cargohosted.Setup,
		cargoproxy.Setup,
		cocoapodsproxy.Setup,
		composerproxy.Setup,
		conangroup.Setup,
		conanhosted.Setup,
		conanproxy.Setup,
		condaproxy.Setup,
		dockergroup.Setup,
		dockerhosted.Setup,
		dockerproxy.Setup,
		gitlfshosted.Setup,
		gogroup.Setup,
		gohosted.Setup,
		goproxy.Setup,
		helmgroup.Setup,
		helmhosted.Setup,
		helmproxy.Setup,
		huggingfaceproxy.Setup,
		maven2group.Setup,
		maven2hosted.Setup,
		maven2proxy.Setup,
		mavengroup.Setup,
		mavenhosted.Setup,
		mavenproxy.Setup,
		npmgroup.Setup,
		npmhosted.Setup,
		npmproxy.Setup,
		nugetgroup.Setup,
		nugethosted.Setup,
		nugetproxy.Setup,
		ocigroup.Setup,
		ocihosted.Setup,
		ociproxy.Setup,
		p2proxy.Setup,
		pubgroup.Setup,
		pubhosted.Setup,
		pubproxy.Setup,
		pypigroup.Setup,
		pypihosted.Setup,
		pypiproxy.Setup,
		rawgroup.Setup,
		rawhosted.Setup,
		rawproxy.Setup,
		rgroup.Setup,
		rhosted.Setup,
		rproxy.Setup,
		rubygemsgroup.Setup,
		rubygemshosted.Setup,
		rubygemslegacygroup.Setup,
		rubygemslegacyhosted.Setup,
		rubygemslegacyproxy.Setup,
		rubygemsproxy.Setup,
		swiftgroup.Setup,
		swiftproxy.Setup,
		terraformgroup.Setup,
		terraformhosted.Setup,
		terraformproxy.Setup,
		yumgroup.Setup,
		yumhosted.Setup,
		yumproxy.Setup,
		rule.Setup,
		oauth2.Setup,
		realms.Setup,
		role.Setup,
		saml.Setup,
		ssltruststore.Setup,
		ssrfprotection.Setup,
		user.Setup,
		usertokens.Setup,
		anonymousaccess.Setup,
		confighttp.Setup,
		configldapconnection.Setup,
		configmail.Setup,
		configproductlicense.Setup,
		iqconnection.Setup,
		blobstorecompact.Setup,
		licenseexpirationnotification.Setup,
		malwareremediator.Setup,
		repaircreatebrowsenodes.Setup,
		repositorydockergc.Setup,
		repositorydockeruploadpurge.Setup,
		repositorymavenremovesnapshots.Setup,
		repositorypurgeunused.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		acs.SetupGated,
		file.SetupGated,
		gcs.SetupGated,
		group.SetupGated,
		s3.SetupGated,
		audit.SetupGated,
		baseurl.SetupGated,
		customs3regions.SetupGated,
		defaultrole.SetupGated,
		firewallauditandquarantine.SetupGated,
		healthcheck.SetupGated,
		outreachmanagement.SetupGated,
		rutauth.SetupGated,
		storagesettings.SetupGated,
		uibranding.SetupGated,
		uisettings.SetupGated,
		webhookglobal.SetupGated,
		webhookrepository.SetupGated,
		policy.SetupGated,
		selector.SetupGated,
		application.SetupGated,
		repositoryadmin.SetupGated,
		repositorycontentselector.SetupGated,
		repositoryview.SetupGated,
		script.SetupGated,
		wildcard.SetupGated,
		providerconfig.SetupGated,
		alpinegroup.SetupGated,
		alpinehosted.SetupGated,
		alpineproxy.SetupGated,
		ansiblegalaxygroup.SetupGated,
		ansiblegalaxyhosted.SetupGated,
		ansiblegalaxyproxy.SetupGated,
		apthosted.SetupGated,
		aptproxy.SetupGated,
		cargogroup.SetupGated,
		cargohosted.SetupGated,
		cargoproxy.SetupGated,
		cocoapodsproxy.SetupGated,
		composerproxy.SetupGated,
		conangroup.SetupGated,
		conanhosted.SetupGated,
		conanproxy.SetupGated,
		condaproxy.SetupGated,
		dockergroup.SetupGated,
		dockerhosted.SetupGated,
		dockerproxy.SetupGated,
		gitlfshosted.SetupGated,
		gogroup.SetupGated,
		gohosted.SetupGated,
		goproxy.SetupGated,
		helmgroup.SetupGated,
		helmhosted.SetupGated,
		helmproxy.SetupGated,
		huggingfaceproxy.SetupGated,
		maven2group.SetupGated,
		maven2hosted.SetupGated,
		maven2proxy.SetupGated,
		mavengroup.SetupGated,
		mavenhosted.SetupGated,
		mavenproxy.SetupGated,
		npmgroup.SetupGated,
		npmhosted.SetupGated,
		npmproxy.SetupGated,
		nugetgroup.SetupGated,
		nugethosted.SetupGated,
		nugetproxy.SetupGated,
		ocigroup.SetupGated,
		ocihosted.SetupGated,
		ociproxy.SetupGated,
		p2proxy.SetupGated,
		pubgroup.SetupGated,
		pubhosted.SetupGated,
		pubproxy.SetupGated,
		pypigroup.SetupGated,
		pypihosted.SetupGated,
		pypiproxy.SetupGated,
		rawgroup.SetupGated,
		rawhosted.SetupGated,
		rawproxy.SetupGated,
		rgroup.SetupGated,
		rhosted.SetupGated,
		rproxy.SetupGated,
		rubygemsgroup.SetupGated,
		rubygemshosted.SetupGated,
		rubygemslegacygroup.SetupGated,
		rubygemslegacyhosted.SetupGated,
		rubygemslegacyproxy.SetupGated,
		rubygemsproxy.SetupGated,
		swiftgroup.SetupGated,
		swiftproxy.SetupGated,
		terraformgroup.SetupGated,
		terraformhosted.SetupGated,
		terraformproxy.SetupGated,
		yumgroup.SetupGated,
		yumhosted.SetupGated,
		yumproxy.SetupGated,
		rule.SetupGated,
		oauth2.SetupGated,
		realms.SetupGated,
		role.SetupGated,
		saml.SetupGated,
		ssltruststore.SetupGated,
		ssrfprotection.SetupGated,
		user.SetupGated,
		usertokens.SetupGated,
		anonymousaccess.SetupGated,
		confighttp.SetupGated,
		configldapconnection.SetupGated,
		configmail.SetupGated,
		configproductlicense.SetupGated,
		iqconnection.SetupGated,
		blobstorecompact.SetupGated,
		licenseexpirationnotification.SetupGated,
		malwareremediator.SetupGated,
		repaircreatebrowsenodes.SetupGated,
		repositorydockergc.SetupGated,
		repositorydockeruploadpurge.SetupGated,
		repositorymavenremovesnapshots.SetupGated,
		repositorypurgeunused.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupWebhookWithManager registers conversion webhooks for all resource kinds in the group.
func SetupWebhookWithManager(mgr ctrl.Manager) error {
	for _, setup := range []func(ctrl.Manager) error{
		acs.SetupWebhookWithManager,
		file.SetupWebhookWithManager,
		gcs.SetupWebhookWithManager,
		group.SetupWebhookWithManager,
		s3.SetupWebhookWithManager,
		audit.SetupWebhookWithManager,
		baseurl.SetupWebhookWithManager,
		customs3regions.SetupWebhookWithManager,
		defaultrole.SetupWebhookWithManager,
		firewallauditandquarantine.SetupWebhookWithManager,
		healthcheck.SetupWebhookWithManager,
		outreachmanagement.SetupWebhookWithManager,
		rutauth.SetupWebhookWithManager,
		storagesettings.SetupWebhookWithManager,
		uibranding.SetupWebhookWithManager,
		uisettings.SetupWebhookWithManager,
		webhookglobal.SetupWebhookWithManager,
		webhookrepository.SetupWebhookWithManager,
		policy.SetupWebhookWithManager,
		selector.SetupWebhookWithManager,
		application.SetupWebhookWithManager,
		repositoryadmin.SetupWebhookWithManager,
		repositorycontentselector.SetupWebhookWithManager,
		repositoryview.SetupWebhookWithManager,
		script.SetupWebhookWithManager,
		wildcard.SetupWebhookWithManager,
		providerconfig.SetupWebhookWithManager,
		alpinegroup.SetupWebhookWithManager,
		alpinehosted.SetupWebhookWithManager,
		alpineproxy.SetupWebhookWithManager,
		ansiblegalaxygroup.SetupWebhookWithManager,
		ansiblegalaxyhosted.SetupWebhookWithManager,
		ansiblegalaxyproxy.SetupWebhookWithManager,
		apthosted.SetupWebhookWithManager,
		aptproxy.SetupWebhookWithManager,
		cargogroup.SetupWebhookWithManager,
		cargohosted.SetupWebhookWithManager,
		cargoproxy.SetupWebhookWithManager,
		cocoapodsproxy.SetupWebhookWithManager,
		composerproxy.SetupWebhookWithManager,
		conangroup.SetupWebhookWithManager,
		conanhosted.SetupWebhookWithManager,
		conanproxy.SetupWebhookWithManager,
		condaproxy.SetupWebhookWithManager,
		dockergroup.SetupWebhookWithManager,
		dockerhosted.SetupWebhookWithManager,
		dockerproxy.SetupWebhookWithManager,
		gitlfshosted.SetupWebhookWithManager,
		gogroup.SetupWebhookWithManager,
		gohosted.SetupWebhookWithManager,
		goproxy.SetupWebhookWithManager,
		helmgroup.SetupWebhookWithManager,
		helmhosted.SetupWebhookWithManager,
		helmproxy.SetupWebhookWithManager,
		huggingfaceproxy.SetupWebhookWithManager,
		maven2group.SetupWebhookWithManager,
		maven2hosted.SetupWebhookWithManager,
		maven2proxy.SetupWebhookWithManager,
		mavengroup.SetupWebhookWithManager,
		mavenhosted.SetupWebhookWithManager,
		mavenproxy.SetupWebhookWithManager,
		npmgroup.SetupWebhookWithManager,
		npmhosted.SetupWebhookWithManager,
		npmproxy.SetupWebhookWithManager,
		nugetgroup.SetupWebhookWithManager,
		nugethosted.SetupWebhookWithManager,
		nugetproxy.SetupWebhookWithManager,
		ocigroup.SetupWebhookWithManager,
		ocihosted.SetupWebhookWithManager,
		ociproxy.SetupWebhookWithManager,
		p2proxy.SetupWebhookWithManager,
		pubgroup.SetupWebhookWithManager,
		pubhosted.SetupWebhookWithManager,
		pubproxy.SetupWebhookWithManager,
		pypigroup.SetupWebhookWithManager,
		pypihosted.SetupWebhookWithManager,
		pypiproxy.SetupWebhookWithManager,
		rawgroup.SetupWebhookWithManager,
		rawhosted.SetupWebhookWithManager,
		rawproxy.SetupWebhookWithManager,
		rgroup.SetupWebhookWithManager,
		rhosted.SetupWebhookWithManager,
		rproxy.SetupWebhookWithManager,
		rubygemsgroup.SetupWebhookWithManager,
		rubygemshosted.SetupWebhookWithManager,
		rubygemslegacygroup.SetupWebhookWithManager,
		rubygemslegacyhosted.SetupWebhookWithManager,
		rubygemslegacyproxy.SetupWebhookWithManager,
		rubygemsproxy.SetupWebhookWithManager,
		swiftgroup.SetupWebhookWithManager,
		swiftproxy.SetupWebhookWithManager,
		terraformgroup.SetupWebhookWithManager,
		terraformhosted.SetupWebhookWithManager,
		terraformproxy.SetupWebhookWithManager,
		yumgroup.SetupWebhookWithManager,
		yumhosted.SetupWebhookWithManager,
		yumproxy.SetupWebhookWithManager,
		rule.SetupWebhookWithManager,
		oauth2.SetupWebhookWithManager,
		realms.SetupWebhookWithManager,
		role.SetupWebhookWithManager,
		saml.SetupWebhookWithManager,
		ssltruststore.SetupWebhookWithManager,
		ssrfprotection.SetupWebhookWithManager,
		user.SetupWebhookWithManager,
		usertokens.SetupWebhookWithManager,
		anonymousaccess.SetupWebhookWithManager,
		confighttp.SetupWebhookWithManager,
		configldapconnection.SetupWebhookWithManager,
		configmail.SetupWebhookWithManager,
		configproductlicense.SetupWebhookWithManager,
		iqconnection.SetupWebhookWithManager,
		blobstorecompact.SetupWebhookWithManager,
		licenseexpirationnotification.SetupWebhookWithManager,
		malwareremediator.SetupWebhookWithManager,
		repaircreatebrowsenodes.SetupWebhookWithManager,
		repositorydockergc.SetupWebhookWithManager,
		repositorydockeruploadpurge.SetupWebhookWithManager,
		repositorymavenremovesnapshots.SetupWebhookWithManager,
		repositorypurgeunused.SetupWebhookWithManager,
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}
