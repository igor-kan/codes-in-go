# Kuhn Augmenting-Path Bipartite Maximum Matching (`bipartite_kuhn_match_d56`)

> Category: **Network Flows & Matching** | Language: **Go 1.22+**

## Overview

Computes exact maximum-cardinality bipartite matching on unweighted bipartite graphs via DFS augmenting paths with epoch-stamped visited slices.

## Complexity Guarantees

| Metric | Asymptotic Bound |
| :--- | :--- |
| **Time Complexity** | `O(V_L * E)` |
| **Space Complexity** | `O(V_L + V_R)` |

## Mathematical & Safety Invariants

- Matched edges form a valid injective bijection with maximum cardinality.

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
go test -v ./src/modules/bipartite_kuhn_match_d56
```

---
© 2026 Igor Kan
