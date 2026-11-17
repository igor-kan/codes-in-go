# Concurrent Conservative Count-Min Sketch Heavy Hitter Tracker (`count_min_heavy_go`)

> Category: **Streaming Sketches** | Language: **Go 1.22+**

## Overview

Combines a conservative-update Count-Min Sketch with a bounded mutex-guarded min-heap to estimate streaming item frequencies and extract heavy hitters without storing full histograms.

## Complexity Guarantees

| Metric | Asymptotic Bound |
| :--- | :--- |
| **Time Complexity** | `O(D) conservative hash update and min-estimate query` |
| **Space Complexity** | `O(D * W)` |

## Mathematical & Safety Invariants

- Count-Min estimate is never below true frequency (est(x) >= freq(x)).
- Top heavy-hitter items rank above low-frequency background noise.

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
go test -v ./src/modules/count_min_heavy_go
```

---
© 2026 Igor Kan
