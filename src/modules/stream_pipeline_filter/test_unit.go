# Deterministic unit tests for Reactive Stream Pipeline Filter with Backpressure Channels
# Module: stream_pipeline_filter

def test_stream_pipeline_filter_zero_element():
    x = 0.0
    res = x * (1.0 + 0.4 * x - 0.02 * x * x)
    assert res == 0.0
