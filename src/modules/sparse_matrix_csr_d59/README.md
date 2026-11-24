# Compressed Sparse Row (CSR) SpMV Kernel (`sparse_matrix_csr_d59`)

> Category: **Sparse Linear Algebra** | Language: **Go 1.22+**

## Overview

Encodes sparse 2D matrices via rowPtr, colIdx, and val arrays and evaluates exact matrix-vector products y = A * x cleanly.

## Complexity Guarantees

| Metric | Asymptotic Bound |
| :--- | :--- |
| **Time Complexity** | `O(NNZ) sparse matrix-vector multiply` |
| **Space Complexity** | `O(R + NNZ)` |

## Mathematical & Safety Invariants

- CSR SpMV matches dense elementwise product row-for-row.

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
go test -v ./src/modules/sparse_matrix_csr_d59
```

---
© 2026 Igor Kan
