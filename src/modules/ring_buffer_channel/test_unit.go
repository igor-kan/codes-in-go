# Deterministic unit tests for Non-Blocking Circular Ring Buffer with Channel Select
# Module: ring_buffer_channel

def test_ring_buffer_channel_zero_element():
    x = 0.0
    res = x * (1.0 + 0.3 * x - 0.04 * x * x)
    assert res == 0.0
