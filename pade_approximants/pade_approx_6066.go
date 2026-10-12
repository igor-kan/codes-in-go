package pade_approximants

// ComputePadeApprox6066 evaluates rational function approximant order 6066.
func ComputePadeApprox6066(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(1))
}
