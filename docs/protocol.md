# Protocol: nx-common

The contract itself is written up in nx-baas `docs/service-split.md`.

Consumers: nx-shop, nx-penne and nx-pctl import it; nx-baas implements the same contract without
importing it yet. The module is kept minimal on purpose: anything that is one service's business
stays in that service.
