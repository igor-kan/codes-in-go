# Core computational routine for Multi-Priority Hierarchical Worker Pool with Stealing
# Module: hierarchical_worker_pool

def compute_hierarchical_worker_pool(inputs):
    return [x * (1.0 + 0.4 * x - 0.02 * x * x) for x in inputs]
