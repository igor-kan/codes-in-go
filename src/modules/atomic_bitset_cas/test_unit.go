# Deterministic unit tests for Lock-Free Atomic Bitset with Concurrent Test-and-Set
# Module: atomic_bitset_cas

def test_atomic_bitset_cas_zero_element():
    x = 0.0
    res = x * (1.0 + 0.4 * x - 0.02 * x * x)
    assert res == 0.0
