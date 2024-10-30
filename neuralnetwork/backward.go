package neuralnetwork

import "go-mnist-nn/utils"

// Backpropagation handles the backward pass of the network
func (nn *NeuralNetwork) Backpropagation(input, hiddenInput, hidden, output, target []float64) {
	// Output layer learning rate (higher learning rate for the output layer)
	outputLearningRate := nn.LearningRate

	// Calculate output layer error (error = output - target)
	outputError := make([]float64, nn.OutputSize)
	for i := 0; i < nn.OutputSize; i++ {
		outputError[i] = output[i] - target[i]
	}

	// Backpropagate the error to hidden layer
	hiddenError := make([]float64, nn.HiddenSize)
	for i := 0; i < nn.HiddenSize; i++ {
		for j := 0; j < nn.OutputSize; j++ {
			hiddenError[i] += outputError[j] * nn.WeightsHO[i][j]
		}
		// Use the pre-activation `hiddenInput` to calculate ReLU derivative
		hiddenError[i] *= utils.ReLUDerivative(hiddenInput[i])
	}

	// Update weights from hidden to output layer with a larger learning rate
	for i := 0; i < nn.HiddenSize; i++ {
		for j := 0; j < nn.OutputSize; j++ {
			nn.WeightsHO[i][j] -= outputLearningRate * outputError[j] * hidden[i]
		}
	}

	// Update biases for output layer with the larger learning rate
	for j := 0; j < nn.OutputSize; j++ {
		nn.BiasOutput[j] -= outputLearningRate * outputError[j]
	}

	// Update weights from input to hidden layer
	for i := 0; i < nn.InputSize; i++ {
		for j := 0; j < nn.HiddenSize; j++ {
			nn.WeightsIH[i][j] -= nn.LearningRate * hiddenError[j] * input[i]
		}
	}

	// Update biases for hidden layer
	for j := 0; j < nn.HiddenSize; j++ {
		nn.BiasHidden[j] -= nn.LearningRate * hiddenError[j]
	}
}

