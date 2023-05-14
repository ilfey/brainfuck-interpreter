package app

import (
	"bufio"
	"fmt"
	"os"
)

type instructions struct {
	Output           func() error
	Input            func() error
	Increment        func() error
	Decrement        func() error
	IncrementPointer func() error
	DecrementPointer func() error
	Begin            func() error
	End              func() error
}

type Interpreter struct {
	*Console
	MemorySize   uint16
	Memory       []uint16
	JumpStack    []uint16
	Pointer      uint16
	NestingLevel uint16
	Instructions *instructions
	*Scanner
}

func NewInterpreter(memSize uint16) *Interpreter {
	i := &Interpreter{
		Console:    NewConsole(),
		MemorySize: memSize,
		Memory:     make([]uint16, memSize),
		Scanner:    NewScanner(),
	}

	i.Instructions = i.newInstructions()

	return i
}

func (i *Interpreter) newInstructions() *instructions {
	return &instructions{
		Output:           i.createOutputInstruction(),
		Input:            i.createInputInstruction(),
		Increment:        i.createIncrementInstruction(),
		Decrement:        i.createDecrementInstruction(),
		IncrementPointer: i.createIncrementPointerInstruction(),
		DecrementPointer: i.createDecrementPointerInstruction(),
		Begin:            i.createBeginInstruction(),
		End:              i.createEndInstruction(),
	}
}

func (i *Interpreter) Execute(code string) error {
	tokens := i.Scanner.AssignLine(code)

	return i.Interpretate(tokens)
}

func (i *Interpreter) Interpretate(tokens []Token) error {
	var err error
	for _, token := range tokens {
		switch token {
		case OUTPUT:
			err = i.Instructions.Output()
		case INPUT:
			err = i.Instructions.Input()
		case INCREMENT:
			err = i.Instructions.Increment()
		case DECREMENT:
			err = i.Instructions.Decrement()
		case INCREMENT_POINTER:
			err = i.Instructions.IncrementPointer()
		case DECREMENT_POINTER:
			err = i.Instructions.DecrementPointer()
		case BEGIN:
			err = i.Instructions.Begin()
		case END:
			err = i.Instructions.End()
		default:
			continue
		}

		if err != nil {
			return err
		}
	}

	return nil
}

func (i *Interpreter) createOutputInstruction() func() error {
	return func() error {
		i.Console.Write(fmt.Sprintf("%c", i.Memory[i.Pointer]))

		return nil
	}
}

func (i *Interpreter) createInputInstruction() func() error {
	reader := bufio.NewReader(os.Stdin)
	return func() error {
		b, err := reader.ReadByte()
		if err != nil {
			return err
		}

		i.Memory[i.Pointer] = uint16(b)

		return nil
	}
}

func (i *Interpreter) createIncrementInstruction() func() error {
	return func() error {
		i.Memory[i.Pointer]++

		return nil
	}
}

func (i *Interpreter) createDecrementInstruction() func() error {
	return func() error {
		i.Memory[i.Pointer]--

		return nil
	}
}

func (i *Interpreter) createIncrementPointerInstruction() func() error {
	return func() error {
		i.Pointer++

		return nil
	}
}

func (i *Interpreter) createDecrementPointerInstruction() func() error {
	return func() error {
		i.Pointer--

		return nil
	}
}

// ... ... ... .. ... ... 0 [ 1 ... ... 1 [ 2 .. .. 2 ] 1 ... ... .1 ] 0 ....... . . .. . . . . ..

func (i *Interpreter) createBeginInstruction() func() error {
	return func() error {
		if i.Memory[i.Pointer] != 0 {
			i.NestingLevel++
		}

		// TODO implement begin instruction

		return nil
	}
}

func (i *Interpreter) createEndInstruction() func() error {
	return func() error {
		if i.Memory[i.Pointer] == 0 {
			i.NestingLevel--
		}

		// TODO implement back to BEGIN

		return nil
	}
}
