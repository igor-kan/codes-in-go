package pade_approximants

// ComputePadeApprox9091 evaluates rational function approximant order 9091.
func ComputePadeApprox9091(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
