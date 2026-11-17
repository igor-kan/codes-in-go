# Mirrokni-Tardos Bounded-Load Consistent Hash Ring (`consistent_bounded_load`)

> Category: **Distributed Systems** | Language: **Go 1.22+**

## Overview

Implements Google's Consistent Hashing with Bounded Loads (cap = ceil(c * totalLoad / numServers)), spilling clockwise around the 64-bit ring when a server hits capacity.

## Complexity Guarantees

| Metric | Asymptotic Bound |
| :--- | :--- |
| **Time Complexity** | `O(log V + probe) clockwise virtual-node scan with load cap` |
| **Space Complexity** | `O(Servers * Replicas)` |

## Mathematical & Safety Invariants

- No active server's assigned load ever exceeds ceil(factor * totalLoad / N).
- Releasing assigned request decrements exact assigned server load.

## Module Architecture

| File | Lifecycle Role |
| :--- | :--- |
| `types.go` | Stage 01 |
| `errors.go` | Stage 02 |
| `interface.go` | Stage 03 |
| `core.go` | Stage 04 |
| `validation.go` | Stage 05 |
| `fast_path.go` | Stage 06 |
| `concurrent.go` | Stage 07 |
| `invariants.go` | Stage 08 |
| `unit_test.go` | Stage 09 |
| `property_test.go` | Stage 10 |
| `benchmark_test.go` | Stage 11 |
| `README.md` | Stage 12 |

## Verification & Execution Example

```text
go test -v ./src/modules/consistent_bounded_load
```

---
© 2026 Igor Kan
