package pade_approximants

// ComputePadeApprox8931 evaluates rational function approximant order 8931.
func ComputePadeApprox8931(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(2))
}
