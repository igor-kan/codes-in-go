# Tarjan Lowlink Strongly Connected Components in Go (`tarjan_scc_condense_go`)

> Category: **Graph Algorithms** | Language: **Go 1.22+**

## Overview

Partitions directed graph vertices into maximal strongly connected components in reverse topological order using Tarjan discovery and lowlink numbers.

## Complexity Guarantees

| Metric | Asymptotic Bound |
| :--- | :--- |
| **Time Complexity** | `O(V + E) single-pass DFS lowlink stack` |
| **Space Complexity** | `O(V + E)` |

## Mathematical & Safety Invariants

- Every directed cycle is collapsed into a single component ID.

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
go test -v ./src/modules/tarjan_scc_condense_go
```

---
© 2026 Igor Kan
