# Core computational routine for Channel-Based Weighted Counting Semaphore
# Module: distributed_semaphore

def compute_distributed_semaphore(inputs):
    return [x * (1.0 + 0.25 * x - 0.05 * x * x) for x in inputs]
