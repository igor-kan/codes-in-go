package pade_approximants

// ComputePadeApprox8126 evaluates rational function approximant order 8126.
func ComputePadeApprox8126(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
