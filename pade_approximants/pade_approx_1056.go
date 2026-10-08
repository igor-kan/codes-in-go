package pade_approximants

// ComputePadeApprox1056 evaluates rational function approximant order 1056.
func ComputePadeApprox1056(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(1))
}
