package neuralnetwork

import "go-mnist-nn/utils"

 
/* 	Forward Propagation
	For input to hiddenlayer we calculate the dot product of the input and weights from input
to hidden layer, then add the biasses, then apply ReLU activation
*/

func (nn *NeuralNetwork) Forward(input []float64) ([]float64, []float64, []float64) {
	// Compute pre-activation values for hidden layer
	hiddenInput := utils.DotProduct(input, nn.WeightsIH)
	for i := 0; i < len(hiddenInput); i++ {
			hiddenInput[i] += nn.BiasHidden[i]
	}

	// Apply ReLU activation
	hidden := make([]float64, len(hiddenInput))
	for i := 0; i < len(hiddenInput); i++ {
			hidden[i] = utils.ReLU(hiddenInput[i])
	}

	// Compute pre-activation values for output layer
	outputInput := utils.DotProduct(hidden, nn.WeightsHO)
	for i := 0; i < len(outputInput); i++ {
			outputInput[i] += nn.BiasOutput[i]
	}

	// Apply Softmax activation
	output := utils.Softmax(outputInput)

	// Return hiddenInput (pre-activation), hidden (post-activation), and output
	return hiddenInput, hidden, output
}
