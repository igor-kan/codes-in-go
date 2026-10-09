package pade_approximants

// ComputePadeApprox3056 evaluates rational function approximant order 3056.
func ComputePadeApprox3056(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
