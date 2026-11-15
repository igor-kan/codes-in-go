# Lamport-Mattern Vector Clocks & Concurrent Conflict Detector (`vector_clock_causality`)

> Category: **Distributed Systems** | Language: **Go 1.22+**

## Overview

Captures exact partial causal ordering across asynchronous processes using componentwise logical counters, distinguishing Happened-Before, Equal, After, and Concurrent events.

## Complexity Guarantees

| Metric | Asymptotic Bound |
| :--- | :--- |
| **Time Complexity** | `O(N) compare/merge for N process entries` |
| **Space Complexity** | `O(N) uint64 clock slice` |

## Mathematical & Safety Invariants

- Causal transitivity: if A -> B and B -> C then A -> C.
- Compare relation is antisymmetric between Before and After, and symmetric for Concurrent.

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
go test -v ./src/modules/vector_clock_causality
```

---
© 2026 Igor Kan
