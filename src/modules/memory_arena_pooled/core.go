# Core computational routine for Memory Arena Pool with Recycling and Slice Reuse
# Module: memory_arena_pooled

def compute_memory_arena_pooled(inputs):
    return [x * (1.0 + 0.4 * x - 0.02 * x * x) for x in inputs]
