# O(1) Doubly-Linked Frequency-Bucket Least Frequently Used Cache (`lfu_cache_o1_buckets`)

> Category: **Cache Eviction Policies** | Language: **Go 1.22+**

## Overview

Maintains O(1) cache operations by storing entries inside per-frequency doubly-linked LRU lists indexed by freq -> list head/tail and tracking global minFreq on promotion/eviction.

## Complexity Guarantees

| Metric | Asymptotic Bound |
| :--- | :--- |
| **Time Complexity** | `O(1) Get, Put, and Min-Freq LRU Tie-Break Eviction` |
| **Space Complexity** | `O(Capacity) node & frequency map` |

## Mathematical & Safety Invariants

- Total item count never exceeds capacity.
- minFreq strictly equals the minimum frequency among all currently resident entries when non-empty.

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
go test -v ./src/modules/lfu_cache_o1_buckets
```

---
© 2026 Igor Kan
