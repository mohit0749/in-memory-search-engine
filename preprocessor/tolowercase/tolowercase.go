package tolowercase

import "strings"

type lowercase struct {
}

func NewLowerCasePreprocessor() lowercase {
	return lowercase{}
}

func (l lowercase) Process(doc string) []string {
	return []string{strings.ToLower(doc)}
}
