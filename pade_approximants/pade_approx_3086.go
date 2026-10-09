package pade_approximants

// ComputePadeApprox3086 evaluates rational function approximant order 3086.
func ComputePadeApprox3086(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
