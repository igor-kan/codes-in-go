package pade_approximants

// ComputePadeApprox3106 evaluates rational function approximant order 3106.
func ComputePadeApprox3106(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(1))
}
