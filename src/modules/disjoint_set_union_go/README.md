# Path-Compressing & Rank-Balanced Disjoint Set Forest (`disjoint_set_union_go`)

> Category: **Graph Connectivity** | Language: **Go 1.22+**

## Overview

Maintains dynamic partition connectivity over N nodes via path compression and union by subtree size with live component counting.

## Complexity Guarantees

| Metric | Asymptotic Bound |
| :--- | :--- |
| **Time Complexity** | `O(alpha(N)) amortized Find and Union` |
| **Space Complexity** | `O(N)` |

## Mathematical & Safety Invariants

- Component count starts at N and decreases by 1 strictly upon merging two distinct sets.

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
go test -v ./src/modules/disjoint_set_union_go
```

---
© 2026 Igor Kan
