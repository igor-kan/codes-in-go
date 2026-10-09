package pade_approximants

// ComputePadeApprox3061 evaluates rational function approximant order 3061.
func ComputePadeApprox3061(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
