# Last-Write-Wins Element-Set (LWW-Element-Set) Conflict-Free Replicated Data Type (`crdt_lww_element_set`)

> Category: **Distributed Systems & CRDTs** | Language: **Go 1.22+**

## Overview

Implements a state-based LWW-Element-Set CRDT with bias-resolved equal-timestamp tie-breaking, associative/commutative/idempotent join semilattice merge, and concurrent replica convergence verification.

## Complexity Guarantees

| Metric | Asymptotic Bound |
| :--- | :--- |
| **Time Complexity** | `O(1) Add/Remove, O(N) Merge` |
| **Space Complexity** | `O(Keys)` |

## Mathematical & Safety Invariants

- Merge join satisfies commutativity: Merge(A, B) == Merge(B, A).
- Merge join satisfies idempotence (Merge(A, A) == A) and associativity.
- An element belongs to the active set iff max_add_ts >= max_remove_ts (with remove/add bias on ties).

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
go test -v ./src/modules/crdt_lww_element_set
```

---
© 2026 Igor Kan
