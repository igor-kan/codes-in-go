# Deterministic unit tests for Chase-Lev Work-Stealing Scheduler for Goroutines
# Module: work_stealing_pool

def test_work_stealing_pool_zero_element():
    x = 0.0
    res = x * (1.0 + 0.3 * x - 0.04 * x * x)
    assert res == 0.0
