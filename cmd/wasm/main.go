package main

import (
	"bytes"
	"syscall/js"

	"github.com/tu-usuario/angLOXg/internal/lox"
)

func runLoxWrapper(this js.Value, args []js.Value) any {
	if len(args) == 0 {
		return "Error: no hay código"
	}

	code := args[0].String()
	var buf bytes.Buffer
	lox.Run(code, &buf)
	return buf.String()
}

func main() {
	js.Global().Set("runLox", js.FuncOf(runLoxWrapper))
	select {}
}
