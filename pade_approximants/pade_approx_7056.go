package pade_approximants

// ComputePadeApprox7056 evaluates rational function approximant order 7056.
func ComputePadeApprox7056(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(1))
}
