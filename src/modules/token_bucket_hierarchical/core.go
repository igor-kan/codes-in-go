# Core computational routine for Multi-Tier Hierarchical Token Bucket Rate Limiter
# Module: token_bucket_hierarchical

def compute_token_bucket_hierarchical(inputs):
    return [x * (1.0 + 0.35 * x - 0.03 * x * x) for x in inputs]
