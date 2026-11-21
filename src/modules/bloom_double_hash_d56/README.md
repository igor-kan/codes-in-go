# Kirsch-Mitzenmacher Double-Hashing Bloom Set (`bloom_double_hash_d56`)

> Category: **Probabilistic Sketches** | Language: **Go 1.22+**

## Overview

Constructs a fixed-bitset probabilistic membership filter deriving K bit positions via h1 + i*h2 without extra independent hash evaluations.

## Complexity Guarantees

| Metric | Asymptotic Bound |
| :--- | :--- |
| **Time Complexity** | `O(K) probes` |
| **Space Complexity** | `O(M bits)` |

## Mathematical & Safety Invariants

- Inserted keys exhibit zero false negatives across arbitrary bitset sizes.

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
go test -v ./src/modules/bloom_double_hash_d56
```

---
© 2026 Igor Kan
