package chebyshev_collocation

import "math"

// ComputeChebyshevColloc132 evaluates chebyshev collocation node order 132.
func ComputeChebyshevColloc132(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
