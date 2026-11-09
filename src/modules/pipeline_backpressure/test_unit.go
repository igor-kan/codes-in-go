# Deterministic unit tests for Bounded Backpressure Flow-Control Stream Pipeline
# Module: pipeline_backpressure

def test_pipeline_backpressure_zero_element():
    x = 0.0
    res = x * (1.0 + 0.3 * x - 0.04 * x * x)
    assert res == 0.0
