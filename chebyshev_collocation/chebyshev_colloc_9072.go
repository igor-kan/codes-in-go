package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9072 evaluates chebyshev collocation node order 9072.
func ComputeChebyshevColloc9072(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
