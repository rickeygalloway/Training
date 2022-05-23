package nlp

import (
	"regexp"
	"strings"
)

var (
	// "Who's on first?" -> [Who s on first]
	wordRe = regexp.MustCompile(`[a-zA-Z]+`)
)

/* Capital T in the Tokenize - means it can be used by anything using the nlp package.
If it was lower t, it would be scoped to local use only.

Below is how to do documentation  // [func name] description.

This will give this from the terminal window:

		rickey.galloway@TN40W7F43T nlp % go doc Tokenize
		package nlp // import "."

		func Tokenize(text string) []string
		Tokenize returns slice of tokens (lower case) found in text.


*you can use godoc external tool for documentation
*/

// Tokenize returns slice of tokens (lower case) found in text.
func Tokenize(text string) []string {
	words := wordRe.FindAllString(text, -1)
	var tokens []string
	for _, w := range words {
		token := strings.ToLower(w)
		//token = stemmer.Stem(token)	//Got an error here: undefined: stemmer, but I'm not concerned with errors right now
		if len(token) != 0 {
			tokens = append(tokens, token)
		}
	}
	return tokens
}

// TestFunc is another function that does very little.
func TestFunc(text string) string {
	return text
}
