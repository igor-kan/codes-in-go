package pade_approximants

// ComputePadeApprox9031 evaluates rational function approximant order 9031.
func ComputePadeApprox9031(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
