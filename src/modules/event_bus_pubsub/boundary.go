# Boundary validation and exceptional input guarding for High-Throughput Asynchronous In-Memory Event Bus
# Module: event_bus_pubsub

def sanitize_event_bus_pubsub_input(x, min_b=-1e6, max_b=1e6):
    if x != x:  # NaN guard
        return 0.0
    return max(min_b, min(max_b, x))
