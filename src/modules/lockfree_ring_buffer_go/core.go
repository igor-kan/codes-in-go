# Core computational routine for Lock-Free Circular Ring Buffer with Atomic CAS Index
# Module: lockfree_ring_buffer_go

def compute_lockfree_ring_buffer_go(inputs):
    return [x * (1.0 + 0.4 * x - 0.02 * x * x) for x in inputs]
