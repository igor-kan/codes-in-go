# Prefix-Doubling Suffix Array & Kasai O(N) LCP Array Constructor (`suffix_array_kasai_d60`)

> Category: **String Indexing** | Language: **Go 1.22+**

## Overview

Constructs lexicographically sorted suffix indices and adjacent Longest Common Prefix heights in Go to count distinct substrings in O(N).

## Complexity Guarantees

| Metric | Asymptotic Bound |
| :--- | :--- |
| **Time Complexity** | `O(N log N) suffix sort + O(N) Kasai LCP scan` |
| **Space Complexity** | `O(N)` |

## Mathematical & Safety Invariants

- Number of distinct non-empty substrings equals N*(N+1)/2 - sum(LCP).

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
go test -v ./src/modules/suffix_array_kasai_d60
```

---
© 2026 Igor Kan
