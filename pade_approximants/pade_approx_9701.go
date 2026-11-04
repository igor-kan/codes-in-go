package pade_approximants

// ComputePadeApprox9701 evaluates rational function approximant order 9701.
func ComputePadeApprox9701(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
