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
			name:    "success/number/integer",
			args:    args{text: "23"},
			want:    []Token{{Kind: NumberToken, Value: "23", Position: 0}},
			wantErr: assert.NoError,
		},
		{
			name:    "success/number/float",
			args:    args{text: "23.5"},
			want:    []Token{{Kind: NumberToken, Value: "23.5", Position: 0}},
			wantErr: assert.NoError,
		},
		{
			name:    "success/number/float/starts with a decinal point",
			args:    args{text: ".23"},
			want:    []Token{{Kind: NumberToken, Value: ".23", Position: 0}},
			wantErr: assert.NoError,
		},
		{
			name:    "success/number/float/ends with a decinal point",
			args:    args{text: "23."},
			want:    []Token{{Kind: NumberToken, Value: "23.", Position: 0}},
			wantErr: assert.NoError,
		},
		{
			name: "success/number/several numbers separated by punctuation",
			args: args{text: "12+34-56*78/90%12^34(56)78,90"},
			want: []Token{
				{Kind: NumberToken, Value: "12", Position: 0},
				{Kind: PlusToken, Position: 2},
				{Kind: NumberToken, Value: "34", Position: 3},
				{Kind: MinusToken, Position: 5},
				{Kind: NumberToken, Value: "56", Position: 6},
				{Kind: AsteriskToken, Position: 8},
				{Kind: NumberToken, Value: "78", Position: 9},
				{Kind: SlashToken, Position: 11},
				{Kind: NumberToken, Value: "90", Position: 12},
				{Kind: PercentToken, Position: 14},
				{Kind: NumberToken, Value: "12", Position: 15},
				{Kind: ExponentiationToken, Position: 17},
				{Kind: NumberToken, Value: "34", Position: 18},
				{Kind: LeftParenthesisToken, Position: 20},
				{Kind: NumberToken, Value: "56", Position: 21},
				{Kind: RightParenthesisToken, Position: 23},
				{Kind: NumberToken, Value: "78", Position: 24},
				{Kind: CommaToken, Position: 26},
				{Kind: NumberToken, Value: "90", Position: 27},
			},
			wantErr: assert.NoError,
		},
		{
			name: "success/number/several numbers separated by spaces/integers",
			args: args{text: "12 34 56 78 90 12 34 56 78 90"},
			want: []Token{
				{Kind: NumberToken, Value: "12", Position: 0},
				{Kind: NumberToken, Value: "34", Position: 3},
				{Kind: NumberToken, Value: "56", Position: 6},
				{Kind: NumberToken, Value: "78", Position: 9},
				{Kind: NumberToken, Value: "90", Position: 12},
				{Kind: NumberToken, Value: "12", Position: 15},
				{Kind: NumberToken, Value: "34", Position: 18},
				{Kind: NumberToken, Value: "56", Position: 21},
				{Kind: NumberToken, Value: "78", Position: 24},
				{Kind: NumberToken, Value: "90", Position: 27},
			},
			wantErr: assert.NoError,
		},
		{
			name: "success/number/several numbers separated by spaces/floats",
			args: args{text: "12.5 34.5 56.5 78.5 90.5 12.5 34.5 56.5 78.5 90.5"},
			want: []Token{
				{Kind: NumberToken, Value: "12.5", Position: 0},
				{Kind: NumberToken, Value: "34.5", Position: 5},
				{Kind: NumberToken, Value: "56.5", Position: 10},
				{Kind: NumberToken, Value: "78.5", Position: 15},
				{Kind: NumberToken, Value: "90.5", Position: 20},
				{Kind: NumberToken, Value: "12.5", Position: 25},
				{Kind: NumberToken, Value: "34.5", Position: 30},
				{Kind: NumberToken, Value: "56.5", Position: 35},
				{Kind: NumberToken, Value: "78.5", Position: 40},
				{Kind: NumberToken, Value: "90.5", Position: 45},
			},
			wantErr: assert.NoError,
		},
		{
			name:    "success/identifier/without an underscore",
			args:    args{text: "test"},
			want:    []Token{{Kind: IdentifierToken, Value: "test", Position: 0}},
			wantErr: assert.NoError,
		},
		{
			name:    "success/identifier/with an underscore",
			args:    args{text: "_test"},
			want:    []Token{{Kind: IdentifierToken, Value: "_test", Position: 0}},
			wantErr: assert.NoError,
		},
		{
			name:    "success/identifier/with a number in the middle",
			args:    args{text: "te23st"},
			want:    []Token{{Kind: IdentifierToken, Value: "te23st", Position: 0}},
			wantErr: assert.NoError,
		},
		{
			name:    "success/identifier/with a number at the end/integer",
			args:    args{text: "test23"},
			want:    []Token{{Kind: IdentifierToken, Value: "test23", Position: 0}},
			wantErr: assert.NoError,
		},
		{
			name: "success/identifier/with a number at the end/float",
			args: args{text: "test.23"},
			want: []Token{
				{Kind: IdentifierToken, Value: "test", Position: 0},
				{Kind: NumberToken, Value: ".23", Position: 4},
			},
			wantErr: assert.NoError,
		},
		{
			name: "success/identifier/with a number at the beginning",
			args: args{text: "23test"},
			want: []Token{
				{Kind: NumberToken, Value: "23", Position: 0},
				{Kind: IdentifierToken, Value: "test", Position: 2},
			},
			wantErr: assert.NoError,
		},
		{
			name: "success/identifier/several identifiers separated by punctuation",
			args: args{text: "xy+xy-xy*xy/xy%xy^xy(xy)xy,xy"},
			want: []Token{
				{Kind: IdentifierToken, Value: "xy", Position: 0},
				{Kind: PlusToken, Position: 2},
				{Kind: IdentifierToken, Value: "xy", Position: 3},
				{Kind: MinusToken, Position: 5},
				{Kind: IdentifierToken, Value: "xy", Position: 6},
				{Kind: AsteriskToken, Position: 8},
				{Kind: IdentifierToken, Value: "xy", Position: 9},
				{Kind: SlashToken, Position: 11},
				{Kind: IdentifierToken, Value: "xy", Position: 12},
				{Kind: PercentToken, Position: 14},
				{Kind: IdentifierToken, Value: "xy", Position: 15},
				{Kind: ExponentiationToken, Position: 17},
				{Kind: IdentifierToken, Value: "xy", Position: 18},
				{Kind: LeftParenthesisToken, Position: 20},
				{Kind: IdentifierToken, Value: "xy", Position: 21},
				{Kind: RightParenthesisToken, Position: 23},
				{Kind: IdentifierToken, Value: "xy", Position: 24},
				{Kind: CommaToken, Position: 26},
				{Kind: IdentifierToken, Value: "xy", Position: 27},
			},
			wantErr: assert.NoError,
		},
		{
			name: "success/identifier/several identifiers separated by spaces",
			args: args{text: "xy xy xy xy xy xy xy xy xy xy"},
			want: []Token{
				{Kind: IdentifierToken, Value: "xy", Position: 0},
				{Kind: IdentifierToken, Value: "xy", Position: 3},
				{Kind: IdentifierToken, Value: "xy", Position: 6},
				{Kind: IdentifierToken, Value: "xy", Position: 9},
				{Kind: IdentifierToken, Value: "xy", Position: 12},
				{Kind: IdentifierToken, Value: "xy", Position: 15},
				{Kind: IdentifierToken, Value: "xy", Position: 18},
				{Kind: IdentifierToken, Value: "xy", Position: 21},
				{Kind: IdentifierToken, Value: "xy", Position: 24},
				{Kind: IdentifierToken, Value: "xy", Position: 27},
			},
			wantErr: assert.NoError,
		},
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
			name:    "error/number/duplicate decimal point",
			args:    args{text: "2.3.5"},
			want:    nil,
			wantErr: assert.Error,
		},
		{
			name:    "error/number/has only a decimal point",
			args:    args{text: "."},
			want:    nil,
			wantErr: assert.Error,
		},
		{
			name:    "error/unknown character",
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
