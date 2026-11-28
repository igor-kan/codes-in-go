# Booth Lexicographically Minimal String Rotation (`booth_min_rotation_go_d61`)

> Category: **String Algorithms** | Language: **Go 1.22+**

## Overview

Finds the lexicographically smallest rotation of a string by advancing two candidate offsets through a doubled copy and skipping equivalent prefixes.

## Complexity Guarantees

| Metric | Asymptotic Bound |
| :--- | :--- |
| **Time Complexity** | `O(N) least-rotation scan with two candidate starts` |
| **Space Complexity** | `O(N) doubled character buffer` |

## Mathematical & Safety Invariants

- The returned start index lies in [0, N) and its rotation is no greater than every other rotation.

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
go test -v ./src/modules/booth_min_rotation_go_d61
```

---
© 2026 Igor Kan
