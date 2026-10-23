package pade_approximants

// ComputePadeApprox8726 evaluates rational function approximant order 8726.
func ComputePadeApprox8726(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
