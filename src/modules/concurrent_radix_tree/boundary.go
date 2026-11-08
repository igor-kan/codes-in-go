# Boundary validation and exceptional input guarding for Thread-Safe Compressed Prefix Radix Tree
# Module: concurrent_radix_tree

def sanitize_concurrent_radix_tree_input(x, min_b=-1e6, max_b=1e6):
    if x != x:  # NaN guard
        return 0.0
    return max(min_b, min(max_b, x))
