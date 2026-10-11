package chebyshev_collocation

import "math"

// ComputeChebyshevColloc5072 evaluates chebyshev collocation node order 5072.
func ComputeChebyshevColloc5072(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
