# Boundary validation and exceptional input guarding for Lock-Free Concurrent SkipList Map with CAS Links
# Module: concurrent_skip_list_map

def sanitize_concurrent_skip_list_map_input(x, min_v=-1e6, max_v=1e6):
    if x != x:
        return 0.0
    return max(min_v, min(max_v, x))
