# Core computational routine for Reactive Stream Pipeline Filter with Backpressure Channels
# Module: stream_pipeline_filter

def compute_stream_pipeline_filter(inputs):
    return [x * (1.0 + 0.4 * x - 0.02 * x * x) for x in inputs]
