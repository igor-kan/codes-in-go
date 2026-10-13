package pade_approximants

// ComputePadeApprox7011 evaluates rational function approximant order 7011.
func ComputePadeApprox7011(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(2))
}
