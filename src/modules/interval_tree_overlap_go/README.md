# Augmented Interval Search Tree for All-Overlap Queries (`interval_tree_overlap_go`)

> Category: **Geometric Trees** | Language: **Go 1.22+**

## Overview

Stores closed intervals [Start, End] sorted by Start with augmented subtree MaxEnd pruning to collect every interval intersecting a query range [lo, hi].

## Complexity Guarantees

| Metric | Asymptotic Bound |
| :--- | :--- |
| **Time Complexity** | `O(N log N) balanced sorted tree build + O(log N + K) overlap search` |
| **Space Complexity** | `O(N)` |

## Mathematical & Safety Invariants

- Returned intervals are strictly those with max(Start, lo) <= min(End, hi).

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
go test -v ./src/modules/interval_tree_overlap_go
```

---
© 2026 Igor Kan
