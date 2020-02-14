package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"

	"github.com/nna774/mado/pkg/canvas"
	"github.com/nna774/mado/pkg/debug"
)

func main() {
	fmt.Println("Hello, WebAssembly!")
	canvas := canvas.GetCanvas("canvas")
	canvas.SetSize(256, 240)
	canvas.FillTestRect()
	debugger := debug.NewDebugger("debug")
	debugger.Log("debug messsage!")
	debugger.Log("second debug messsage!")

	var i image.Image
	i = canvas

	p := i.At(11, 11)
	debugger.Log(fmt.Sprintf("p! : %v", p))

	var d draw.Image
	d = canvas

	d.Set(15, 15, color.Black)
	p = i.At(15, 15)
	debugger.Log(fmt.Sprintf("p2! : %v", p))
}
