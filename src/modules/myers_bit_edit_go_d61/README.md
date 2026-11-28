# Myers Bit-Vector Levenshtein Distance Engine (`myers_bit_edit_go_d61`)

> Category: **String Metrics** | Language: **Go 1.22+**

## Overview

Computes exact Levenshtein edit distance by propagating match, positive, and negative bit vectors over the text while keeping a running score.

## Complexity Guarantees

| Metric | Asymptotic Bound |
| :--- | :--- |
| **Time Complexity** | `O(N * ceil(M/64)) word-parallel edit distance` |
| **Space Complexity** | `O(M/64 + 256) bit masks` |

## Mathematical & Safety Invariants

- Distance equals zero exactly when pattern and text are identical and falls back to row DP beyond 64 bytes.

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
go test -v ./src/modules/myers_bit_edit_go_d61
```

---
© 2026 Igor Kan
