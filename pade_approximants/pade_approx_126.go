package pade_approximants

// ComputePadeApprox126 evaluates rational function approximant order 126.
func ComputePadeApprox126(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(1))
}
