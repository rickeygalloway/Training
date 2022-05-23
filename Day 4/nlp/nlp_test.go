package nlp //<- same package, so now we have access to internal functions
import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTokenize(t *testing.T) {
	text := "What's on second?"
	expected := []string{"what", "s", "on", "second"}
	tokens := Tokenize(text)

	//this uses require.Equal(t, expected, tokens)
	//you can run: go mod tidy to pull in reference and dependencies - see go.mod file
	require.Equal(t, expected, tokens)

	//if !reflect.DeepEqual(expected, tokens) {
	//	t.Fatalf("%#v-> %v", expected, tokens)
	//}

}

//anonymous struct
var tokenizeCases = []struct {
	text   string
	tokens []string
}{
	{"Who's on first?", []string{"who", "s", "on", "first"}},
	{"What's on second?", []string{"what", "s", "on", "second"}},
	{"", nil},
}

func TestTokenizeMany(t *testing.T) {
	for _, tc := range tokenizeCases {
		t.Run(tc.text, func(t *testing.T) {
			tokens := Tokenize(tc.text)
			require.Equal(t, tokens, tokens)
			//if !reflect.DeepEqual(tc.tokens, tokens) {
			//	t.Fatalf("%#v-> %v", tc.tokens, tokens)
			//}
		})
	}
}

/*
=== RUN   TestTokenizeMany
=== RUN   TestTokenizeMany/Who's_on_first?
=== RUN   TestTokenizeMany/What's_on_second?
=== RUN   TestTokenizeMany/#00
--- PASS: TestTokenizeMany (0.00s)
    --- PASS: TestTokenizeMany/Who's_on_first? (0.00s)
    --- PASS: TestTokenizeMany/What's_on_second? (0.00s)
    --- PASS: TestTokenizeMany/#00 (0.00s)
PASS
*/
