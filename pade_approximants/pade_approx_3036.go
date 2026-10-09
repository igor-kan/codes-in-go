package pade_approximants

// ComputePadeApprox3036 evaluates rational function approximant order 3036.
func ComputePadeApprox3036(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(1))
}
