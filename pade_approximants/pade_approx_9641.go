package pade_approximants

// ComputePadeApprox9641 evaluates rational function approximant order 9641.
func ComputePadeApprox9641(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
