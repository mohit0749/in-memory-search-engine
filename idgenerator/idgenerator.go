package idgenerator

import "sync"

var defaultGenerator Generator
var once sync.Once

type Generator interface {
	Generate(datasetName, doc string) (string, error)
}

func Init(g Generator) {
	once.Do(func() {
		defaultGenerator = g
	})
}

func Get() Generator {
	return defaultGenerator
}
