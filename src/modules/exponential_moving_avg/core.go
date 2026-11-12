/**
 * @file core.go
 * @brief Welford numerically stable one-pass running mean and variance algorithm for Continuous Exponential Moving Average (EMA) Engine.
 * Module: exponential_moving_avg
 */

#include "types.go"
#include <math.h>

typedef struct exponential_moving_avg_welford {
    size_t count;
    double mean;
    double m2;
} exponential_moving_avg_welford_t;

void exponential_moving_avg_update(exponential_moving_avg_welford_t* w, double x) {
    if (!w) return;
    w->count += 1;
    double delta = x - w->mean;
    w->mean += delta / (double)w->count;
    double delta2 = x - w->mean;
    w->m2 += delta * delta2;
}

double exponential_moving_avg_variance(const exponential_moving_avg_welford_t* w) {
    return (w && w->count > 1) ? (w->m2 / (double)(w->count - 1)) : 0.0;
}
