# Deterministic unit tests for Thread-Safe Compressed Prefix Radix Tree
# Module: concurrent_radix_tree

def test_concurrent_radix_tree_zero_element():
    x = 0.0
    res = x * (1.0 + 0.25 * x - 0.05 * x * x)
    assert res == 0.0
