package neuralnetwork

import (
	"fmt"
	"go-mnist-nn/utils"
	"math/rand"
	"math"
)

// ShuffleData shuffles the inputs and targets in unison.
func ShuffleData(inputs, targets [][]float64) {
	for i := range inputs {
		j := rand.Intn(i + 1)
		inputs[i], inputs[j] = inputs[j], inputs[i]
		targets[i], targets[j] = targets[j], targets[i]
	}
}

// LogWeights prints the mean and standard deviation of the weights.
func (nn *NeuralNetwork) LogWeights() {
	fmt.Println("Weights between Input and Hidden Layer:")
	fmt.Printf("Mean: %.6f, StdDev: %.6f\n", Mean(nn.WeightsIH), StdDev(nn.WeightsIH))

	fmt.Println("Weights between Hidden and Output Layer:")
	fmt.Printf("Mean: %.6f, StdDev: %.6f\n", Mean(nn.WeightsHO), StdDev(nn.WeightsHO))
}

// Calculate the mean of the weight matrix.
func Mean(matrix [][]float64) float64 {
	sum := 0.0
	count := 0
	for i := range matrix {
		for j := range matrix[i] {
			sum += matrix[i][j]
			count++
		}
	}
	return sum / float64(count)
}

// Calculate the standard deviation of the weight matrix.
func StdDev(matrix [][]float64) float64 {
	mean := Mean(matrix)
	sum := 0.0
	count := 0
	for i := range matrix {
		for j := range matrix[i] {
			diff := matrix[i][j] - mean
			sum += diff * diff
			count++
		}
	}
	return math.Sqrt(sum / float64(count))
}

// Train trains the neural network using cross-entropy loss.
func (nn *NeuralNetwork) Train(inputs, targets [][]float64, epochs int) {
	for epoch := 0; epoch < epochs; epoch++ {
		// Shuffle data before each epoch
		ShuffleData(inputs, targets)
		
		var totalLoss float64
		for i := 0; i < len(inputs); i++ {
			// Forward pass
			hiddenInput, hidden, output := nn.Forward(inputs[i])

			// Calculate cross-entropy loss
			loss := utils.CrossEntropyLoss(output, targets[i])
			totalLoss += loss

			// Backpropagation
			nn.Backpropagation(inputs[i], hiddenInput, hidden, output, targets[i])
		}
		// Print the average loss for this epoch
		fmt.Printf("Epoch %d, Loss: %.4f\n", epoch+1, totalLoss/float64(len(inputs)))

		// Log the weights after each epoch
		nn.LogWeights()
	}
}
