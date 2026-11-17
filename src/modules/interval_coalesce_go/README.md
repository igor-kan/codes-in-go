# Thread-Safe Sorted Disjoint Interval Ledger (`interval_coalesce_go`)

> Category: **Geometric & Interval Engines** | Language: **Go 1.22+**

## Overview

Maintains a sorted list of strictly disjoint closed intervals under thread-safe RWMutex synchronization, merging overlapping or adjacent intervals on insert and checking point coverage.

## Complexity Guarantees

| Metric | Asymptotic Bound |
| :--- | :--- |
| **Time Complexity** | `O(N) sorted insertion/coalesce & O(log N) binary-search point probe` |
| **Space Complexity** | `O(N)` |

## Mathematical & Safety Invariants

- Stored intervals maintain strict ordering iv[i].End + 1 < iv[i+1].Start (disjoint and non-adjacent).
- Total covered integer span matches exact union set cardinality.

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
go test -v ./src/modules/interval_coalesce_go
```

---
© 2026 Igor Kan
