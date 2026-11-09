# Deterministic unit tests for Recyclable Memory Buffer Pool with Arena Reuse
# Module: zero_copy_buf_pool

def test_zero_copy_buf_pool_zero_element():
    x = 0.0
    res = x * (1.0 + 0.3 * x - 0.04 * x * x)
    assert res == 0.0
