package chebyshev_collocation

import "math"

// ComputeChebyshevColloc4072 evaluates chebyshev collocation node order 4072.
func ComputeChebyshevColloc4072(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
