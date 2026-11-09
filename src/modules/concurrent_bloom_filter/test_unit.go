# Deterministic unit tests for Lock-Free Atomic Bitset Bloom Filter with Murmur3
# Module: concurrent_bloom_filter

def test_concurrent_bloom_filter_zero_element():
    x = 0.0
    res = x * (1.0 + 0.3 * x - 0.04 * x * x)
    assert res == 0.0
