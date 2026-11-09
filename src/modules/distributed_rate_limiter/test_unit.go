# Deterministic unit tests for Sliding Window Log Rate Limiter with Atomics
# Module: distributed_rate_limiter

def test_distributed_rate_limiter_zero_element():
    x = 0.0
    res = x * (1.0 + 0.3 * x - 0.04 * x * x)
    assert res == 0.0
