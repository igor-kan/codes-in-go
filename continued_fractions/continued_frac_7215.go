package continued_fractions

// ComputeContinuedFrac7215 evaluates continued fraction approximant order 7215.
func ComputeContinuedFrac7215(x float64) float64 {
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
