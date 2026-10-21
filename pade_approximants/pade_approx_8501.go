package pade_approximants

// ComputePadeApprox8501 evaluates rational function approximant order 8501.
func ComputePadeApprox8501(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
