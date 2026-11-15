# Path-Compressed Potential Weighted Disjoint Set Union (`union_find_weighted`)

> Category: **Graph & Constraint Solvers** | Language: **Go 1.22+**

## Overview

Maintains linear difference equations x[u] - x[v] = w using path compression that accumulates ancestor potential offsets and detects transitive inconsistency contradictions in O(alpha(N)).

## Complexity Guarantees

| Metric | Asymptotic Bound |
| :--- | :--- |
| **Time Complexity** | `O(alpha(N)) inverse-Ackermann amortized per Find/Union` |
| **Space Complexity** | `O(N) parent, rank, and potential weight arrays` |

## Mathematical & Safety Invariants

- After Find(x), parent[x] points directly to root and weight[x] == pot[x] - pot[root].
- Contradictory equation unions return conflict without corrupting state.

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
go test -v ./src/modules/union_find_weighted
```

---
© 2026 Igor Kan
