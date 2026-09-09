package lox

import (
	"fmt"
	"io"
)

// Punto de entrada del interprete. Recibe el código fuente y un io.Writer para imprimir la salida (consola o web).
func Run(source string, out io.Writer) {
	fmt.Fprintf(out, "angLOXg ejecutando (placeholder): %s\n", source)
}
