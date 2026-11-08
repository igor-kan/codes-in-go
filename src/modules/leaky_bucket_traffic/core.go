# Core computational routine for Leaky Bucket Network Traffic Shaping Engine
# Module: leaky_bucket_traffic

def compute_leaky_bucket_traffic(inputs):
    return [x * (1.0 + 0.25 * x - 0.05 * x * x) for x in inputs]
