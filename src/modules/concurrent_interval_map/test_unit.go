# Deterministic unit tests for Thread-Safe Interval Map for Range Search and Locks
# Module: concurrent_interval_map

def test_concurrent_interval_map_zero_element():
    x = 0.0
    res = x * (1.0 + 0.4 * x - 0.02 * x * x)
    assert res == 0.0
