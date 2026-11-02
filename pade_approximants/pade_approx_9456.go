package pade_approximants

// ComputePadeApprox9456 evaluates rational function approximant order 9456.
func ComputePadeApprox9456(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(1))
}
