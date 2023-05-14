package main

import (
	"strings"

	"github.com/ilfey/brainfuck-interpreter/internal/app"
)

func main() {

	i := app.NewInterpreter(65535)
	c := app.NewConsole()

	c.Write("Type ;exit to exit the REPL")

	var line string
	for {
		line = c.Read()
		
		if strings.ToLower(line) == ";exit" {
			break
		}

		i.Execute(line)
	}
}
