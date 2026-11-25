# HyperLogLog Distinct Cardinality Estimator with Linear Counting (`hyperloglog_card_d60`)

> Category: **Streaming Sketches** | Language: **Go 1.22+**

## Overview

Maintains M=64 leading-zero rank registers via splitmix64 bit mixing and harmonic-mean aggregation with small-range linear-counting correction.

## Complexity Guarantees

| Metric | Asymptotic Bound |
| :--- | :--- |
| **Time Complexity** | `O(1) register update and O(M) harmonic-mean estimation` |
| **Space Complexity** | `O(2^p) byte registers` |

## Mathematical & Safety Invariants

- Duplicate insertions are idempotent on every register.

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
go test -v ./src/modules/hyperloglog_card_d60
```

---
© 2026 Igor Kan
