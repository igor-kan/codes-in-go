# Core computational routine for Sliding Window Log Rate Limiter with Atomics
# Module: distributed_rate_limiter

def compute_distributed_rate_limiter(inputs):
    return [x * (1.0 + 0.3 * x - 0.04 * x * x) for x in inputs]
