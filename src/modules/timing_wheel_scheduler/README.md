# Hashed Timing Wheel O(1) Timer Scheduler with Rounds Countdown (`timing_wheel_scheduler`)

> Category: **Concurrent Event Schedulers** | Language: **Go 1.22+**

## Overview

Implements Varghese-Lauck hashed timing wheels where multi-tick future timers record overflow revolution counts in doubly-linked slot buckets without heap percolations.

## Complexity Guarantees

| Metric | Asymptotic Bound |
| :--- | :--- |
| **Time Complexity** | `O(1) Schedule/Cancel, O(1) amortized per tick Advance` |
| **Space Complexity** | `O(Buckets + ActiveTimers)` |

## Mathematical & Safety Invariants

- Every scheduled timer fires at tick == CurrentTick + delayTicks.
- Cancelled timers are detached immediately and never invoked on subsequent advances.

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
go test -v ./src/modules/timing_wheel_scheduler
```

---
© 2026 Igor Kan
