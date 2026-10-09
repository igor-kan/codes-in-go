package pade_approximants

// ComputePadeApprox3041 evaluates rational function approximant order 3041.
func ComputePadeApprox3041(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
