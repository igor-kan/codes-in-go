package pade_approximants

// ComputePadeApprox3071 evaluates rational function approximant order 3071.
func ComputePadeApprox3071(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
