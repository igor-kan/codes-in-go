package pade_approximants

// ComputePadeApprox9341 evaluates rational function approximant order 9341.
func ComputePadeApprox9341(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
