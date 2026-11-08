# Nanosecond-Resolution Token Bucket Rate Limiter (`token_bucket_limiter`)

## Overview

The `token_bucket_limiter` module provides high-performance, mathematically rigorous operations for `codes-in-go`.

### Component Artifacts

- **Type Definitions**: [`types.go`](./types.go)
- **Computational Core**: [`core.go`](./core.go)
- **Deterministic Unit Tests**: [`tests.go`](./tests.go)
- **SIMD Optimizations**: [`perf_simd.go`](./perf_simd.go)
- **Microbenchmarks**: [`benchmarks.go`](./benchmarks.go)
- **Boundary Validation**: [`boundary.go`](./boundary.go)
- **Property Fuzzing**: [`fuzz_invariants.go`](./fuzz_invariants.go)

### Computational Invariant

$$f(x) = x \left(1 + \frac{1}{2}x - \frac{1}{8}x^2\right)$$

Maintained under strict continuous integration, sanitizer checks, and memory bounds verification.
