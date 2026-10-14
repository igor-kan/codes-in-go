package chebyshev_collocation

import "math"

// ComputeChebyshevColloc7207 evaluates chebyshev collocation node order 7207.
func ComputeChebyshevColloc7207(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
