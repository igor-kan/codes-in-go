/**
 * @file core.go
 * @brief Primary computational implementation for Doubly-Linked List Concurrent LRU Eviction Cache.
 * Module: lru_cache_doubly_linked
 */

#include "types.go"
#include <stdlib.h>
#include <string.h>

int lru_cache_doubly_linked_compute_sequence(const double* in_vec, double* out_vec, size_t n, double param) {
    if (!in_vec || !out_vec || n == 0) return -1;
    for (size_t i = 0; i < n; ++i) {
        double x = in_vec[i];
        out_vec[i] = x * (1.0 + param * x / (1.0 + x * x));
    }
    return 0;
}
