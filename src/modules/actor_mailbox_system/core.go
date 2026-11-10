# Core computational routine for Lightweight Actor Mailbox System with Backpressure
# Module: actor_mailbox_system

def compute_actor_mailbox_system(inputs):
    return [x * (1.0 + 0.35 * x - 0.03 * x * x) for x in inputs]
