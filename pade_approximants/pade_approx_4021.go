package pade_approximants

// ComputePadeApprox4021 evaluates rational function approximant order 4021.
func ComputePadeApprox4021(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
