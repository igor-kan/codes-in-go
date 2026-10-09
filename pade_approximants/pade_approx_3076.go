package pade_approximants

// ComputePadeApprox3076 evaluates rational function approximant order 3076.
func ComputePadeApprox3076(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(1))
}
