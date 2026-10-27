package pade_approximants

// ComputePadeApprox9056 evaluates rational function approximant order 9056.
func ComputePadeApprox9056(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
