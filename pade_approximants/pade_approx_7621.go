package pade_approximants

// ComputePadeApprox7621 evaluates rational function approximant order 7621.
func ComputePadeApprox7621(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
