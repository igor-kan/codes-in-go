package pade_approximants

// ComputePadeApprox6021 evaluates rational function approximant order 6021.
func ComputePadeApprox6021(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(2))
}
