package translator

import (
	"fmt"

	"github.com/rmaidveo/go-calculator/containers"
	"github.com/rmaidveo/go-calculator/tokenizer"
)

type CommandKind int

const (
	PushNumberCommand CommandKind = iota
	PushVariableCommand
	CallFunctionCommand
)

type Command struct {
	Kind     CommandKind
	Operand  string
	Position int
}

func Translate(tokens []tokenizer.Token, functions map[string]struct{}) ([]Command, error) {
	var commands []Command
	var tokenStack containers.Stack[tokenizer.Token]
	for _, token := range tokens {
		switch {
		case token.Kind == tokenizer.NumberToken:
			commands = append(commands, Command{
				Kind:     PushNumberCommand,
				Operand:  token.Value,
				Position: token.Position,
			})
		case token.Kind == tokenizer.IdentifierToken:
			if _, ok := functions[token.Value]; ok {
				tokenStack.Push(token)
				continue
			}

			commands = append(commands, Command{
				Kind:     PushVariableCommand,
				Operand:  token.Value,
				Position: token.Position,
			})
		case token.Kind.IsOperator():
			for {
				lastStackToken, ok := tokenStack.Pop()
				if !ok {
					break
				}
				if !lastStackToken.Kind.IsOperator() || lastStackToken.Kind.Precedence() <= token.Kind.Precedence() {
					tokenStack.Push(lastStackToken)
					break
				}

				commands = append(commands, Command{
					Kind:     CallFunctionCommand,
					Operand:  lastStackToken.Kind.String(),
					Position: lastStackToken.Position,
				})
			}

			tokenStack.Push(token)
		case token.Kind == tokenizer.LeftParenthesisToken:
			tokenStack.Push(token)
		case token.Kind == tokenizer.RightParenthesisToken:
			for {
				lastStackToken, ok := tokenStack.Pop()
				if !ok {
					break
				}
				if lastStackToken.Kind == tokenizer.LeftParenthesisToken {
					tokenStack.Push(lastStackToken)
					break
				}

				commands = append(commands, Command{
					Kind:     CallFunctionCommand,
					Operand:  lastStackToken.Kind.String(),
					Position: lastStackToken.Position,
				})
			}

			if _, ok := tokenStack.Pop(); !ok {
				return nil, fmt.Errorf("no left parenthesis is found, but a right parenthesis at position %d", token.Position)
			}

			lastStackToken, ok := tokenStack.Pop()
			if ok {
				if lastStackToken.Kind == tokenizer.IdentifierToken {
					commands = append(commands, Command{
						Kind:     CallFunctionCommand,
						Operand:  lastStackToken.Value,
						Position: lastStackToken.Position,
					})
				} else {
					tokenStack.Push(lastStackToken)
				}
			}
		case token.Kind == tokenizer.CommaToken:
			for {
				lastStackToken, ok := tokenStack.Pop()
				if !ok {
					break
				}
				if lastStackToken.Kind == tokenizer.LeftParenthesisToken {
					tokenStack.Push(lastStackToken)
					break
				}

				commands = append(commands, Command{
					Kind:     CallFunctionCommand,
					Operand:  lastStackToken.Kind.String(),
					Position: lastStackToken.Position,
				})
			}

			if tokenStack.IsEmpty() {
				return nil, fmt.Errorf("no left parenthesis is found, but a comma at position %d", token.Position)
			}
		}
	}

	for {
		lastStackToken, ok := tokenStack.Pop()
		if !ok {
			break
		}
		if lastStackToken.Kind == tokenizer.LeftParenthesisToken {
			return nil, fmt.Errorf("unexpected left parenthesis is found at position %d", lastStackToken.Position)
		}

		commands = append(commands, Command{
			Kind:     CallFunctionCommand,
			Operand:  lastStackToken.Kind.String(),
			Position: lastStackToken.Position,
		})
	}

	return commands, nil
}
