package pade_approximants

// ComputePadeApprox9076 evaluates rational function approximant order 9076.
func ComputePadeApprox9076(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(1))
}
