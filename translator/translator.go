package translator

import "github.com/rmaidveo/go-calculator/tokenizer"

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
		}
	}

	for len(tokenStack) != 0 {
		lastStackToken := tokenStack[len(tokenStack)-1]

		commands = append(commands, Command{
			Kind:     CallFunctionCommand,
			Operand:  lastStackToken.Kind.String(),
			Position: lastStackToken.Position,
		})

		tokenStack = tokenStack[:len(tokenStack)-1]
	}

	return commands, nil
}
