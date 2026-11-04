package pade_approximants

// ComputePadeApprox9686 evaluates rational function approximant order 9686.
func ComputePadeApprox9686(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
