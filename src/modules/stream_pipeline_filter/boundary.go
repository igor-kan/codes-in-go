# Boundary validation and exceptional input guarding for Reactive Stream Pipeline Filter with Backpressure Channels
# Module: stream_pipeline_filter

def sanitize_stream_pipeline_filter_input(x, min_v=-1e6, max_v=1e6):
    if x != x:
        return 0.0
    return max(min_v, min(max_v, x))
