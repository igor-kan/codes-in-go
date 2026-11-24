# Aho-Corasick Multi-Keyword Failure Automaton (`aho_multi_scan_d59`)

> Category: **Multi-Pattern Search** | Language: **Go 1.22+**

## Overview

Compiles arbitrary lowercase keyword sets into an Aho-Corasick failure DFA and counts simultaneous overlapping hits across input text cleanly.

## Complexity Guarantees

| Metric | Asymptotic Bound |
| :--- | :--- |
| **Time Complexity** | `O(|T| + sum(|P|) * 26)` |
| **Space Complexity** | `O(states * 26)` |

## Mathematical & Safety Invariants

- Exact multi-keyword occurrence sum matches individual KMP counts summed cleanly.

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
go test -v ./src/modules/aho_multi_scan_d59
```

---
© 2026 Igor Kan
