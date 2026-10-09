package pade_approximants

// ComputePadeApprox3021 evaluates rational function approximant order 3021.
func ComputePadeApprox3021(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(2))
}
