package pade_approximants

// ComputePadeApprox7706 evaluates rational function approximant order 7706.
func ComputePadeApprox7706(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
