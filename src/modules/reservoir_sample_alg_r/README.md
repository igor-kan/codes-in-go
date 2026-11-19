# Deterministic Vitter Algorithm-R Streaming Reservoir Sampler (`reservoir_sample_alg_r`)

> Category: **Streaming Algorithms** | Language: **Go 1.22+**

## Overview

Maintains a uniform random K-element sample over unbounded streams in O(K) memory using SplitMix64 deterministic PRNG replacement probabilities K / i.

## Complexity Guarantees

| Metric | Asymptotic Bound |
| :--- | :--- |
| **Time Complexity** | `O(N) single-pass bounded-memory uniform k-sample` |
| **Space Complexity** | `O(K)` |

## Mathematical & Safety Invariants

- Reservoir holds min(seen, K) elements and deterministic seed replays identical samples.

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
go test -v ./src/modules/reservoir_sample_alg_r
```

---
© 2026 Igor Kan
