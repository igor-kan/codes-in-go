# Linear-Time Knuth-Morris-Pratt Prefix Function Stream Scanner (`knuth_morris_pratt_go`)

> Category: **String Matching** | Language: **Go 1.22+**

## Overview

Computes the KMP longest-proper-prefix-suffix array pi[] in O(M) time and emits all starting indices of pattern occurrences inside arbitrary texts in O(N).

## Complexity Guarantees

| Metric | Asymptotic Bound |
| :--- | :--- |
| **Time Complexity** | `O(N + M) zero-backtrack prefix-function DFA scan` |
| **Space Complexity** | `O(M)` |

## Mathematical & Safety Invariants

- Every emitted start offset matches exact string equality.

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
go test -v ./src/modules/knuth_morris_pratt_go
```

---
© 2026 Igor Kan
