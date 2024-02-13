package tokenizer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenize(t *testing.T) {
	type args struct {
		text string
	}

	tests := []struct {
		name    string
		args    args
		want    []Token
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "success/+",
			args:    args{text: "+"},
			want:    []Token{{Kind: PlusToken, Position: 0}},
			wantErr: assert.NoError,
		},
		{
			name:    "success/-",
			args:    args{text: "-"},
			want:    []Token{{Kind: MinusToken, Position: 0}},
			wantErr: assert.NoError,
		},
		{
			name:    "success/*",
			args:    args{text: "*"},
			want:    []Token{{Kind: AsteriskToken, Position: 0}},
			wantErr: assert.NoError,
		},
		{
			name:    "success//",
			args:    args{text: "/"},
			want:    []Token{{Kind: SlashToken, Position: 0}},
			wantErr: assert.NoError,
		},
		{
			name:    "success/%",
			args:    args{text: "%"},
			want:    []Token{{Kind: PercentToken, Position: 0}},
			wantErr: assert.NoError,
		},
		{
			name:    "success/^",
			args:    args{text: "^"},
			want:    []Token{{Kind: ExponentiationToken, Position: 0}},
			wantErr: assert.NoError,
		},
		{
			name:    "success/(",
			args:    args{text: "("},
			want:    []Token{{Kind: LeftParenthesisToken, Position: 0}},
			wantErr: assert.NoError,
		},
		{
			name:    "success/)",
			args:    args{text: ")"},
			want:    []Token{{Kind: RightParenthesisToken, Position: 0}},
			wantErr: assert.NoError,
		},
		{
			name:    "success/,",
			args:    args{text: ","},
			want:    []Token{{Kind: CommaToken, Position: 0}},
			wantErr: assert.NoError,
		},
		{
			name: "success/all punctuation",
			args: args{text: "+-*/%^(),"},
			want: []Token{
				{Kind: PlusToken, Position: 0},
				{Kind: MinusToken, Position: 1},
				{Kind: AsteriskToken, Position: 2},
				{Kind: SlashToken, Position: 3},
				{Kind: PercentToken, Position: 4},
				{Kind: ExponentiationToken, Position: 5},
				{Kind: LeftParenthesisToken, Position: 6},
				{Kind: RightParenthesisToken, Position: 7},
				{Kind: CommaToken, Position: 8},
			},
			wantErr: assert.NoError,
		},
		{
			name: "success/all punctuation with spaces",
			args: args{text: "+ - * / % ^ ( ) ,"},
			want: []Token{
				{Kind: PlusToken, Position: 0},
				{Kind: MinusToken, Position: 2},
				{Kind: AsteriskToken, Position: 4},
				{Kind: SlashToken, Position: 6},
				{Kind: PercentToken, Position: 8},
				{Kind: ExponentiationToken, Position: 10},
				{Kind: LeftParenthesisToken, Position: 12},
				{Kind: RightParenthesisToken, Position: 14},
				{Kind: CommaToken, Position: 16},
			},
			wantErr: assert.NoError,
		},
		{
			name:    "error",
			args:    args{text: "@"},
			want:    nil,
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Tokenize(tt.args.text)

			assert.Equal(t, tt.want, got)
			tt.wantErr(t, err)
		})
	}
}
