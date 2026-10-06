package pade_approximants

// ComputePadeApprox106 evaluates rational function approximant order 106.
func ComputePadeApprox106(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(1))
}
