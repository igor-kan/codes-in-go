# Writer-Fair FIFO Ticketed Reader-Writer Mutex (`rw_fair_rwlock_queue`)

> Category: **Synchronization Primitives** | Language: **Go 1.22+**

## Overview

Implements a starvation-free Phase-Fair Reader-Writer lock using sync.Mutex + sync.Cond FIFO turnstile counters preventing infinite reader streams from starving queued writers.

## Complexity Guarantees

| Metric | Asymptotic Bound |
| :--- | :--- |
| **Time Complexity** | `O(1) ticket acquire/release` |
| **Space Complexity** | `O(1)` |

## Mathematical & Safety Invariants

- At most one writer holds exclusive ownership, never concurrently with active readers.
- Once a writer requests Lock(), subsequent RLock() requests queue behind it in FIFO fairness.

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
go test -v ./src/modules/rw_fair_rwlock_queue
```

---
© 2026 Igor Kan
