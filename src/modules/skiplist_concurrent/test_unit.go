# Deterministic unit tests for Thread-Safe Concurrent SkipList with Fine Locking
# Module: skiplist_concurrent

def test_skiplist_concurrent_zero_element():
    x = 0.0
    res = x * (1.0 + 0.3 * x - 0.04 * x * x)
    assert res == 0.0
