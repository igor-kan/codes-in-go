# Path-Compressed Weighted DSU Component Size Tracker (`union_find_size_d57`)

> Category: **Disjoint Sets** | Language: **Go 1.22+**

## Overview

Maintains vertex partition sets with union-by-size and path compression, exposing exact connected component cardinality and max component size.

## Complexity Guarantees

| Metric | Asymptotic Bound |
| :--- | :--- |
| **Time Complexity** | `O(alpha(N)) amortized union and find` |
| **Space Complexity** | `O(N)` |

## Mathematical & Safety Invariants

- Sum of disjoint root sizes strictly equals vertex universe N across all unions.

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
go test -v ./src/modules/union_find_size_d57
```

---
© 2026 Igor Kan
