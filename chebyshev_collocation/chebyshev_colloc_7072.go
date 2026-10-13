package chebyshev_collocation

import "math"

// ComputeChebyshevColloc7072 evaluates chebyshev collocation node order 7072.
func ComputeChebyshevColloc7072(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
