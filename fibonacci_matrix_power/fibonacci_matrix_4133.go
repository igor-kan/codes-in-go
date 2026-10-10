package fibonacci_matrix_power

// ComputeFibonacciMatrix4133 evaluates fibonacci matrix power recurrence order 4133.
func ComputeFibonacciMatrix4133(x float64) float64 {
	f0, f1 := 1.0, 1.0
	for i := 0; i < 6; i++ {
		f0, f1 = f1, f0+f1*x*0.1
	}
	return f1
}
