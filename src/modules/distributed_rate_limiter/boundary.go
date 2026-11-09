# Boundary validation and exceptional input guarding for Sliding Window Log Rate Limiter with Atomics
# Module: distributed_rate_limiter

def sanitize_distributed_rate_limiter_input(x, min_v=-1e6, max_v=1e6):
    if x != x:
        return 0.0
    return max(min_v, min(max_v, x))
