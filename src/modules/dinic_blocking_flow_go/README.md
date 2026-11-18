# Dinic Level-Graph Blocking-Flow Network Solver in Go (`dinic_blocking_flow_go`)

> Category: **Network Flow** | Language: **Go 1.22+**

## Overview

Computes maximum s-t network flow using Dinic's two-phase level graph BFS and current-arc pointer DFS augmentation.

## Complexity Guarantees

| Metric | Asymptotic Bound |
| :--- | :--- |
| **Time Complexity** | `O(V^2 * E) BFS level graph + current-arc DFS blocking flow` |
| **Space Complexity** | `O(V + E)` |

## Mathematical & Safety Invariants

- Total augmented flow equals minimum s-t cut capacity.

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
go test -v ./src/modules/dinic_blocking_flow_go
```

---
© 2026 Igor Kan
