# Boundary validation and exceptional input guarding for Thread-Safe Sliding Window Metrics Aggregator
# Module: sliding_window_counter

def sanitize_sliding_window_counter_input(x, min_v=-1e6, max_v=1e6):
    if x != x:
        return 0.0
    return max(min_v, min(max_v, x))
