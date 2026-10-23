package pade_approximants

// ComputePadeApprox8771 evaluates rational function approximant order 8771.
func ComputePadeApprox8771(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
