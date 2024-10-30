package utils

import (
    "encoding/binary"
    "fmt"
    "io"
    "os"
)

// ReadImages reads images from an IDX file.
func ReadImages(filename string) ([][]float64, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var magicNumber int32
    err = binary.Read(file, binary.BigEndian, &magicNumber)
    if err != nil {
        return nil, err
    }
    if magicNumber != 2051 {
        return nil, fmt.Errorf("invalid magic number %d in %s", magicNumber, filename)
    }

    var numImages, numRows, numCols int32
    err = binary.Read(file, binary.BigEndian, &numImages)
    err = binary.Read(file, binary.BigEndian, &numRows)
    err = binary.Read(file, binary.BigEndian, &numCols)

    images := make([][]float64, numImages)
    imageSize := numRows * numCols

    for i := int32(0); i < numImages; i++ {
        buf := make([]byte, imageSize)
        _, err := io.ReadFull(file, buf)
        if err != nil {
            return nil, err
        }
        img := make([]float64, imageSize)
        for j := int32(0); j < imageSize; j++ {
            img[j] = float64(buf[j]) / 255.0 // Normalize pixel values
        }
        images[i] = img
    }

    return images, nil
}

// ReadLabels reads labels from an IDX file.
func ReadLabels(filename string) ([]int, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var magicNumber int32
    err = binary.Read(file, binary.BigEndian, &magicNumber)
    if err != nil {
        return nil, err
    }
    if magicNumber != 2049 {
        return nil, fmt.Errorf("invalid magic number %d in %s", magicNumber, filename)
    }

    var numLabels int32
    err = binary.Read(file, binary.BigEndian, &numLabels)

    labels := make([]int, numLabels)
    buf := make([]byte, numLabels)
    _, err = io.ReadFull(file, buf)
    if err != nil {
        return nil, err
    }
    for i := int32(0); i < numLabels; i++ {
        labels[i] = int(buf[i])
    }

    return labels, nil
}
