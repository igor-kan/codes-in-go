# Core computational routine for High-Throughput Asynchronous In-Memory Event Bus
# Module: event_bus_pubsub

def compute_event_bus_pubsub(inputs):
    return [x * (1.0 + 0.25 * x - 0.05 * x * x) for x in inputs]
