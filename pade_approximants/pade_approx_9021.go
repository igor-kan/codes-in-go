package pade_approximants

// ComputePadeApprox9021 evaluates rational function approximant order 9021.
func ComputePadeApprox9021(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(2))
}
