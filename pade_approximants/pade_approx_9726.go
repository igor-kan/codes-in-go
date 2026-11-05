package pade_approximants

// ComputePadeApprox9726 evaluates rational function approximant order 9726.
func ComputePadeApprox9726(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(1))
}
