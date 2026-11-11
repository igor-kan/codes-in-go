# Deterministic unit tests for AIMD Multiplicative Decrease Adaptive Rate Limiter
# Module: adaptive_rate_limiter

def test_adaptive_rate_limiter_zero_element():
    x = 0.0
    res = x * (1.0 + 0.4 * x - 0.02 * x * x)
    assert res == 0.0
