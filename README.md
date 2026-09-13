# provider-nexus

A [Crossplane](https://crossplane.io) provider for
[Sonatype Nexus Repository](https://www.sonatype.com/products/sonatype-nexus-repository),
generated with [Upjet](https://github.com/crossplane/upjet) from the community
Terraform provider
[`sonatype-nexus-community/terraform-provider-sonatyperepo`](https://github.com/sonatype-nexus-community/terraform-provider-sonatyperepo).

Every resource runs **in-process** through terraform-plugin-framework — the
provider image contains no Terraform CLI. The Terraform provider is embedded
via the fork
[`tinyorbitvn/terraform-provider-sonatyperepo`](https://github.com/tinyorbitvn/terraform-provider-sonatyperepo),
required directly as `github.com/tinyorbitvn/terraform-provider-sonatyperepo`
v1.19.0-xp.2 (fork branch `xp`; `xpprovider` branch = upstream PR material —
it only adds a public `xpprovider` package re-exporting the provider
constructor; we are proposing the shim upstream; until it is merged **and**
upstream adopts a dotted Go module path, the fork stays required).

## Install

```yaml
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-nexus
spec:
  package: ghcr.io/tinyorbitvn/provider-nexus:v0.1.0
```

Crossplane v2 is required (both cluster-scoped `*.nexus.tinyorbit.vn` and
namespaced `*.nexus.m.tinyorbit.vn` APIs are generated).

## Configure

```yaml
apiVersion: nexus.m.tinyorbit.vn/v1beta1
kind: ProviderConfig
metadata:
  name: default
  namespace: nexus
spec:
  url: http://nexus.nexus.svc:8081
  username: admin
  credentials:
    source: Secret
    secretRef:
      name: nexus-admin
      namespace: nexus
      key: NEXUS_ADMIN_PASSWORD   # raw password, not JSON
```

The user needs at least `nx-metrics-all` (the Terraform provider checks
cluster status on start) plus the privileges of whatever you manage.

## Resources

All 119 resources of terraform-provider-sonatyperepo 1.19.0 are generated.
API groups: `repository` (70 formats × hosted/proxy/group), `blobstore`,
`security` (realms, ssrfprotection, users, roles, …), `system`, `capability`,
`privilege`, `cleanup`, `content`, `routing`, `task`. See `package/crds/` and
`examples-generated/`.

External names (see `config/external_name.go` for the full reasoning):

- repositories, blob stores, privileges, routing rules, content selectors and
  cleanup policies: identified by their Nexus **name**, read back from state
  (`crossplane.io/external-name: <name>` adopts an existing one; defaults to
  `metadata.name`, so names that are not valid Kubernetes object names —
  upper case, `_`, … — need the annotation);
- roles: identified by the role **`id`** (annotation; defaults to
  `metadata.name`), `spec.forProvider.name` is the display name;
- users: identified by **`user_id`** (annotation); `source` is read-only and
  not part of the identity;
- capabilities, `Realms`, SSL truststore, tasks and other kinds whose
  Terraform schema has an `id`: ID assigned by Nexus, filled in after the
  first reconcile;
- the 9 one-per-server singletons (`AnonymousAccess`, `SsrfProtection`,
  `Oauth2`, `Saml`, `UserTokens`, `ConfigHttp`, `ConfigMail`,
  `ConfigProductLicense`, `IqConnection`): external name is always the
  constant `default`.

## Known limitations

1. **Observe-only Docker repositories** — `DockerHosted`/`DockerProxy`/`DockerGroup`
   with `managementPolicies: ["Observe"]` must populate the required blocks
   (`storage`, `docker`, and `proxy`/`group` where applicable) in
   `spec.forProvider`: the upstream Terraform models use non-pointer structs
   and reject a null config ("Received null value…"). See the comments in
   `examples/nexus/observe/repositories.yaml`. Create/Update are unaffected.
2. **Deleting a singleton MR writes to Nexus** — the Terraform provider's
   Delete for singletons is not a no-op: `AnonymousAccess` → sets
   `enabled=false`; `ConfigHttp` → resets HTTP settings; `SsrfProtection` →
   overwrites; `UserTokens` → disables; `ConfigMail` → deletes the mail
   config (see the fork's `internal/provider/system/*_resource.go`). A GitOps
   prune therefore reconfigures the live
   server — set `spec.deletionPolicy: Orphan` on singleton MRs unless that is
   intended.
3. **`secretRef.namespace`** is required by the CRD in both scopes; for the
   namespaced `ProviderConfig` it must equal the ProviderConfig's own
   namespace (the provider overrides it with the MR namespace).

## Develop

```bash
make submodules     # once
make generate       # schema from the Terraform registry + docs + code generation
make test
make build          # binary + image for the host platform
make local-deploy   # kind cluster with Crossplane and this provider
make run            # controller out-of-cluster against $KUBECONFIG
```

### Bumping the Terraform provider

1. Rebase the fork branch `xpprovider` onto the new upstream tag, tag it
   `v<upstream>-xp.1`, push (upstream PR material only — no dotted module
   path here).
2. Rebase fork branch `xp` onto the updated `xpprovider`, tag it
   `v<upstream>-xp.2`, push.
3. Update `TERRAFORM_PROVIDER_VERSION` in `Makefile` and the `require` line
   (`github.com/tinyorbitvn/terraform-provider-sonatyperepo`) in `go.mod`;
   `go mod tidy`.
4. `make generate && make check-diff`; review CRD changes reported by
   `make report-breaking-changes` (crddiff) in CI.
5. Tag `vX.Y.Z`; the `Publish` workflow pushes
   `ghcr.io/tinyorbitvn/provider-nexus:vX.Y.Z`.
