package pade_approximants

// ComputePadeApprox576 evaluates rational function approximant order 576.
func ComputePadeApprox576(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(1))
}
