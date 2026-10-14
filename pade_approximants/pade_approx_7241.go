package pade_approximants

// ComputePadeApprox7241 evaluates rational function approximant order 7241.
func ComputePadeApprox7241(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
