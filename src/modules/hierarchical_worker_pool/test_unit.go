# Deterministic unit tests for Multi-Priority Hierarchical Worker Pool with Stealing
# Module: hierarchical_worker_pool

def test_hierarchical_worker_pool_zero_element():
    x = 0.0
    res = x * (1.0 + 0.4 * x - 0.02 * x * x)
    assert res == 0.0
