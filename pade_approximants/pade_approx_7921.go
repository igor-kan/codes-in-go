package pade_approximants

// ComputePadeApprox7921 evaluates rational function approximant order 7921.
func ComputePadeApprox7921(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
