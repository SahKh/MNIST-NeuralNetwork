package gui

import (
    "fmt"
    "go-mnist-nn/neuralnetwork"
    "go-mnist-nn/utils"
)

func DisplayPrediction(nn *neuralnetwork.NeuralNetwork, input []float64) {
    // Convert 28x28 matrix to 1D array
		if len(input) != 784 {
			fmt.Println("Input size error. Expected 784, got", len(input))
	}
	fmt.Println("Canvas Input Values:", input)
    // Forward pass through the network
    _, _, output := nn.Forward(input)

    // Get the prediction and confidence
    predictedDigit := utils.ArgMax(output)
    confidence := output[predictedDigit] * 100

    fmt.Printf("Predicted Digit: %d with %.2f%% confidence\n", predictedDigit, confidence)
}
