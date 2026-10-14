package pade_approximants

// ComputePadeApprox7271 evaluates rational function approximant order 7271.
func ComputePadeApprox7271(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
