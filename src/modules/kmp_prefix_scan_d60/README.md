# Knuth-Morris-Pratt Prefix Automaton & Overlapping Matcher (`kmp_prefix_scan_d60`)

> Category: **String Matching** | Language: **Go 1.22+**

## Overview

Computes the Knuth-Morris-Pratt longest-proper-prefix-suffix (pi) array and finds every overlapping match offset without text cursor backtracking.

## Complexity Guarantees

| Metric | Asymptotic Bound |
| :--- | :--- |
| **Time Complexity** | `O(N + M) zero-backtracking text scan` |
| **Space Complexity** | `O(M) LPS table` |

## Mathematical & Safety Invariants

- All reported match start indices satisfy text[idx:idx+len(pat)] == pat.

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
go test -v ./src/modules/kmp_prefix_scan_d60
```

---
© 2026 Igor Kan
