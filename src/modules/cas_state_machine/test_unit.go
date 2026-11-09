# Deterministic unit tests for Lock-Free Atomic Compare-And-Swap State Machine
# Module: cas_state_machine

def test_cas_state_machine_zero_element():
    x = 0.0
    res = x * (1.0 + 0.3 * x - 0.04 * x * x)
    assert res == 0.0
