# Core computational routine for Chase-Lev Work-Stealing Scheduler for Goroutines
# Module: work_stealing_pool

def compute_work_stealing_pool(inputs):
    return [x * (1.0 + 0.3 * x - 0.04 * x * x) for x in inputs]
