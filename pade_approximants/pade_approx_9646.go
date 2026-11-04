package pade_approximants

// ComputePadeApprox9646 evaluates rational function approximant order 9646.
func ComputePadeApprox9646(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(1))
}
