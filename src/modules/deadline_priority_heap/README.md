# Indexed Cancelable Deadline Timer Min-Heap in Go (`deadline_priority_heap`)

> Category: **Scheduling & Concurrency** | Language: **Go 1.22+**

## Overview

Implements an indexed binary min-heap keyed by nanosecond deadline supporting O(log N) deadline rescheduling and task cancellation via live heap-index tracking.

## Complexity Guarantees

| Metric | Asymptotic Bound |
| :--- | :--- |
| **Time Complexity** | `O(log N) push, update-deadline, cancel, and expire-ready drain` |
| **Space Complexity** | `O(N) heap + index map` |

## Mathematical & Safety Invariants

- Heap property deadline[parent] <= deadline[child] holds after every push, cancel, update, and pop.
- DrainExpired emits all ready items in monotonically non-decreasing deadline order.

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
go test -v ./src/modules/deadline_priority_heap
```

---
© 2026 Igor Kan
