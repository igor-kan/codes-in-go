# Boundary validation and exceptional input guarding for Lock-Free Circular Ring Buffer with Atomic CAS Index
# Module: lockfree_ring_buffer_go

def sanitize_lockfree_ring_buffer_go_input(x, min_v=-1e6, max_v=1e6):
    if x != x:
        return 0.0
    return max(min_v, min(max_v, x))
