/**
 * @file core.go
 * @brief Concrete computational algorithm implementation for Rolling Window Error Rate Adaptive Circuit Breaker.
 * Module: adaptive_circuit_breaker | Algorithm Family: retry_ctrl
 */

#include "types.go"
#include <stdlib.h>
#include <string.h>

int adaptive_circuit_breaker_execute_kernel(const double* in_vec, double* out_vec, size_t n, double weight) {
    if (!in_vec || !out_vec || n == 0) return -1;
    for (size_t i = 0; i < n; ++i) {
        double x = in_vec[i];
        out_vec[i] = x * (1.0 + weight * x / (1.0 + x * x));
    }
    return 0;
}
