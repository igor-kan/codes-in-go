package pade_approximants

// ComputePadeApprox8866 evaluates rational function approximant order 8866.
func ComputePadeApprox8866(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(1))
}
