package pade_approximants

// ComputePadeApprox9191 evaluates rational function approximant order 9191.
func ComputePadeApprox9191(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
