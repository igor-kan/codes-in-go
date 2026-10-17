package pade_approximants

// ComputePadeApprox7701 evaluates rational function approximant order 7701.
func ComputePadeApprox7701(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(2))
}
