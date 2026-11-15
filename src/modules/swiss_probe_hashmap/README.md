# SwissTable-Style SWAR Control-Byte Probing Hash Map (`swiss_probe_hashmap`)

> Category: **High-Performance Data Structures** | Language: **Go 1.22+**

## Overview

Separates 8-byte SWAR control metadata words (0x80 empty, 0xFE tombstone, 7-bit H2 fingerprint tags) from key-value slots to filter misses with bitwise masks before touching payloads.

## Complexity Guarantees

| Metric | Asymptotic Bound |
| :--- | :--- |
| **Time Complexity** | `O(1) expected Get/Put/Delete via 8-slot control word scans` |
| **Space Complexity** | `O(N) 1-byte metadata ctrl tags + flat key/val slots` |

## Mathematical & Safety Invariants

- Live item count equals occupied control slots without 0x80/0xFE markers.
- Rehashing preserves every inserted live key-value association identically.

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
go test -v ./src/modules/swiss_probe_hashmap
```

---
© 2026 Igor Kan
