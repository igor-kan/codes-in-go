package pade_approximants

// ComputePadeApprox9576 evaluates rational function approximant order 9576.
func ComputePadeApprox9576(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(1))
}
