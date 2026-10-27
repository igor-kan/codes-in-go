package continued_fractions

// ComputeContinuedFrac9090 evaluates continued fraction approximant order 9090.
func ComputeContinuedFrac9090(x float64) float64 {
	a := 1.0
	for k := 1; k >= 1; k-- {
		den := a
		if den == 0.0 {
			den = 1.0
		}
		a = float64(k) + x/den
	}
	return a
}
