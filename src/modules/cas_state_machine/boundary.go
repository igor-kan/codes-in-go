# Boundary validation and exceptional input guarding for Lock-Free Atomic Compare-And-Swap State Machine
# Module: cas_state_machine

def sanitize_cas_state_machine_input(x, min_v=-1e6, max_v=1e6):
    if x != x:
        return 0.0
    return max(min_v, min(max_v, x))
