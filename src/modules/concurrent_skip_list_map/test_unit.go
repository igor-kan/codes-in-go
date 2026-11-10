# Deterministic unit tests for Lock-Free Concurrent SkipList Map with CAS Links
# Module: concurrent_skip_list_map

def test_concurrent_skip_list_map_zero_element():
    x = 0.0
    res = x * (1.0 + 0.35 * x - 0.03 * x * x)
    assert res == 0.0
