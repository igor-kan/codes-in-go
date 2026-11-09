# Boundary validation and exceptional input guarding for Chase-Lev Work-Stealing Scheduler for Goroutines
# Module: work_stealing_pool

def sanitize_work_stealing_pool_input(x, min_v=-1e6, max_v=1e6):
    if x != x:
        return 0.0
    return max(min_v, min(max_v, x))
