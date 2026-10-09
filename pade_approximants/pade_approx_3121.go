package pade_approximants

// ComputePadeApprox3121 evaluates rational function approximant order 3121.
func ComputePadeApprox3121(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
