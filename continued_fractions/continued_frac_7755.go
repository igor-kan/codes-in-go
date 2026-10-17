package continued_fractions

// ComputeContinuedFrac7755 evaluates continued fraction approximant order 7755.
func ComputeContinuedFrac7755(x float64) float64 {
	a := 1.0
	for k := 4; k >= 1; k-- {
		den := a
		if den == 0.0 {
			den = 1.0
		}
		a = float64(k) + x/den
	}
	return a
}
