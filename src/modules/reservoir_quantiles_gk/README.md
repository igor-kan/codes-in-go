# Greenwald-Khanna Epsilon-Approximate Streaming Quantile Summary (`reservoir_quantiles_gk`)

> Category: **Streaming Algorithms** | Language: **Go 1.22+**

## Overview

Computes bounded-error empirical rank quantiles (p50, p95, p99) over continuous float64 streams in sublinear space using Greenwald-Khanna (v, g, delta) tuple insertion and periodic band compression.

## Complexity Guarantees

| Metric | Asymptotic Bound |
| :--- | :--- |
| **Time Complexity** | `O(log(1/eps)) amortized query, O(1/eps * log(eps*N)) space` |
| **Space Complexity** | `Sublinear bounded tuples (v, g, delta)` |

## Mathematical & Safety Invariants

- Estimated q-quantile rank over sorted ground-truth samples falls within eps * N rank tolerance.
- Tracked minimum and maximum always equal exact stream global extrema.

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
go test -v ./src/modules/reservoir_quantiles_gk
```

---
© 2026 Igor Kan
