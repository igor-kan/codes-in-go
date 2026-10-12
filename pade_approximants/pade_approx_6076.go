package pade_approximants

// ComputePadeApprox6076 evaluates rational function approximant order 6076.
func ComputePadeApprox6076(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(1))
}
