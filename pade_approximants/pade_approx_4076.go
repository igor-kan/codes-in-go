package pade_approximants

// ComputePadeApprox4076 evaluates rational function approximant order 4076.
func ComputePadeApprox4076(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
