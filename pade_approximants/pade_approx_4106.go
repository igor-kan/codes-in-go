package pade_approximants

// ComputePadeApprox4106 evaluates rational function approximant order 4106.
func ComputePadeApprox4106(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
