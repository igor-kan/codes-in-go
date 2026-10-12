package pade_approximants

// ComputePadeApprox6056 evaluates rational function approximant order 6056.
func ComputePadeApprox6056(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
