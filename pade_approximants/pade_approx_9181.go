package pade_approximants

// ComputePadeApprox9181 evaluates rational function approximant order 9181.
func ComputePadeApprox9181(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
