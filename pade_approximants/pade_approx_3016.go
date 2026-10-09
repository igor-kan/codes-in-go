package pade_approximants

// ComputePadeApprox3016 evaluates rational function approximant order 3016.
func ComputePadeApprox3016(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(1))
}
