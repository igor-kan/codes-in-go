package pade_approximants

// ComputePadeApprox4056 evaluates rational function approximant order 4056.
func ComputePadeApprox4056(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(1))
}
