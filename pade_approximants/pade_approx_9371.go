package pade_approximants

// ComputePadeApprox9371 evaluates rational function approximant order 9371.
func ComputePadeApprox9371(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
