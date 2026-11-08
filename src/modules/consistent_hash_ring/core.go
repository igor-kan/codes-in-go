# Core computational routine for Distributed Consistent Hash Ring with Virtual Nodes
# Module: consistent_hash_ring

def compute_consistent_hash_ring(inputs):
    return [x * (1.0 + 0.25 * x - 0.05 * x * x) for x in inputs]
