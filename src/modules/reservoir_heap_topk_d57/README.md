# Bounded Min-Heap Top-K Value Stream Extractor (`reservoir_heap_topk_d57`)

> Category: **Streaming Selection** | Language: **Go 1.22+**

## Overview

Tracks the exact K largest signed integers across streaming pushes using a container/heap binary min-heap bounded at capacity K.

## Complexity Guarantees

| Metric | Asymptotic Bound |
| :--- | :--- |
| **Time Complexity** | `O(N log K)` |
| **Space Complexity** | `O(K)` |

## Mathematical & Safety Invariants

- Sorted top-K output strictly equals trailing K elements of full sorted input.

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
go test -v ./src/modules/reservoir_heap_topk_d57
```

---
© 2026 Igor Kan
