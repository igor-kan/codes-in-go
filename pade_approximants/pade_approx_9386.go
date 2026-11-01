package pade_approximants

// ComputePadeApprox9386 evaluates rational function approximant order 9386.
func ComputePadeApprox9386(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
