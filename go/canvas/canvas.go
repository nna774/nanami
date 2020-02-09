package canvas

import (
	"syscall/js"

	"github.com/nna774/mado/go/internal"
)

// Canvas is js canvas
type Canvas struct {
	jsValue js.Value
}

// NewCanvas creates Canvas
func NewCanvas(v js.Value) Canvas {
	return Canvas{jsValue: v}
}

// GetCanvas gets canvas
func GetCanvas(id string) Canvas {
	return NewCanvas(internal.GetElementByID(id))
}

// SetSize sets size
func (c *Canvas) SetSize(width, height int) {
	c.jsValue.Set("width", width)
	c.jsValue.Set("height", height)
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
