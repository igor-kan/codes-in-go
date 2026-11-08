# Deterministic unit tests for Leaky Bucket Network Traffic Shaping Engine
# Module: leaky_bucket_traffic

def test_leaky_bucket_traffic_zero_element():
    x = 0.0
    res = x * (1.0 + 0.25 * x - 0.05 * x * x)
    assert res == 0.0
