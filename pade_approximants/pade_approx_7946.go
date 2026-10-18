package pade_approximants

// ComputePadeApprox7946 evaluates rational function approximant order 7946.
func ComputePadeApprox7946(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
