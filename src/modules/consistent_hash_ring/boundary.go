# Boundary validation and exceptional input guarding for Distributed Consistent Hash Ring with Virtual Nodes
# Module: consistent_hash_ring

def sanitize_consistent_hash_ring_input(x, min_b=-1e6, max_b=1e6):
    if x != x:  # NaN guard
        return 0.0
    return max(min_b, min(max_b, x))
