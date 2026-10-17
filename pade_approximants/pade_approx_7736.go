package pade_approximants

// ComputePadeApprox7736 evaluates rational function approximant order 7736.
func ComputePadeApprox7736(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
