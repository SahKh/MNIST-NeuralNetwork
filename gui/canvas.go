// File: gui/canvas.go
package gui

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/canvas"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
    "image/color"
)

var drawnData [28][28]float64 // Store the drawn image data

// CreateCanvasWindow sets up the window where users can draw digits
func blurImage(input [][]float64) [][]float64 {
	blurred := make([][]float64, 28)
	for i := 0; i < 28; i++ {
			blurred[i] = make([]float64, 28)
			for j := 0; j < 28; j++ {
					sum := input[i][j]
					count := 1
					if i > 0 {
							sum += input[i-1][j]
							count++
					}
					if i < 27 {
							sum += input[i+1][j]
							count++
					}
					if j > 0 {
							sum += input[i][j-1]
							count++
					}
					if j < 27 {
							sum += input[i][j+1]
							count++
					}
					blurred[i][j] = sum / float64(count)
			}
	}
	return blurred
}

func thresholdInput(input [][]float64, threshold float64) [][]float64 {
	for i := 0; i < 28; i++ {
			for j := 0; j < 28; j++ {
					if input[i][j] < threshold {
							input[i][j] = 0
					} else {
							input[i][j] = 1
					}
			}
	}
	return input
}

func CreateCanvasWindow(onSubmit func([][]float64)) {
    myApp := app.New()
    if myApp == nil {
        myApp = app.New()
    }
    myWindow := myApp.NewWindow("Draw")

    var drawCanvas *DrawingCanvas
    drawCanvas = NewDrawingCanvas(func() {
        input := convertDrawnDataToInput()
        onSubmit(input)
        clearDrawnData()
        drawCanvas.Refresh()
    })

    container := container.NewVBox(drawCanvas)
    myWindow.SetContent(container)
    myWindow.Resize(fyne.NewSize(300, 300))
    myWindow.ShowAndRun()
}

// DrawingCanvas is the canvas where the user can draw digits
type DrawingCanvas struct {
    widget.BaseWidget
    img      *canvas.Raster
    onSubmit func()
}

func NewDrawingCanvas(onSubmit func()) *DrawingCanvas {
    dc := &DrawingCanvas{onSubmit: onSubmit}
    dc.img = canvas.NewRasterWithPixels(dc.draw)
    dc.ExtendBaseWidget(dc)
    return dc
}

func (dc *DrawingCanvas) CreateRenderer() fyne.WidgetRenderer {
    return &drawingCanvasRenderer{dc: dc}
}

type drawingCanvasRenderer struct {
    dc *DrawingCanvas
}

func (r *drawingCanvasRenderer) Layout(size fyne.Size) {
    r.dc.img.Resize(size)
}

func (r *drawingCanvasRenderer) MinSize() fyne.Size {
    return fyne.NewSize(280, 280)
}

func (r *drawingCanvasRenderer) Refresh() {
    r.dc.img.Refresh()
}

func (r *drawingCanvasRenderer) BackgroundColor() color.Color {
    return color.White
}

func (r *drawingCanvasRenderer) Objects() []fyne.CanvasObject {
    return []fyne.CanvasObject{r.dc.img}
}

func (r *drawingCanvasRenderer) Destroy() {}



// Tapped handles the mouse click event on the canvas
func (dc *DrawingCanvas) Tapped(ev *fyne.PointEvent) {
    dc.updateCanvas(int(ev.Position.X), int(ev.Position.Y))
}

// Dragged handles the mouse drag event on the canvas
func (dc *DrawingCanvas) Dragged(ev *fyne.DragEvent) {
    dc.updateCanvas(int(ev.Position.X), int(ev.Position.Y))
}

// DragEnd handles the event when the user releases the mouse button
func (dc *DrawingCanvas) DragEnd() {
    dc.onSubmit()
    clearDrawnData()
    dc.Refresh()
}

func (dc *DrawingCanvas) draw(x, y, w, h int) color.Color {
	gridX := x * 28 / w
	gridY := y * 28 / h

	if gridX >= 28 || gridY >= 28 || gridX < 0 || gridY < 0 {
			return color.White
	}

	if drawnData[gridX][gridY] > 0 {
			return color.Black
	}
	return color.White
}

// updateCanvas updates the drawnData array with the user's drawing
func (dc *DrawingCanvas) updateCanvas(x, y int) {
	width := int(dc.Size().Width)
	height := int(dc.Size().Height)

	gridX := x * 28 / width
	gridY := y * 28 / height

	if gridX >= 28 || gridY >= 28 || gridX < 0 || gridY < 0 {
			return
	}

	for dx := -1; dx <= 1; dx++ {
			for dy := -1; dy <= 1; dy++ {
					newX := gridX + dx
					newY := gridY + dy
					if newX >= 0 && newX < 28 && newY >= 0 && newY < 28 {
							drawnData[newX][newY] = 1
					}
			}
	}
	dc.Refresh()
}

// Convert drawnData to a format suitable for the neural network input

func mirrorMatrix(matrix [][]float64) [][]float64 {
	for i := range matrix {
		// Reverse each row
		for j, k := 0, len(matrix[i])-1; j < k; j, k = j+1, k-1 {
			matrix[i][j], matrix[i][k] = matrix[i][k], matrix[i][j]
		}
	}
	return matrix
}

func convertDrawnDataToInput() [][]float64 {
    adjustedData := make([][]float64, 28)
    for i := 0; i < 28; i++ {
        adjustedData[i] = make([]float64, 28)
    }

    for i := 0; i < 28; i++ {
        for j := 0; j < 28; j++ {
            // Rotate 90 degrees clockwise and flip horizontally
            adjustedData[i][j] = drawnData[27-j][i]
        }
    }
    return mirrorMatrix(adjustedData)
}


// clearDrawnData clears the canvas by resetting drawnData
func clearDrawnData() {
    for i := range drawnData {
        for j := range drawnData[i] {
            drawnData[i][j] = 0
        }
    }
}
