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
			name:    "single number",
			args:    args{tokens: []tokenizer.Token{{Kind: tokenizer.NumberToken, Value: "23", Position: 42}}},
			want:    []Command{{Kind: PushNumberCommand, Operand: "23", Position: 42}},
			wantErr: assert.NoError,
		},
		{
			name: "single operator",
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
			name: "multiple operators",
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
			name: "multiple operators (in a different order)",
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Translate(tt.args.tokens, tt.args.functions)

			assert.Equal(t, tt.want, got)
			tt.wantErr(t, err)
		})
	}
}