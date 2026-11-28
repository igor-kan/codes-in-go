# Eertree Distinct Palindromic Substring Counter (`eertree_palindrome_go_d61`)

> Category: **String Indexing** | Language: **Go 1.22+**

## Overview

Builds a palindromic tree with imaginary length -1 and length 0 roots, suffix links, and transition maps so every distinct palindromic substring is counted once.

## Complexity Guarantees

| Metric | Asymptotic Bound |
| :--- | :--- |
| **Time Complexity** | `O(N) amortized palindromic-tree construction` |
| **Space Complexity** | `O(N) nodes and suffix links` |

## Mathematical & Safety Invariants

- The number of nodes minus two equals the number of distinct non-empty palindromic substrings.

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
go test -v ./src/modules/eertree_palindrome_go_d61
```

---
© 2026 Igor Kan
