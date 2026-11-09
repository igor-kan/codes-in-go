# Core computational routine for Lock-Free Atomic Bitset Bloom Filter with Murmur3
# Module: concurrent_bloom_filter

def compute_concurrent_bloom_filter(inputs):
    return [x * (1.0 + 0.3 * x - 0.04 * x * x) for x in inputs]
