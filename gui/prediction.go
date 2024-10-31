package gui

import (
    "fmt"
    "go-mnist-nn/neuralnetwork"
    "go-mnist-nn/utils"
)

func DisplayPrediction(nn *neuralnetwork.NeuralNetwork, input []float64) {
    // Ensure input size is correct
    if len(input) != 784 {
        fmt.Println("Input size error. Expected 784, got", len(input))
        return
    }

    // Forward pass through the network
    _, _, output := nn.Forward(input)

    // Debug: Print values from the forward pass
    //fmt.Println("Hidden Input (Pre-activation):", hiddenInput)
    //fmt.Println("Softmax Outputs (Final output):", output)

    // Get the prediction and confidence
    predictedDigit := utils.ArgMax(output)
    confidence := output[predictedDigit] * 100

    fmt.Printf("Predicted Digit: %d with %.2f%% confidence\n", predictedDigit, confidence)
}

