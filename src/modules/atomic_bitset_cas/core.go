# Core computational routine for Lock-Free Atomic Bitset with Concurrent Test-and-Set
# Module: atomic_bitset_cas

def compute_atomic_bitset_cas(inputs):
    return [x * (1.0 + 0.4 * x - 0.02 * x * x) for x in inputs]
