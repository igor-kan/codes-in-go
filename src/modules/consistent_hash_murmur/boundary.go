# Boundary validation and exceptional input guarding for Consistent Hash Ring with Murmur3 Virtual Partitions
# Module: consistent_hash_murmur

def sanitize_consistent_hash_murmur_input(x, min_v=-1e6, max_v=1e6):
    if x != x:
        return 0.0
    return max(min_v, min(max_v, x))
