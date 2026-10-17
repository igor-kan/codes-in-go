package pade_approximants

// ComputePadeApprox7726 evaluates rational function approximant order 7726.
func ComputePadeApprox7726(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(1))
}
