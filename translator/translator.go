package translator

import (
	"fmt"

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
	var tokenStack []tokenizer.Token
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
				tokenStack = append(tokenStack, token)
				continue
			}

			commands = append(commands, Command{
				Kind:     PushVariableCommand,
				Operand:  token.Value,
				Position: token.Position,
			})
		case token.Kind.IsOperator():
			for len(tokenStack) != 0 {
				lastStackToken := tokenStack[len(tokenStack)-1]
				if !lastStackToken.Kind.IsOperator() || lastStackToken.Kind.Precedence() <= token.Kind.Precedence() {
					break
				}

				commands = append(commands, Command{
					Kind:     CallFunctionCommand,
					Operand:  lastStackToken.Kind.String(),
					Position: lastStackToken.Position,
				})

				tokenStack = tokenStack[:len(tokenStack)-1]
			}

			tokenStack = append(tokenStack, token)
		case token.Kind == tokenizer.LeftParenthesisToken:
			tokenStack = append(tokenStack, token)
		case token.Kind == tokenizer.RightParenthesisToken:
			for len(tokenStack) != 0 {
				lastStackToken := tokenStack[len(tokenStack)-1]
				if lastStackToken.Kind == tokenizer.LeftParenthesisToken {
					break
				}

				commands = append(commands, Command{
					Kind:     CallFunctionCommand,
					Operand:  lastStackToken.Kind.String(),
					Position: lastStackToken.Position,
				})

				tokenStack = tokenStack[:len(tokenStack)-1]
			}

			if len(tokenStack) == 0 {
				return nil, fmt.Errorf("no left parenthesis is found, but a right parenthesis at position %d", token.Position)
			}

			tokenStack = tokenStack[:len(tokenStack)-1]

			if len(tokenStack) != 0 {
				lastStackToken := tokenStack[len(tokenStack)-1]
				if lastStackToken.Kind == tokenizer.IdentifierToken {
					commands = append(commands, Command{
						Kind:     CallFunctionCommand,
						Operand:  lastStackToken.Value,
						Position: lastStackToken.Position,
					})

					tokenStack = tokenStack[:len(tokenStack)-1]
				}
			}
		case token.Kind == tokenizer.CommaToken:
			for len(tokenStack) != 0 {
				lastStackToken := tokenStack[len(tokenStack)-1]
				if lastStackToken.Kind == tokenizer.LeftParenthesisToken {
					break
				}

				commands = append(commands, Command{
					Kind:     CallFunctionCommand,
					Operand:  lastStackToken.Kind.String(),
					Position: lastStackToken.Position,
				})

				tokenStack = tokenStack[:len(tokenStack)-1]
			}

			if len(tokenStack) == 0 {
				return nil, fmt.Errorf("no left parenthesis is found, but a comma at position %d", token.Position)
			}
		}
	}

	for len(tokenStack) != 0 {
		lastStackToken := tokenStack[len(tokenStack)-1]
		if lastStackToken.Kind == tokenizer.LeftParenthesisToken {
			return nil, fmt.Errorf("unexpected left parenthesis is found at position %d", lastStackToken.Position)
		}

		commands = append(commands, Command{
			Kind:     CallFunctionCommand,
			Operand:  lastStackToken.Kind.String(),
			Position: lastStackToken.Position,
		})

		tokenStack = tokenStack[:len(tokenStack)-1]
	}

	return commands, nil
}
