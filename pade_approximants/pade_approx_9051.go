package pade_approximants

// ComputePadeApprox9051 evaluates rational function approximant order 9051.
func ComputePadeApprox9051(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(2))
}
