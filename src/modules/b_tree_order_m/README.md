# Ordered In-Memory B-Tree Dictionary with Node Splitting (`b_tree_order_m`)

> Category: **Ordered Search Trees** | Language: **Go 1.22+**

## Overview

Implements an ordered B-Tree dictionary with configurable minimum degree t >= 2 supporting preemptive child splitting on downward insertion, exact key search, and sorted in-order range enumeration.

## Complexity Guarantees

| Metric | Asymptotic Bound |
| :--- | :--- |
| **Time Complexity** | `O(log N) search/insert, O(N) ordered scan` |
| **Space Complexity** | `O(N)` |

## Mathematical & Safety Invariants

- Every non-root node holds keys k satisfying t - 1 <= len(keys) <= 2*t - 1.
- In-order traversal yields keys in strictly ascending sorted sequence matching a Go reference map.

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
go test -v ./src/modules/b_tree_order_m
```

---
© 2026 Igor Kan
