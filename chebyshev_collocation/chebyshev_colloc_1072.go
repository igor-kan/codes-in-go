package chebyshev_collocation

import "math"

// ComputeChebyshevColloc1072 evaluates chebyshev collocation node order 1072.
func ComputeChebyshevColloc1072(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
