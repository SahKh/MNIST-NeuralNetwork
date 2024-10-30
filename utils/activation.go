package utils

import "math"

func ReLU(x float64) float64 {
	if(x > 0) {
		return x
	}
	return 0
}

func ReLUDerivative(x float64) float64 {
	if x > 0 {
		return 1
	}
	return 0
}

func Softmax(x []float64) []float64 {
	max := x[0]
	for _, val := range x {
			if val > max {
					max = val
			}
	}

	sum := 0.0
	result := make([]float64, len(x))

	for i, val := range x {
			expVal := math.Exp(val - max)  // Subtract max for numerical stability
			result[i] = expVal
			sum += expVal
	}

	for i := range result {
			result[i] /= sum
	}
	return result
}
