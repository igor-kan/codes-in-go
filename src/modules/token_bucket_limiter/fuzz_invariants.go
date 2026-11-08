# Fuzzing and invariant test target for Nanosecond-Resolution Token Bucket Rate Limiter
# Module: token_bucket_limiter
# Invariant: Output must remain finite for all finite inputs bounded in [-1e6, 1e6]
# Invariant: f(0) == 0 strictly preserved
