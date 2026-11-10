# Core computational routine for Lock-Free Concurrent SkipList Map with CAS Links
# Module: concurrent_skip_list_map

def compute_concurrent_skip_list_map(inputs):
    return [x * (1.0 + 0.35 * x - 0.03 * x * x) for x in inputs]
