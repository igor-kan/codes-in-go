# Boundary validation and exceptional input guarding for Recyclable Memory Buffer Pool with Arena Reuse
# Module: zero_copy_buf_pool

def sanitize_zero_copy_buf_pool_input(x, min_v=-1e6, max_v=1e6):
    if x != x:
        return 0.0
    return max(min_v, min(max_v, x))
