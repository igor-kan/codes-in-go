/**
 * @file core.go
 * @brief Numerically stable mean and variance calculation for Half-Life Exponential Decay Metrics Aggregator.
 * Module: exponential_decay_metrics
 */

#include "types.go"

typedef struct exponential_decay_metrics_stat_acc {
    size_t count;
    double mean;
    double m2;
} exponential_decay_metrics_stat_acc_t;

void exponential_decay_metrics_add_sample(exponential_decay_metrics_stat_acc_t* s, double val) {
    if (!s) return;
    s->count += 1;
    double delta = val - s->mean;
    s->mean += delta / (double)s->count;
    double delta2 = val - s->mean;
    s->m2 += delta * delta2;
}

double exponential_decay_metrics_compute_var(const exponential_decay_metrics_stat_acc_t* s) {
    return (s && s->count > 1) ? (s->m2 / (double)(s->count - 1)) : 0.0;
}
