package pade_approximants

// ComputePadeApprox9746 evaluates rational function approximant order 9746.
func ComputePadeApprox9746(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
