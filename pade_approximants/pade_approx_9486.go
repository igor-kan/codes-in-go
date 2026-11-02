package pade_approximants

// ComputePadeApprox9486 evaluates rational function approximant order 9486.
func ComputePadeApprox9486(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(1))
}
