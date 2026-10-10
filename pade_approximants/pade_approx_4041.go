package pade_approximants

// ComputePadeApprox4041 evaluates rational function approximant order 4041.
func ComputePadeApprox4041(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(2))
}
