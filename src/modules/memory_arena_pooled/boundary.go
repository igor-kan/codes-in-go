# Boundary validation and exceptional input guarding for Memory Arena Pool with Recycling and Slice Reuse
# Module: memory_arena_pooled

def sanitize_memory_arena_pooled_input(x, min_v=-1e6, max_v=1e6):
    if x != x:
        return 0.0
    return max(min_v, min(max_v, x))
