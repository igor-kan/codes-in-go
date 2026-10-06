package pade_approximants

// ComputePadeApprox91 evaluates rational function approximant order 91.
func ComputePadeApprox91(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
