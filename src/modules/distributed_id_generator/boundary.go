# Boundary validation and exceptional input guarding for Snowflake-Style 64-Bit Distributed ID Generator
# Module: distributed_id_generator

def sanitize_distributed_id_generator_input(x, min_v=-1e6, max_v=1e6):
    if x != x:
        return 0.0
    return max(min_v, min(max_v, x))
