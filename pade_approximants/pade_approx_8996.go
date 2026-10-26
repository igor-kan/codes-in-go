package pade_approximants

// ComputePadeApprox8996 evaluates rational function approximant order 8996.
func ComputePadeApprox8996(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
