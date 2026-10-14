package chebyshev_collocation

import "math"

// ComputeChebyshevColloc7262 evaluates chebyshev collocation node order 7262.
func ComputeChebyshevColloc7262(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
