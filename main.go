package main

import (
    "fmt"
    "go-mnist-nn/gui"
    "go-mnist-nn/neuralnetwork"
    "go-mnist-nn/utils"
    "os"
)
func testModel(nn *neuralnetwork.NeuralNetwork) {
    // Load test images and labels
    testImages, err := utils.ReadImages("dataset/t10k-images.idx3-ubyte")
    if err != nil {
        fmt.Println("Error reading test images:", err)
        return
    }

    testLabels, err := utils.ReadLabels("dataset/t10k-labels.idx1-ubyte")
    if err != nil {
        fmt.Println("Error reading test labels:", err)
        return
    }

    // Normalize test images
    testImages = utils.NormalizeData(testImages)

    // Prepare one-hot encoded targets for comparison
    correctPredictions := 0
    totalPredictions := len(testImages)

    for i := 0; i < totalPredictions; i++ {
        // Forward pass: Get hiddenInput, hidden, and output
        _, _, output := nn.Forward(testImages[i])

        // Get the predicted class (digit) and the actual class
        predictedDigit := utils.ArgMax(output)
        actualDigit := testLabels[i]

        // Count correct predictions
        if predictedDigit == actualDigit {
            correctPredictions++
						//fmt.Printf("Correct: %d\n", predictedDigit)
        } else {
						//fmt.Printf("Incorrect: %d\n", predictedDigit)
				}

        // Optional: Print the first few test results for debugging
        if i < 10 {
            fmt.Printf("Sample %d: Predicted = %d, Actual = %d\n", i+1, predictedDigit, actualDigit)
        }
				//fmt.Printf("Test: %d\n", i)
    }

    // Calculate accuracy
    accuracy := float64(correctPredictions) / float64(totalPredictions) * 100
    fmt.Printf("Accuracy on test set: %.2f%%\n", accuracy)
}


func main() {
    nn := neuralnetwork.NeuralNetworkInit(784, 256, 10, 0.01)

    // Check if a trained model already exists in 'dataset/model/model.gob'
    if _, err := os.Stat("dataset/model/model.gob"); err == nil {
        // Load the trained model
        err := nn.Load("model.gob")
        if err != nil {
            fmt.Println("Error loading model:", err)
            return
        }
        fmt.Println("Model loaded successfully.")
    } else {
        // Load training data (replace test data with training data)
        images, err := utils.ReadImages("dataset/train-images.idx3-ubyte")
        if err != nil {
            fmt.Println("Error reading images:", err)
            return
        }

        labels, err := utils.ReadLabels("dataset/train-labels.idx1-ubyte")
        if err != nil {
            fmt.Println("Error reading labels:", err)
            return
        }
				images = utils.NormalizeData(images)
        // Prepare targets (one-hot encoding)
        targets := make([][]float64, len(labels))
        for i, label := range labels {
            targets[i] = utils.OneHotEncode(label, 10)
        }

        // Train the neural network for 
        nn.Train(images, targets, 100)

        // Save the trained model
        err = nn.Save("model.gob")
        if err != nil {
            fmt.Println("Error saving model:", err)
            return
        }
        fmt.Println("Model saved successfully.")
    }
		testModel(nn)
    // Create the GUI for drawing digits
    gui.CreateCanvasWindow(func(input [][]float64) {
			normalizedInput := utils.NormalizeData(input)
      flatInput := make([]float64, 0)
        for _, row := range normalizedInput {
            flatInput = append(flatInput, row...)
        }
       
        gui.DisplayPrediction(nn, flatInput)
    })
}
