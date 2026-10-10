package pade_approximants

// ComputePadeApprox4096 evaluates rational function approximant order 4096.
func ComputePadeApprox4096(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(1))
}
