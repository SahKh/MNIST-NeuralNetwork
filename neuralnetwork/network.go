package neuralnetwork
import (
	"go-mnist-nn/utils"
)

type NeuralNetwork struct {
	InputSize int // 28x28
	HiddenSize int // hidden layer size
	OutputSize int // 10
	LearningRate float64
	
	WeightsIH [][]float64 //Weights between input and hidden layer
	WeightsHO [][]float64 //Weights between hidedn and output
	BiasHidden []float64 // hidden biases
	BiasOutput []float64 //output biases
}

func NeuralNetworkInit(_InputSize, _HiddenSize, _OutputSize int, _LearningRate float64) *NeuralNetwork {
	network := &NeuralNetwork {
		InputSize: _InputSize,
		HiddenSize: _HiddenSize,
		OutputSize: _OutputSize,
		LearningRate: _LearningRate,
	}

	network.WeightsIH = utils.RandomMatrix(_InputSize, _HiddenSize)
	network.WeightsHO = utils.RandomMatrix(_HiddenSize, _OutputSize)
	network.BiasHidden = utils.ZeroVector(_HiddenSize)
	network.BiasOutput = utils.ZeroVector(_OutputSize)

	return network
}