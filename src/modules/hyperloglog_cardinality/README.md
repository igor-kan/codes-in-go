# HyperLogLog Probabilistic Cardinality Estimator with LinearCounting (`hyperloglog_cardinality`)

> Category: **Streaming Algorithms** | Language: **Go 1.22+**

## Overview

Estimates distinct stream cardinality via 64-bit SplitMix hashing, math/bits LeadingZeros64 rank extraction, Flajolet-Fusy alpha_m correction, and small-cardinality LinearCounting.

## Complexity Guarantees

| Metric | Asymptotic Bound |
| :--- | :--- |
| **Time Complexity** | `O(1) Add, O(2^p) Count & Merge` |
| **Space Complexity** | `O(2^p) uint8 registers` |

## Mathematical & Safety Invariants

- Join merge is commutative, associative, and idempotent across replica sketches.
- Repeated observations of identical keys never alter register max values.

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
go test -v ./src/modules/hyperloglog_cardinality
```

---
© 2026 Igor Kan
