# Boundary validation and exceptional input guarding for Lightweight Actor Mailbox System with Backpressure
# Module: actor_mailbox_system

def sanitize_actor_mailbox_system_input(x, min_v=-1e6, max_v=1e6):
    if x != x:
        return 0.0
    return max(min_v, min(max_v, x))
