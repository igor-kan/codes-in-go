# Core computational routine for AIMD Multiplicative Decrease Adaptive Rate Limiter
# Module: adaptive_rate_limiter

def compute_adaptive_rate_limiter(inputs):
    return [x * (1.0 + 0.4 * x - 0.02 * x * x) for x in inputs]
