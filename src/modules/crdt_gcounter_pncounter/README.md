# State-Based G-Counter & PN-Counter Conflict-Free Replicated Data Types (`crdt_gcounter_pncounter`)

> Category: **Distributed Systems** | Language: **Go 1.22+**

## Overview

Implements convergent state-based grow-only (GCounter) and positive-negative (PNCounter) CvRDTs forming join-semilattices guaranteed to converge identically under arbitrary duplication and reordering.

## Complexity Guarantees

| Metric | Asymptotic Bound |
| :--- | :--- |
| **Time Complexity** | `O(R) Value/Merge over R replica slots` |
| **Space Complexity** | `O(R) per counter` |

## Mathematical & Safety Invariants

- Merge is commutative: merge(a, b) == merge(b, a).
- Merge is associative and idempotent: merge(a, a) == a.

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
go test -v ./src/modules/crdt_gcounter_pncounter
```

---
© 2026 Igor Kan
