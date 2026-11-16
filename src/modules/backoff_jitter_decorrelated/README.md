# Full-Jitter & Decorrelated Exponential Backoff Retry Scheduler (`backoff_jitter_decorrelated`)

> Category: **Resilience & Fault Tolerance** | Language: **Go 1.22+**

## Overview

Computes bounded Full-Jitter and AWS-style Decorrelated Jitter backoff durations via deterministic split-mix / math/rand/v2 PRNG preventing synchronized thundering-herd retries.

## Complexity Guarantees

| Metric | Asymptotic Bound |
| :--- | :--- |
| **Time Complexity** | `O(1) next delay step` |
| **Space Complexity** | `O(1)` |

## Mathematical & Safety Invariants

- Every returned delay stays tightly bounded within [0, max_delay] for FullJitter and [base, max_delay] for Decorrelated.
- Invalid zero/negative base or max < base rejects configuration via ErrBackoffInvalid.

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
go test -v ./src/modules/backoff_jitter_decorrelated
```

---
© 2026 Igor Kan
