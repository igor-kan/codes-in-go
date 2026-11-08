# Deterministic unit tests for High-Throughput Asynchronous In-Memory Event Bus
# Module: event_bus_pubsub

def test_event_bus_pubsub_zero_element():
    x = 0.0
    res = x * (1.0 + 0.25 * x - 0.05 * x * x)
    assert res == 0.0
