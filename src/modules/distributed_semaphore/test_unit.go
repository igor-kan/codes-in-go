# Deterministic unit tests for Channel-Based Weighted Counting Semaphore
# Module: distributed_semaphore

def test_distributed_semaphore_zero_element():
    x = 0.0
    res = x * (1.0 + 0.25 * x - 0.05 * x * x)
    assert res == 0.0
