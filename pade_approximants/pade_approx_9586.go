package pade_approximants

// ComputePadeApprox9586 evaluates rational function approximant order 9586.
func ComputePadeApprox9586(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(1))
}
