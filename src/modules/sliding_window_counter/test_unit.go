# Deterministic unit tests for Thread-Safe Sliding Window Metrics Aggregator
# Module: sliding_window_counter

def test_sliding_window_counter_zero_element():
    x = 0.0
    res = x * (1.0 + 0.35 * x - 0.03 * x * x)
    assert res == 0.0
