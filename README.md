# Switch Common Library

OpenPak's Switch Common Library.

| | |
| --- | --- |
| Network | Nintendo Switch |
| Language | Go module `github.com/openpak/nx-common` |
| Used by | [`nx-shop`](https://github.com/openpak/nx-shop), [`nx-penne`](https://github.com/openpak/nx-penne), [`nx-pctl`](https://github.com/openpak/nx-pctl) |
| Licence | AGPL-3.0-only |

## Packages

| Package | What |
| --- | --- |
| `nxtls` | TLS termination: fixed multi-SAN cert and on-demand certs from the OpenPak CA |
| `nxjwt` | RS256 sign/verify with the BAAS key, key loading, id helpers |
| `nxhttp` | JSON reply, host helper, full-request dump, 421 for a foreign host |
| `internalkey` | `X-Internal-Key` check (strict and lenient) and a guard |
| `nxapi` | Internal identity/penne contract (types, paths) and its client |

## Use

```go
import "github.com/openpak/nx-common/nxjwt"
```

Each service uses `replace github.com/openpak/nx-common => ../nx-common` and commits `vendor/`; details in [deploy](docs/deploy.md).

## Docs

- [Protocol](docs/protocol.md)
- [Evidence](docs/evidence.md)
- [Deploy](docs/deploy.md)
- [Notes](docs/notes.md)

## Licence

AGPL-3.0-only. See [LICENSE](LICENSE).
