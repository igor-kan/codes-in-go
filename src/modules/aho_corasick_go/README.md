# Immutable Multi-Pattern Aho-Corasick Byte Automaton (`aho_corasick_go`)

> Category: **String Algorithms** | Language: **Go 1.22+**

## Overview

Builds a failure-linked DFA automaton over exact multi-pattern byte dictionaries for thread-safe concurrent log/corpus match counting.

## Complexity Guarantees

| Metric | Asymptotic Bound |
| :--- | :--- |
| **Time Complexity** | `O(sum(|P|) + |T|) DFA construction & parallel scanning` |
| **Space Complexity** | `O(States * 256)` |

## Mathematical & Safety Invariants

- Match counts across arbitrary texts agree identically with brute substring scan.

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
go test -v ./src/modules/aho_corasick_go
```

---
© 2026 Igor Kan
