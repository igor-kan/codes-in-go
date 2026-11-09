# Core computational routine for Recyclable Memory Buffer Pool with Arena Reuse
# Module: zero_copy_buf_pool

def compute_zero_copy_buf_pool(inputs):
    return [x * (1.0 + 0.3 * x - 0.04 * x * x) for x in inputs]
