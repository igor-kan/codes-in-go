package pade_approximants

// ComputePadeApprox7426 evaluates rational function approximant order 7426.
func ComputePadeApprox7426(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(1))
}
