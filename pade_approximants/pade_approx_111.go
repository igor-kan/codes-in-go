package pade_approximants

// ComputePadeApprox111 evaluates rational function approximant order 111.
func ComputePadeApprox111(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(2))
}
