# Sorted-Page B+Tree Leaf Chain with Range Cursor (`bplus_leaf_cursor_go`)

> Category: **Storage Indexes** | Language: **Go 1.22+**

## Overview

Maintains sorted key-value entries partitioned across linked leaf pages with O(log N) binary search entry point and forward leaf-chain cursor iteration.

## Complexity Guarantees

| Metric | Asymptotic Bound |
| :--- | :--- |
| **Time Complexity** | `O(log N + K) binary-searched leaf chain range scan` |
| **Space Complexity** | `O(N)` |

## Mathematical & Safety Invariants

- Range scan [lo, hi] emits keys in strictly ascending order.

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
go test -v ./src/modules/bplus_leaf_cursor_go
```

---
© 2026 Igor Kan
