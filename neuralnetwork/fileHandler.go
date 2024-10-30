// File: neuralnetwork/save_load.go
package neuralnetwork

import (
    "encoding/gob"
    "os"
    "path/filepath"
)

// Save the trained neural network model to a file inside the 'dataset/model' subfolder
func (nn *NeuralNetwork) Save(filename string) error {
    // Ensure the directory exists
    modelDir := filepath.Join("dataset", "model")
    if _, err := os.Stat(modelDir); os.IsNotExist(err) {
        err := os.MkdirAll(modelDir, os.ModePerm)
        if err != nil {
            return err
        }
    }

    // Save the model to 'dataset/model/filename'
    filePath := filepath.Join(modelDir, filename)
    file, err := os.Create(filePath)
    if err != nil {
        return err
    }
    defer file.Close()

    encoder := gob.NewEncoder(file)
    err = encoder.Encode(nn)
    if err != nil {
        return err
    }
    return nil
}

// Load the neural network model from a file inside the 'dataset/model' subfolder
func (nn *NeuralNetwork) Load(filename string) error {
    // Load the model from 'dataset/model/filename'
    filePath := filepath.Join("dataset", "model", filename)
    file, err := os.Open(filePath)
    if err != nil {
        return err
    }
    defer file.Close()

    decoder := gob.NewDecoder(file)
    err = decoder.Decode(nn)
    if err != nil {
        return err
    }
    return nil
}
