package autoincrementor

import "strconv"

type autoIncrement struct {
	ids map[string]int
}

func NewAutoIncrementGenerator() *autoIncrement {
	return &autoIncrement{make(map[string]int)}
}

func (ai autoIncrement) Generate(datasetName, doc string) (string, error) {
	ai.ids[datasetName]++
	return strconv.Itoa(ai.ids[datasetName]), nil
}
