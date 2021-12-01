package tokenizer

import "strings"

type tokenizer struct {
	delim string
}

func NewTokenizer(delim string) *tokenizer {
	return &tokenizer{delim: delim}
}

func (t tokenizer) Process(doc string) []string {
	return strings.Split(doc, t.delim)
}
