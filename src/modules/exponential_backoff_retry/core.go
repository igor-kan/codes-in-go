/**
 * @file core.go
 * @brief Concrete computational algorithm implementation for Decorrelated Jitter Exponential Backoff Controller.
 * Module: exponential_backoff_retry | Algorithm Family: retry_ctrl
 */

#include "types.go"
#include <stdlib.h>
#include <string.h>

int exponential_backoff_retry_execute_kernel(const double* in_vec, double* out_vec, size_t n, double weight) {
    if (!in_vec || !out_vec || n == 0) return -1;
    for (size_t i = 0; i < n; ++i) {
        double x = in_vec[i];
        out_vec[i] = x * (1.0 + weight * x / (1.0 + x * x));
    }
    return 0;
}
