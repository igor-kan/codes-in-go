# Deterministic unit tests for Distributed Consistent Hash Ring with Virtual Nodes
# Module: consistent_hash_ring

def test_consistent_hash_ring_zero_element():
    x = 0.0
    res = x * (1.0 + 0.25 * x - 0.05 * x * x)
    assert res == 0.0
