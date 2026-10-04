# nx-common

The small Go module the split-out Switch services share (nx-shop, nx-penne, nx-pctl; nx-baas
implements the same contract without importing it yet). Kept minimal on purpose: anything
that is one service's business stays in that service.

| package       | what                                                                    |
|---------------|-------------------------------------------------------------------------|
| `nxtls`       | TLS termination: fixed multi-SAN cert + on-demand certs from the OpenPak CA |
| `nxjwt`       | RS256 sign/verify with the BAAS key, key loading, id helpers            |
| `nxhttp`      | the JSON reply, host helper, full-request dump, 421 for a foreign host  |
| `internalkey` | the `X-Internal-Key` check (strict and lenient) and a guard             |
| `nxapi`       | the internal identity/penne contract (types, paths) and its client      |

The contract itself is written up in nx-baas `docs/service-split.md`.

## How services consume it

No remote yet, so each service has `replace github.com/openpak/nx-common => ../nx-common` and
commits `vendor/` (`go mod vendor`). The image build then needs nothing but the service's own
repository, which keeps the release workflow identical to nx-baas's. After changing this
module, run `go mod vendor` in each service and commit the result there. Once this repository
has a remote and tags, the `replace` line goes and `vendor/` can too.
