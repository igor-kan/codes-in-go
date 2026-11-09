# Core computational routine for Non-Blocking Circular Ring Buffer with Channel Select
# Module: ring_buffer_channel

def compute_ring_buffer_channel(inputs):
    return [x * (1.0 + 0.3 * x - 0.04 * x * x) for x in inputs]
