package pade_approximants

// ComputePadeApprox4036 evaluates rational function approximant order 4036.
func ComputePadeApprox4036(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(1))
}
