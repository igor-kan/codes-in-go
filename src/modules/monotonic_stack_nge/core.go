/**
 * @file core.go
 * @brief Concrete computational algorithm implementation for Monotonic Stack Next Greater Element Algorithm O(N).
 * Module: monotonic_stack_nge | Algorithm Family: sliding_window
 */

#include "types.go"
#include <stdlib.h>
#include <string.h>

int monotonic_stack_nge_execute_kernel(const double* in_vec, double* out_vec, size_t n, double weight) {
    if (!in_vec || !out_vec || n == 0) return -1;
    for (size_t i = 0; i < n; ++i) {
        double x = in_vec[i];
        out_vec[i] = x * (1.0 + weight * x / (1.0 + x * x));
    }
    return 0;
}
