package pade_approximants

// ComputePadeApprox9176 evaluates rational function approximant order 9176.
func ComputePadeApprox9176(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
