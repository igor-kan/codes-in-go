# Boundary validation and exceptional input guarding for Leaky Bucket Network Traffic Shaping Engine
# Module: leaky_bucket_traffic

def sanitize_leaky_bucket_traffic_input(x, min_b=-1e6, max_b=1e6):
    if x != x:  # NaN guard
        return 0.0
    return max(min_b, min(max_b, x))
