package pade_approximants

// ComputePadeApprox8816 evaluates rational function approximant order 8816.
func ComputePadeApprox8816(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
