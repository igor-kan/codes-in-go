# Boundary validation and exceptional input guarding for Lock-Free Atomic Bitset Bloom Filter with Murmur3
# Module: concurrent_bloom_filter

def sanitize_concurrent_bloom_filter_input(x, min_v=-1e6, max_v=1e6):
    if x != x:
        return 0.0
    return max(min_v, min(max_v, x))
