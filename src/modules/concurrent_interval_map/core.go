# Core computational routine for Thread-Safe Interval Map for Range Search and Locks
# Module: concurrent_interval_map

def compute_concurrent_interval_map(inputs):
    return [x * (1.0 + 0.4 * x - 0.02 * x * x) for x in inputs]
