package evaluator

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/rmaidveo/go-calculator/containers"
	"github.com/rmaidveo/go-calculator/translator"
)

type Function struct {
	Arity   int
	Handler func(arguments []float64) (float64, error)
}

func Evaluate(
	commands []translator.Command,
	variables map[string]float64,
	functions map[string]Function,
) (float64, error) {
	var numberStack containers.Stack[float64]
	for _, command := range commands {
		switch command.Kind {
		case translator.PushNumberCommand:
			number, err := strconv.ParseFloat(command.Operand, 64)
			if err != nil {
				return 0, fmt.Errorf("unable to parse the number at position %d: %w", command.Position, err)
			}

			numberStack.Push(number)
		case translator.PushVariableCommand:
			number, ok := variables[command.Operand]
			if !ok {
				return 0, fmt.Errorf("unknown variable %q at position %d", command.Operand, command.Position)
			}

			numberStack.Push(number)
		case translator.CallFunctionCommand:
			function, ok := functions[command.Operand]
			if !ok {
				return 0, fmt.Errorf("unknown function %q at position %d", command.Operand, command.Position)
			}

			var arguments []float64
			for argumentIndex := 0; argumentIndex < function.Arity; argumentIndex++ {
				number, ok := numberStack.Pop()
				if !ok {
					return 0, fmt.Errorf("number stack is empty for argument #%d at position %d", argumentIndex, command.Position)
				}

				arguments = append(arguments, number)
			}

			reverseSlice(arguments)

			number, err := function.Handler(arguments)
			if err != nil {
				return 0, fmt.Errorf("unable to call the function %q at position %d: %w", command.Operand, command.Position, err)
			}

			numberStack.Push(number)
		}
	}

	result, ok := numberStack.Pop()
	if !ok {
		return 0, errors.New("number stack is empty")
	}

	return result, nil
}

func reverseSlice[T any](slice []T) {
	for i := 0; i < len(slice)/2; i++ {
		j := (len(slice) - 1) - i
		slice[i], slice[j] = slice[j], slice[i]
	}
}
