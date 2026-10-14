package pade_approximants

// ComputePadeApprox7226 evaluates rational function approximant order 7226.
func ComputePadeApprox7226(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
