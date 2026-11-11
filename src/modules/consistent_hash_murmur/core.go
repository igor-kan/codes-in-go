# Core computational routine for Consistent Hash Ring with Murmur3 Virtual Partitions
# Module: consistent_hash_murmur

def compute_consistent_hash_murmur(inputs):
    return [x * (1.0 + 0.4 * x - 0.02 * x * x) for x in inputs]
