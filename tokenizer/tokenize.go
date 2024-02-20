package tokenizer

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

const (
	decimalPointCharacter = '.'
)

type Token struct {
	Kind     TokenKind
	Value    string
	Position int
}

type State int

const (
	DefaultState State = iota
	NumberState
	IdentifierState
)

func Tokenize(text string) ([]Token, error) {
	var tokens []Token
	stateCtx := newStateContext()
	for index, character := range text {
		switch {
		case unicode.IsSpace(character):
			token, err := stateCtx.createNumberToken(index)
			if err != nil && !errors.Is(err, errNoToken) {
				return nil, fmt.Errorf("unable to create a number token: %w", err)
			}
			if err == nil {
				tokens = append(tokens, token)
			}

			token, err = stateCtx.createIdentifierToken(index)
			if err != nil && !errors.Is(err, errNoToken) {
				return nil, fmt.Errorf("unable to create a identifier token: %w", err)
			}
			if err == nil {
				tokens = append(tokens, token)
			}

			continue
		case unicode.IsDigit(character) || character == decimalPointCharacter:
			if character == decimalPointCharacter {
				token, err := stateCtx.createIdentifierToken(index)
				if err != nil && !errors.Is(err, errNoToken) {
					return nil, fmt.Errorf("unable to create a identifier token: %w", err)
				}
				if err == nil {
					tokens = append(tokens, token)
				}
			}

			if err := stateCtx.addCharacterToNumber(index, character); err != nil {
				return nil, fmt.Errorf("unable to add a character to the number: %w", err)
			}
		case unicode.IsLetter(character) || character == '_':
			token, err := stateCtx.createNumberToken(index)
			if err != nil && !errors.Is(err, errNoToken) {
				return nil, fmt.Errorf("unable to create a number token: %w", err)
			}
			if err == nil {
				tokens = append(tokens, token)
			}

			stateCtx.addCharacterToIdentifier(character)
		case strings.ContainsRune("+-*/%^(),", character):
			token, err := stateCtx.createNumberToken(index)
			if err != nil && !errors.Is(err, errNoToken) {
				return nil, fmt.Errorf("unable to create a number token: %w", err)
			}
			if err == nil {
				tokens = append(tokens, token)
			}

			token, err = stateCtx.createIdentifierToken(index)
			if err != nil && !errors.Is(err, errNoToken) {
				return nil, fmt.Errorf("unable to create a identifier token: %w", err)
			}
			if err == nil {
				tokens = append(tokens, token)
			}

			kind, err := ParseTokenKind(character)
			if err != nil {
				return nil, fmt.Errorf("unable to parse a token kind at position %d: %w", index, err)
			}

			tokens = append(tokens, Token{
				Kind:     kind,
				Position: index,
			})
		default:
			return nil, fmt.Errorf("unknown character %q at position %d", character, index)
		}
	}

	endOfTextIndex := len(text)
	token, err := stateCtx.createNumberToken(endOfTextIndex)
	if err != nil && !errors.Is(err, errNoToken) {
		return nil, fmt.Errorf("unable to create a number token: %w", err)
	}
	if err == nil {
		tokens = append(tokens, token)
	}

	token, err = stateCtx.createIdentifierToken(endOfTextIndex)
	if err != nil && !errors.Is(err, errNoToken) {
		return nil, fmt.Errorf("unable to create a identifier token: %w", err)
	}
	if err == nil {
		tokens = append(tokens, token)
	}

	return tokens, nil
}
