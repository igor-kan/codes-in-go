package pade_approximants

// ComputePadeApprox7021 evaluates rational function approximant order 7021.
func ComputePadeApprox7021(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
