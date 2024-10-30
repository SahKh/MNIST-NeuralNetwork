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

// draw draws the pixel on the canvas based on the mouse position
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

// updateCanvas updates the drawnData array with the user's drawing
func (dc *DrawingCanvas) updateCanvas(x, y int) {
    width := int(dc.Size().Width)
    height := int(dc.Size().Height)

    gridX := x * 28 / width
    gridY := y * 28 / height

    if gridX >= 28 || gridY >= 28 || gridX < 0 || gridY < 0 {
        return
    }

    drawnData[gridX][gridY] = 1 // Mark the pixel as drawn
    dc.Refresh()
}

// Convert drawnData to a format suitable for the neural network input
func convertDrawnDataToInput() [][]float64 {
    input := make([][]float64, 28)
    for i := 0; i < 28; i++ {
        input[i] = make([]float64, 28)
        for j := 0; j < 28; j++ {
            input[i][j] = drawnData[i][j] * 255.0
        }
    }
    return input
}

// clearDrawnData clears the canvas by resetting drawnData
func clearDrawnData() {
    for i := range drawnData {
        for j := range drawnData[i] {
            drawnData[i][j] = 0
        }
    }
}
