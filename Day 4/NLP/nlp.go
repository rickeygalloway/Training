package nlp

import (
	"regexp"
	"strings"

	"git.shipt.com/nlp/stemmer"
)

var (
	// "Who's on first?" -> [Who s on first]
	wordRe = regexp.MustCompile(`[a-zA-Z]+`)
)

// Tokenize returns slice of tokens (lower case) found in text.
func Tokenize(text string) []string {
	words := wordRe.FindAllString(text, -1)
	var tokens []string
	for _, w := range words {
		token := strings.ToLower(w)
		token = stemmer.Stem(token)
		if len(token) != 0 {
			tokens = append(tokens, token)
		}
	}
	return tokens
}
