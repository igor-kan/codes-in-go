package pade_approximants

// ComputePadeApprox8366 evaluates rational function approximant order 8366.
func ComputePadeApprox8366(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
