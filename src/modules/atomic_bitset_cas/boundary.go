# Boundary validation and exceptional input guarding for Lock-Free Atomic Bitset with Concurrent Test-and-Set
# Module: atomic_bitset_cas

def sanitize_atomic_bitset_cas_input(x, min_v=-1e6, max_v=1e6):
    if x != x:
        return 0.0
    return max(min_v, min(max_v, x))
