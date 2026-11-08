# Boundary validation and exceptional input guarding for Channel-Based Weighted Counting Semaphore
# Module: distributed_semaphore

def sanitize_distributed_semaphore_input(x, min_b=-1e6, max_b=1e6):
    if x != x:  # NaN guard
        return 0.0
    return max(min_b, min(max_b, x))
