package pade_approximants

// ComputePadeApprox9071 evaluates rational function approximant order 9071.
func ComputePadeApprox9071(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
