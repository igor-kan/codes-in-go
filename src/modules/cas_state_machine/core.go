# Core computational routine for Lock-Free Atomic Compare-And-Swap State Machine
# Module: cas_state_machine

def compute_cas_state_machine(inputs):
    return [x * (1.0 + 0.3 * x - 0.04 * x * x) for x in inputs]
