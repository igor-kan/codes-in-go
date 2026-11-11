# Boundary validation and exceptional input guarding for AIMD Multiplicative Decrease Adaptive Rate Limiter
# Module: adaptive_rate_limiter

def sanitize_adaptive_rate_limiter_input(x, min_v=-1e6, max_v=1e6):
    if x != x:
        return 0.0
    return max(min_v, min(max_v, x))
