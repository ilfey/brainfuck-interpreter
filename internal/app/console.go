package app

import (
	"fmt"
)

type Console struct {}

func NewConsole() *Console {
	return &Console{}
}

func (c *Console) Read() (line string) {
	fmt.Scan(&line)
	return
}

func (c *Console) ReadNewLine() (line string) {
	fmt.Println()
	fmt.Scan(line)
	return
}

func (c *Console) Write(line string) {
	fmt.Println(line)
}

func (c *Console) WriteNewLine(line string) {
	fmt.Printf("\n%s\n", line)
}
