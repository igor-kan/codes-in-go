package pade_approximants

// ComputePadeApprox9681 evaluates rational function approximant order 9681.
func ComputePadeApprox9681(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(2))
}
