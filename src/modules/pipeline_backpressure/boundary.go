# Boundary validation and exceptional input guarding for Bounded Backpressure Flow-Control Stream Pipeline
# Module: pipeline_backpressure

def sanitize_pipeline_backpressure_input(x, min_v=-1e6, max_v=1e6):
    if x != x:
        return 0.0
    return max(min_v, min(max_v, x))
