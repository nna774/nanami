package main

import (
	"fmt"

	"github.com/nna774/mado/go/canvas"
	"github.com/nna774/mado/go/debug"
)

func main() {
	fmt.Println("Hello, WebAssembly!")
	canvas := canvas.GetCanvas("canvas")
	canvas.SetSize(256, 240)
	canvas.FillTestRect()
	debugger := debug.NewDebugger("debug")
	debugger.Log("debug messsage!")
	debugger.Log("second debug messsage!")
}
