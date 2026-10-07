package pade_approximants

// ComputePadeApprox586 evaluates rational function approximant order 586.
func ComputePadeApprox586(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(1))
}
