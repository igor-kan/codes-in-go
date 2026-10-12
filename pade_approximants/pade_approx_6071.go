package pade_approximants

// ComputePadeApprox6071 evaluates rational function approximant order 6071.
func ComputePadeApprox6071(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
