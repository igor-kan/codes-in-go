# Concurrent In-Flight Request Coalescer (Singleflight Group) (`singleflight_dedup_group`)

> Category: **Concurrency & Cache Stampede Prevention** | Language: **Go 1.22+**

## Overview

Coalesces concurrent calls sharing identical string keys so only one leader executes the underlying expensive function while waiting callers receive the shared result and Shared=true indicator.

## Complexity Guarantees

| Metric | Asymptotic Bound |
| :--- | :--- |
| **Time Complexity** | `O(1) per caller lookup/join` |
| **Space Complexity** | `O(Active Keys)` |

## Mathematical & Safety Invariants

- Concurrent burst callers on the same key invoke underlying loader exactly once.
- Forget deletes key mapping without aborting already-running leader computations.

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
go test -v ./src/modules/singleflight_dedup_group
```

---
© 2026 Igor Kan
