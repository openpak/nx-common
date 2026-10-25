# Deploy: nx-common

## How services consume it

No remote yet, so each service has `replace github.com/openpak/nx-common => ../nx-common` and
commits `vendor/` (`go mod vendor`). The image build then needs nothing but the service's own
repository, which keeps the release workflow identical to nx-baas's. After changing this
module, run `go mod vendor` in each service and commit the result there. Once this repository
has a remote and tags, the `replace` line goes and `vendor/` can too.
