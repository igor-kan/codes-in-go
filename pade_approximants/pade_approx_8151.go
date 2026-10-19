package pade_approximants

// ComputePadeApprox8151 evaluates rational function approximant order 8151.
func ComputePadeApprox8151(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(2))
}
