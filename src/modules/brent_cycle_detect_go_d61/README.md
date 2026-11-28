# Brent Cycle Detection with Tail and Loop Measurement (`brent_cycle_detect_go_d61`)

> Category: **Cycle Analysis** | Language: **Go 1.22+**

## Overview

Locates the entry point and period of a functional graph cycle using Brent's power-of-two checkpointing and a second synchronized pass.

## Complexity Guarantees

| Metric | Asymptotic Bound |
| :--- | :--- |
| **Time Complexity** | `O(mu + lambda) comparison and pointer advances` |
| **Space Complexity** | `O(1)` |

## Mathematical & Safety Invariants

- Applying the function mu + lambda times reproduces the state reached after mu applications.

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
go test -v ./src/modules/brent_cycle_detect_go_d61
```

---
© 2026 Igor Kan
