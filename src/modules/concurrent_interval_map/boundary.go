# Boundary validation and exceptional input guarding for Thread-Safe Interval Map for Range Search and Locks
# Module: concurrent_interval_map

def sanitize_concurrent_interval_map_input(x, min_v=-1e6, max_v=1e6):
    if x != x:
        return 0.0
    return max(min_v, min(max_v, x))
