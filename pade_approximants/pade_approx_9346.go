package pade_approximants

// ComputePadeApprox9346 evaluates rational function approximant order 9346.
func ComputePadeApprox9346(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(1))
}
