# Knuth-Morris-Pratt Exact Offset Matcher (`kmp_stream_go_d59`)

> Category: **String Matching** | Language: **Go 1.22+**

## Overview

Computes exact prefix-function border shifts over pattern bytes and extracts all matching zero-based text indices cleanly.

## Complexity Guarantees

| Metric | Asymptotic Bound |
| :--- | :--- |
| **Time Complexity** | `O(|T| + |P|)` |
| **Space Complexity** | `O(|P|)` |

## Mathematical & Safety Invariants

- Every emitted index satisfies exact substring equality against pat.

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
go test -v ./src/modules/kmp_stream_go_d59
```

---
© 2026 Igor Kan
