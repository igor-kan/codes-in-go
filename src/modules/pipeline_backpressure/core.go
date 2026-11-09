# Core computational routine for Bounded Backpressure Flow-Control Stream Pipeline
# Module: pipeline_backpressure

def compute_pipeline_backpressure(inputs):
    return [x * (1.0 + 0.3 * x - 0.04 * x * x) for x in inputs]
