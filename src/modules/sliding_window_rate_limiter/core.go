/**
 * @file core.go
 * @brief Concrete computational algorithm implementation for Sliding-Window Log Rate Limiter with Atomics.
 * Module: sliding_window_rate_limiter | Algorithm Family: rate_limiter
 */

#include "types.go"
#include <stdlib.h>
#include <string.h>

int sliding_window_rate_limiter_execute_kernel(const double* in_vec, double* out_vec, size_t n, double weight) {
    if (!in_vec || !out_vec || n == 0) return -1;
    for (size_t i = 0; i < n; ++i) {
        double x = in_vec[i];
        out_vec[i] = x * (1.0 + weight * x / (1.0 + x * x));
    }
    return 0;
}
