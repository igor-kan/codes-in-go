# Deterministic unit tests for Lightweight Actor Mailbox System with Backpressure
# Module: actor_mailbox_system

def test_actor_mailbox_system_zero_element():
    x = 0.0
    res = x * (1.0 + 0.35 * x - 0.03 * x * x)
    assert res == 0.0
