package indexer

import (
	"testSearchEngine/idgenerator"
	"testSearchEngine/preprocessor"
	"testSearchEngine/ranker"
	"testSearchEngine/storage"
)

type indexer struct {
	preprocessors []preprocessor.Preprocessor
	store         storage.Storage
	idGenerator   idgenerator.Generator
	ranker        ranker.Ranker
}

func NewIndexer(store storage.Storage, ranker ranker.Ranker, idGenerator idgenerator.Generator, preprocessors ...preprocessor.Preprocessor) *indexer {
	pp := make([]preprocessor.Preprocessor, 0)
	for _, p := range preprocessors {
		pp = append(pp, p)
	}
	return &indexer{
		preprocessors: pp,
		store:         store,
		ranker:        ranker,
		idGenerator:   idGenerator,
	}
}

func (i *indexer) AddPreProcessor(pp preprocessor.Preprocessor) {
	i.preprocessors = append(i.preprocessors, pp)
}

func (i indexer) Insert(datasetName, doc string) error {
	data := i.preprocess(doc)
	id, err := i.idGenerator.Generate(datasetName, doc)
	if err != nil {
		return err
	}
	return i.store.Add(datasetName, data, doc, id)
}

func (i indexer) preprocess(doc string) []string {
	data := make([]string, 0)
	for _, pp := range i.preprocessors {
		data = pp.Process(doc)
	}
	return data
}

func (i indexer) Search(datasetName string, queryString string) ([]string, error) {
	data := i.preprocess(queryString)
	docIds, err := i.store.Search(datasetName, data)
	if err != nil {
		return []string{}, err
	}
	ranks, err := i.ranker.RankDocument(docIds)
	if err != nil {
		return []string{}, err
	}
	docs := make([]string, 0)
	for _, rd := range ranks {
		docs = append(docs, rd.GetDocId())
	}
	return docs, nil
}
