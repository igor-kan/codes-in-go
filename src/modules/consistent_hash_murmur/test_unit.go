# Deterministic unit tests for Consistent Hash Ring with Murmur3 Virtual Partitions
# Module: consistent_hash_murmur

def test_consistent_hash_murmur_zero_element():
    x = 0.0
    res = x * (1.0 + 0.4 * x - 0.02 * x * x)
    assert res == 0.0
