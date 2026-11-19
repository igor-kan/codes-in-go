# Binary Indexed Tree Point Update & K-th Order Statistic in Go (`fenwick_range_point_go`)

> Category: **Order-Statistic Structures** | Language: **Go 1.22+**

## Overview

Maintains prefix frequency counts inside a contiguous int64 slice supporting O(log N) point updates, interval sums, and binary-lifted k-th order lookup.

## Complexity Guarantees

| Metric | Asymptotic Bound |
| :--- | :--- |
| **Time Complexity** | `O(log N) BIT add, range sum, and binary-lifted lower bound` |
| **Space Complexity** | `O(N)` |

## Mathematical & Safety Invariants

- Binary-lifted K-th lookup returns exact first index whose cumulative sum >= k.

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
go test -v ./src/modules/fenwick_range_point_go
```

---
© 2026 Igor Kan
