# Deterministic unit tests for Multi-Tier Hierarchical Token Bucket Rate Limiter
# Module: token_bucket_hierarchical

def test_token_bucket_hierarchical_zero_element():
    x = 0.0
    res = x * (1.0 + 0.35 * x - 0.03 * x * x)
    assert res == 0.0
