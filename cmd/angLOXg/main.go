package main

import (
	"fmt"
	"os"

	"github.com/FelipeAscencio/angLOXg/internal/lox"
)

func main() {
	fmt.Println("=== angLOXg CLI ===")
	lox.Run("print 'Hola FIUBA';", os.Stdout)
}
