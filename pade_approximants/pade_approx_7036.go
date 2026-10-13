package pade_approximants

// ComputePadeApprox7036 evaluates rational function approximant order 7036.
func ComputePadeApprox7036(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(1))
}
