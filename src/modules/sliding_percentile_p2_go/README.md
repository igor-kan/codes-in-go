# Exact Sliding Window Order-Statistic Tracker in Go (`sliding_percentile_p2_go`)

> Category: **Streaming Telemetry** | Language: **Go 1.22+**

## Overview

Maintains the last W numeric observations inside a circular buffer and computes exact p-quantiles over the active window.

## Complexity Guarantees

| Metric | Asymptotic Bound |
| :--- | :--- |
| **Time Complexity** | `O(W log W) bounded window ring + exact percentile extraction` |
| **Space Complexity** | `O(W)` |

## Mathematical & Safety Invariants

- Median of symmetric window equals exact midpoint.

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
go test -v ./src/modules/sliding_percentile_p2_go
```

---
© 2026 Igor Kan
