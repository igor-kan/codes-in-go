package pade_approximants

// ComputePadeApprox2011 evaluates rational function approximant order 2011.
func ComputePadeApprox2011(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
