package translator

import (
	"testing"

	"github.com/rmaidveo/go-calculator/tokenizer"
	"github.com/stretchr/testify/assert"
)

func TestTranslate(t *testing.T) {
	type args struct {
		tokens    []tokenizer.Token
		functions map[string]struct{}
	}

	tests := []struct {
		name    string
		args    args
		want    []Command
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "success/single number",
			args:    args{tokens: []tokenizer.Token{{Kind: tokenizer.NumberToken, Value: "23", Position: 42}}},
			want:    []Command{{Kind: PushNumberCommand, Operand: "23", Position: 42}},
			wantErr: assert.NoError,
		},
		{
			name:    "success/single variable",
			args:    args{tokens: []tokenizer.Token{{Kind: tokenizer.IdentifierToken, Value: "x", Position: 42}}},
			want:    []Command{{Kind: PushVariableCommand, Operand: "x", Position: 42}},
			wantErr: assert.NoError,
		},
		{
			name: "success/single operator",
			args: args{
				tokens: []tokenizer.Token{
					{Kind: tokenizer.NumberToken, Value: "12", Position: 112},
					{Kind: tokenizer.PlusToken, Value: "+", Position: 120},
					{Kind: tokenizer.NumberToken, Value: "23", Position: 123},
				},
			},
			want: []Command{
				{Kind: PushNumberCommand, Operand: "12", Position: 112},
				{Kind: PushNumberCommand, Operand: "23", Position: 123},
				{Kind: CallFunctionCommand, Operand: "+", Position: 120},
			},
			wantErr: assert.NoError,
		},
		{
			name: "success/multiple operators",
			args: args{
				tokens: []tokenizer.Token{
					{Kind: tokenizer.NumberToken, Value: "12", Position: 112},
					{Kind: tokenizer.PlusToken, Value: "+", Position: 120},
					{Kind: tokenizer.NumberToken, Value: "23", Position: 123},
					{Kind: tokenizer.AsteriskToken, Value: "*", Position: 140},
					{Kind: tokenizer.NumberToken, Value: "42", Position: 142},
				},
			},
			want: []Command{
				{Kind: PushNumberCommand, Operand: "12", Position: 112},
				{Kind: PushNumberCommand, Operand: "23", Position: 123},
				{Kind: PushNumberCommand, Operand: "42", Position: 142},
				{Kind: CallFunctionCommand, Operand: "*", Position: 140},
				{Kind: CallFunctionCommand, Operand: "+", Position: 120},
			},
			wantErr: assert.NoError,
		},
		{
			name: "success/multiple operators (in a different order)",
			args: args{
				tokens: []tokenizer.Token{
					{Kind: tokenizer.NumberToken, Value: "12", Position: 112},
					{Kind: tokenizer.AsteriskToken, Value: "*", Position: 120},
					{Kind: tokenizer.NumberToken, Value: "23", Position: 123},
					{Kind: tokenizer.PlusToken, Value: "+", Position: 140},
					{Kind: tokenizer.NumberToken, Value: "42", Position: 142},
				},
			},
			want: []Command{
				{Kind: PushNumberCommand, Operand: "12", Position: 112},
				{Kind: PushNumberCommand, Operand: "23", Position: 123},
				{Kind: CallFunctionCommand, Operand: "*", Position: 120},
				{Kind: PushNumberCommand, Operand: "42", Position: 142},
				{Kind: CallFunctionCommand, Operand: "+", Position: 140},
			},
			wantErr: assert.NoError,
		},
		{
			name: "success/multiple operators (with parentheses)",
			args: args{
				tokens: []tokenizer.Token{
					{Kind: tokenizer.LeftParenthesisToken, Value: "(", Position: 110},
					{Kind: tokenizer.NumberToken, Value: "12", Position: 112},
					{Kind: tokenizer.PlusToken, Value: "+", Position: 120},
					{Kind: tokenizer.NumberToken, Value: "23", Position: 123},
					{Kind: tokenizer.RightParenthesisToken, Value: ")", Position: 130},
					{Kind: tokenizer.AsteriskToken, Value: "*", Position: 140},
					{Kind: tokenizer.NumberToken, Value: "42", Position: 142},
				},
			},
			want: []Command{
				{Kind: PushNumberCommand, Operand: "12", Position: 112},
				{Kind: PushNumberCommand, Operand: "23", Position: 123},
				{Kind: CallFunctionCommand, Operand: "+", Position: 120},
				{Kind: PushNumberCommand, Operand: "42", Position: 142},
				{Kind: CallFunctionCommand, Operand: "*", Position: 140},
			},
			wantErr: assert.NoError,
		},
		{
			name: "success/multiple operators (with parentheses in a different order)",
			args: args{
				tokens: []tokenizer.Token{
					{Kind: tokenizer.NumberToken, Value: "42", Position: 100},
					{Kind: tokenizer.AsteriskToken, Value: "*", Position: 105},
					{Kind: tokenizer.LeftParenthesisToken, Value: "(", Position: 110},
					{Kind: tokenizer.NumberToken, Value: "12", Position: 112},
					{Kind: tokenizer.PlusToken, Value: "+", Position: 120},
					{Kind: tokenizer.NumberToken, Value: "23", Position: 123},
					{Kind: tokenizer.RightParenthesisToken, Value: ")", Position: 130},
				},
			},
			want: []Command{
				{Kind: PushNumberCommand, Operand: "42", Position: 100},
				{Kind: PushNumberCommand, Operand: "12", Position: 112},
				{Kind: PushNumberCommand, Operand: "23", Position: 123},
				{Kind: CallFunctionCommand, Operand: "+", Position: 120},
				{Kind: CallFunctionCommand, Operand: "*", Position: 105},
			},
			wantErr: assert.NoError,
		},
		{
			name: "success/function call",
			args: args{
				tokens: []tokenizer.Token{
					{Kind: tokenizer.IdentifierToken, Value: "sin", Position: 105},
					{Kind: tokenizer.LeftParenthesisToken, Value: "(", Position: 110},
					{Kind: tokenizer.NumberToken, Value: "12", Position: 112},
					{Kind: tokenizer.PlusToken, Value: "+", Position: 120},
					{Kind: tokenizer.NumberToken, Value: "23", Position: 123},
					{Kind: tokenizer.RightParenthesisToken, Value: ")", Position: 130},
				},
				functions: map[string]struct{}{"sin": {}},
			},
			want: []Command{
				{Kind: PushNumberCommand, Operand: "12", Position: 112},
				{Kind: PushNumberCommand, Operand: "23", Position: 123},
				{Kind: CallFunctionCommand, Operand: "+", Position: 120},
				{Kind: CallFunctionCommand, Operand: "sin", Position: 105},
			},
			wantErr: assert.NoError,
		},
		{
			name: "success/function call (with several operands)",
			args: args{
				tokens: []tokenizer.Token{
					{Kind: tokenizer.IdentifierToken, Value: "sin", Position: 105},
					{Kind: tokenizer.LeftParenthesisToken, Value: "(", Position: 110},
					{Kind: tokenizer.NumberToken, Value: "5", Position: 112},
					{Kind: tokenizer.PlusToken, Value: "+", Position: 120},
					{Kind: tokenizer.NumberToken, Value: "12", Position: 123},
					{Kind: tokenizer.CommaToken, Value: ",", Position: 125},
					{Kind: tokenizer.NumberToken, Value: "23", Position: 130},
					{Kind: tokenizer.AsteriskToken, Value: "*", Position: 135},
					{Kind: tokenizer.NumberToken, Value: "42", Position: 142},
					{Kind: tokenizer.RightParenthesisToken, Value: ")", Position: 150},
				},
				functions: map[string]struct{}{"sin": {}},
			},
			want: []Command{
				{Kind: PushNumberCommand, Operand: "5", Position: 112},
				{Kind: PushNumberCommand, Operand: "12", Position: 123},
				{Kind: CallFunctionCommand, Operand: "+", Position: 120},
				{Kind: PushNumberCommand, Operand: "23", Position: 130},
				{Kind: PushNumberCommand, Operand: "42", Position: 142},
				{Kind: CallFunctionCommand, Operand: "*", Position: 135},
				{Kind: CallFunctionCommand, Operand: "sin", Position: 105},
			},
			wantErr: assert.NoError,
		},
		{
			name: "error/no left parenthesis is found",
			args: args{
				tokens: []tokenizer.Token{
					{Kind: tokenizer.NumberToken, Value: "12", Position: 112},
					{Kind: tokenizer.PlusToken, Value: "+", Position: 120},
					{Kind: tokenizer.NumberToken, Value: "23", Position: 123},
					{Kind: tokenizer.RightParenthesisToken, Value: ")", Position: 130},
					{Kind: tokenizer.AsteriskToken, Value: "*", Position: 140},
					{Kind: tokenizer.NumberToken, Value: "42", Position: 142},
				},
			},
			want:    nil,
			wantErr: assert.Error,
		},
		{
			name: "error/unexpected left parenthesis",
			args: args{
				tokens: []tokenizer.Token{
					{Kind: tokenizer.LeftParenthesisToken, Value: "(", Position: 110},
					{Kind: tokenizer.NumberToken, Value: "12", Position: 112},
					{Kind: tokenizer.PlusToken, Value: "+", Position: 120},
					{Kind: tokenizer.NumberToken, Value: "23", Position: 123},
					{Kind: tokenizer.AsteriskToken, Value: "*", Position: 140},
					{Kind: tokenizer.NumberToken, Value: "42", Position: 142},
				},
			},
			want:    nil,
			wantErr: assert.Error,
		},
		{
			name: "error/function call/no left parenthesis is found",
			args: args{
				tokens: []tokenizer.Token{
					{Kind: tokenizer.IdentifierToken, Value: "sin", Position: 105},
					{Kind: tokenizer.NumberToken, Value: "5", Position: 112},
					{Kind: tokenizer.PlusToken, Value: "+", Position: 120},
					{Kind: tokenizer.NumberToken, Value: "12", Position: 123},
					{Kind: tokenizer.CommaToken, Value: ",", Position: 125},
					{Kind: tokenizer.NumberToken, Value: "23", Position: 130},
					{Kind: tokenizer.AsteriskToken, Value: "*", Position: 135},
					{Kind: tokenizer.NumberToken, Value: "42", Position: 142},
					{Kind: tokenizer.RightParenthesisToken, Value: ")", Position: 150},
				},
				functions: map[string]struct{}{"sin": {}},
			},
			want:    nil,
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Translate(tt.args.tokens, tt.args.functions)

			assert.Equal(t, tt.want, got)
			tt.wantErr(t, err)
		})
	}
}
