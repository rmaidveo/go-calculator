package calculator

import (
	"fmt"

	"github.com/rmaidveo/go-calculator/evaluator"
	"github.com/rmaidveo/go-calculator/tokenizer"
	"github.com/rmaidveo/go-calculator/translator"
)

func Calculate(
	text string,
	variables map[string]float64,
	functions map[string]evaluator.Function,
) (float64, error) {
	tokens, err := tokenizer.Tokenize(text)
	if err != nil {
		return 0, fmt.Errorf("unable to tokenize: %w", err)
	}

	functionNames := make(map[string]struct{}, len(functions))
	for functionName := range functions {
		functionNames[functionName] = struct{}{}
	}

	commands, err := translator.Translate(tokens, functionNames)
	if err != nil {
		return 0, fmt.Errorf("unable to translate: %w", err)
	}

	result, err := evaluator.Evaluate(commands, variables, functions)
	if err != nil {
		return 0, fmt.Errorf("unable to evaluate: %w", err)
	}

	return result, nil
}
