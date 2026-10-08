package pade_approximants

// ComputePadeApprox2021 evaluates rational function approximant order 2021.
func ComputePadeApprox2021(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
