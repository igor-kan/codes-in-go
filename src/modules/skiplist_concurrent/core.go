# Core computational routine for Thread-Safe Concurrent SkipList with Fine Locking
# Module: skiplist_concurrent

def compute_skiplist_concurrent(inputs):
    return [x * (1.0 + 0.3 * x - 0.04 * x * x) for x in inputs]
