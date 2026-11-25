# Rollback-Capable Disjoint Set Union by Rank Without Compression (`union_find_rollback_d60`)

> Category: **Disjoint Sets** | Language: **Go 1.22+**

## Overview

Tracks DSU parent and rank mutations on an undo stack so dynamic graph connectivity queries can rewind cleanly across DFS/offline timelines.

## Complexity Guarantees

| Metric | Asymptotic Bound |
| :--- | :--- |
| **Time Complexity** | `O(log N) union/find + O(1) per-merge rollback snapshot` |
| **Space Complexity** | `O(N + U)` |

## Mathematical & Safety Invariants

- Rewinding to a saved checkpoint restores exact component count and connectivity.

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
go test -v ./src/modules/union_find_rollback_d60
```

---
© 2026 Igor Kan
