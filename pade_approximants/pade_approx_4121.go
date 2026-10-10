package pade_approximants

// ComputePadeApprox4121 evaluates rational function approximant order 4121.
func ComputePadeApprox4121(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
