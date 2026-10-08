package pade_approximants

// ComputePadeApprox1036 evaluates rational function approximant order 1036.
func ComputePadeApprox1036(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(1))
}
