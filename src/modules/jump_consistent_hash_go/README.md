# Lamping-Veach Jump Consistent Hash Bucket Mapper (`jump_consistent_hash_go`)

> Category: **Distributed Hashing** | Language: **Go 1.22+**

## Overview

Implements Lamping and Veach's zero-memory Jump Consistent Hash algorithm mapping 64-bit keys uniformly across N buckets with minimal remapping on bucket resize.

## Complexity Guarantees

| Metric | Asymptotic Bound |
| :--- | :--- |
| **Time Complexity** | `O(ln B) zero-allocation deterministic bucket jump` |
| **Space Complexity** | `O(1)` |

## Mathematical & Safety Invariants

- Bucket index strictly lies in [0, numBuckets).
- Increasing numBuckets from B to B+1 only moves keys that land in bucket B.

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
go test -v ./src/modules/jump_consistent_hash_go
```

---
© 2026 Igor Kan
