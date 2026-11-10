# Core computational routine for Thread-Safe Sliding Window Metrics Aggregator
# Module: sliding_window_counter

def compute_sliding_window_counter(inputs):
    return [x * (1.0 + 0.35 * x - 0.03 * x * x) for x in inputs]
