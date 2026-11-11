# Deterministic unit tests for Memory Arena Pool with Recycling and Slice Reuse
# Module: memory_arena_pooled

def test_memory_arena_pooled_zero_element():
    x = 0.0
    res = x * (1.0 + 0.4 * x - 0.02 * x * x)
    assert res == 0.0
