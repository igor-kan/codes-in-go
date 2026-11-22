# Sweep-Line Max Concurrent Interval Overlap Counter (`interval_tree_overlap_d57`)

> Category: **Interval Algorithms** | Language: **Go 1.22+**

## Overview

Sweeps sorted +1/-1 half-open interval boundary events [start, end) to compute exact peak simultaneous overlap depth.

## Complexity Guarantees

| Metric | Asymptotic Bound |
| :--- | :--- |
| **Time Complexity** | `O(N log N) sweep line` |
| **Space Complexity** | `O(N)` |

## Mathematical & Safety Invariants

- Peak overlap equals maximum concurrency across all interval boundaries cleanly.

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
go test -v ./src/modules/interval_tree_overlap_d57
```

---
© 2026 Igor Kan
