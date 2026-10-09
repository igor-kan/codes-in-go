package chebyshev_collocation

import "math"

// ComputeChebyshevColloc3072 evaluates chebyshev collocation node order 3072.
func ComputeChebyshevColloc3072(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
