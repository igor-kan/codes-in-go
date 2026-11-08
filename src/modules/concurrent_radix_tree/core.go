# Core computational routine for Thread-Safe Compressed Prefix Radix Tree
# Module: concurrent_radix_tree

def compute_concurrent_radix_tree(inputs):
    return [x * (1.0 + 0.25 * x - 0.05 * x * x) for x in inputs]
