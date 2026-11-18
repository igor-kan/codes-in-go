# Smooth Nginx-Style Interleaved Weighted Round-Robin Balancer (`weighted_round_robin_gcd`)

> Category: **Load Balancing** | Language: **Go 1.22+**

## Overview

Selects backends in smooth interleaved order proportional to static weights without bursty clustering by adding effective_weight to current_weight and subtracting total_weight from the winner.

## Complexity Guarantees

| Metric | Asymptotic Bound |
| :--- | :--- |
| **Time Complexity** | `O(N) current-weight promotion and total-weight subtraction per pick` |
| **Space Complexity** | `O(N)` |

## Mathematical & Safety Invariants

- Over sum(weights) selections, each server i is picked exactly weight[i] times.
- Sum of current_weights remains identically zero after every pick.

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
go test -v ./src/modules/weighted_round_robin_gcd
```

---
© 2026 Igor Kan
