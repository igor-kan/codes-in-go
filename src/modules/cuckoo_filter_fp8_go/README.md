# 4-Way Bucket 8-Bit Fingerprint Cuckoo Filter in Go (`cuckoo_filter_fp8_go`)

> Category: **Probabilistic Filters** | Language: **Go 1.22+**

## Overview

Implements Fan et al.'s Cuckoo Filter storing 8-bit non-zero fingerprints in 4-slot buckets with partial-key cuckoo hashing i2 = i1 ^ hash(fp).

## Complexity Guarantees

| Metric | Asymptotic Bound |
| :--- | :--- |
| **Time Complexity** | `O(1) dual-bucket lookup, insert, and delete` |
| **Space Complexity** | `O(B * 4) byte slots` |

## Mathematical & Safety Invariants

- Inserted keys always test positive and can be deleted cleanly.

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
go test -v ./src/modules/cuckoo_filter_fp8_go
```

---
© 2026 Igor Kan
