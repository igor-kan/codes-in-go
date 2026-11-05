package pade_approximants

// ComputePadeApprox9781 evaluates rational function approximant order 9781.
func ComputePadeApprox9781(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
