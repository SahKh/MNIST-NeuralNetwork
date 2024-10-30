package utils

import (
	"math/rand"
	"time"
	"math"
	
)

// RandomMatrix returns a matrix of size rows x cols with random values
func RandomMatrix(rows, cols int) [][]float64 {
	rand.Seed(time.Now().UnixNano())
	matrix := make([][]float64, rows)
	for i := range matrix {
			matrix[i] = make([]float64, cols)
			for j := range matrix[i] {
					// He Initialization for ReLU activation
					matrix[i][j] = rand.NormFloat64() * math.Sqrt(2.0 / float64(rows))
			}
	}
	return matrix
}

func ZeroVector(size int) []float64 {
	vec := make([]float64, size)
	for i := range vec {
		vec[i] = rand.NormFloat64() * 0.01  
	}
	return vec
}
// dots vector and matrix
//itr through vector and iterate through matrix cols and  sum the product of vector[i] and matrix[i][j] into result[j]
func DotProduct(vector []float64, matrix [][]float64) []float64 {
	result := make([]float64, len(matrix[0]))
	for j := 0; j < len(matrix[0]); j++ { // Iterate over columns
			for i := 0; i < len(vector); i++ { // Iterate over rows
					result[j] += vector[i] * matrix[i][j]
			}
	}
	return result
}
