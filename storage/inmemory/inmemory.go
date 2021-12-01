package inmemory

import (
	"errors"
	"testSearchEngine/storage"
	"time"
)

type index struct {
	wordMap map[string]map[string][]document
}

func NewInMemoryStorage() *index {
	return &index{make(map[string]map[string][]document)}
}

type document struct {
	text        string
	createdTime *time.Time
	id          string
}

func (d document) GetID() string {
	return d.id
}

func (d document) GetCreatedTime() *time.Time {
	return d.createdTime
}

func (i *index) Add(datasetName string, tokens []string, docText, docId string) error {
	if _, ok := i.wordMap[datasetName]; !ok {
		i.wordMap[datasetName] = make(map[string][]document)
	}
	now := time.Now()
	for _, token := range tokens {
		if _, ok := i.wordMap[datasetName][token]; !ok {
			i.wordMap[datasetName][token] = make([]document, 0)
		}
		i.wordMap[datasetName][token] = append(i.wordMap[datasetName][token], document{
			text:        docText,
			createdTime: &now,
			id:          docId,
		})
	}
	return nil
}

func (i index) Search(datasetName string, tokens []string) ([]storage.Attributes, error) {
	if _, ok := i.wordMap[datasetName]; !ok {
		return nil, errors.New("dataset does not exists")
	}
	docList := make([]storage.Attributes, 0)
	for _, token := range tokens {
		if docs, ok := i.wordMap[datasetName][token]; ok {
			for _, doc := range docs {
				docList = append(docList, doc)
			}
		}
	}
	return docList, nil
}
