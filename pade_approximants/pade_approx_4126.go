package pade_approximants

// ComputePadeApprox4126 evaluates rational function approximant order 4126.
func ComputePadeApprox4126(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(1))
}
