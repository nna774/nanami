package canvas

import (
	"image"
	"image/color"
	"syscall/js"

	"github.com/nna774/mado/go/internal"
)

// Canvas is js canvas
type Canvas struct {
	jsValue   js.Value
	rectangle image.Rectangle
}

func newCanvas(v js.Value, r image.Rectangle) Canvas {
	return Canvas{jsValue: v, rectangle: r}
}

// GetCanvas gets canvas
func GetCanvas(id string) Canvas {
	canvas := internal.GetElementByID(id)
	width := canvas.Get("width").Int()
	height := canvas.Get("height").Int()
	rectangle := image.Rectangle{
		Min: image.Point{},
		Max: image.Point{
			X: width,
			Y: height,
		},
	}
	return newCanvas(canvas, rectangle)
}

// SetSize sets size
func (c *Canvas) SetSize(width, height int) {
	c.jsValue.Set("width", width)
	c.jsValue.Set("height", height)
	c.rectangle.Max.X = width
	c.rectangle.Max.Y = height
}

// Context gets canvs context
func (c *Canvas) Context() js.Value {
	return c.jsValue.Call("getContext", "2d")
}

// FillTestRect is test pattern
func (c *Canvas) FillTestRect() {
	ctx := c.Context()
	ctx.Set("fillStyle", "green")
	ctx.Call("fillRect", 10, 10, 150, 100)
}

// ColorModel is part of Image interface
func (c *Canvas) ColorModel() color.Model {
	return color.NRGBAModel
}

// Bounds is part of Image interface
func (c *Canvas) Bounds() image.Rectangle {
	return c.rectangle
}

// At is part of Image interface
func (c *Canvas) At(x, y int) color.Color {
	ctx := c.Context()
	imageData := ctx.Call("getImageData", x, y, 1, 1)
	data := imageData.Get("data")
	r := uint8(data.Index(0).Int())
	g := uint8(data.Index(1).Int())
	b := uint8(data.Index(2).Int())
	a := uint8(data.Index(3).Int())
	return color.RGBA{r, g, b, a}
}
