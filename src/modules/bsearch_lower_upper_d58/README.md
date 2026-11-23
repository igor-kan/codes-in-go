# LowerBound & UpperBound Frequency Span Finder (`bsearch_lower_upper_d58`)

> Category: **Binary Search** | Language: **Go 1.22+**

## Overview

Computes half-open [lower, upper) bounds over ascending integer arrays cleanly avoiding integer midpoint overflow.

## Complexity Guarantees

| Metric | Asymptotic Bound |
| :--- | :--- |
| **Time Complexity** | `O(log N)` |
| **Space Complexity** | `O(1)` |

## Mathematical & Safety Invariants

- Target frequency strictly equals upperBound - lowerBound across all keys.

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
go test -v ./src/modules/bsearch_lower_upper_d58
```

---
© 2026 Igor Kan
