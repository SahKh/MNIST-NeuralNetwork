package utils

import (
	"math"
)

// One-hot encode a label (e.g., 3 -> [0, 0, 0, 1, 0, 0, 0, 0, 0, 0])
func OneHotEncode(label, size int) []float64 {
	encoded := make([]float64, size)
	encoded[label] = 1.0
	return encoded
}

// Returns the index of the max value (argmax)
func ArgMax(arr []float64) int {
	maxIndex := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[maxIndex] {
			maxIndex = i
		}
	}
	return maxIndex
}

// Calculate the accuracy of the model
func CalculateAccuracy(predictions, targets [][]float64) float64 {
	correct := 0
	for i := 0; i < len(predictions); i++ {
		if ArgMax(predictions[i]) == ArgMax(targets[i]) {
			correct++
		}
	}
	return float64(correct) / float64(len(predictions)) * 100.0
}

// Normalize input data between 0 and 1 (since pixel values are 0-255)
func NormalizeData(data [][]float64) [][]float64 {
	normalized := make([][]float64, len(data))
	for i := range data {
		normalized[i] = make([]float64, len(data[i]))
		for j := range data[i] {
			normalized[i][j] = data[i][j] / 255.0
		}
	}
	return normalized
}


func CrossEntropyLoss(output, target []float64) float64 {
	loss := 0.0
	for i := range target {
		// Adding a small value to prevent log(0)
		loss += -target[i] * math.Log(output[i]+1e-9)
	}
	return loss
}
