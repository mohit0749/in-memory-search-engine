package ranker

import (
	"sync"
	"testSearchEngine/storage"
)

var defaultRanker Ranker
var once sync.Once

type Ranker interface {
	RankDocument(docIds []storage.Attributes) ([]Rank, error)
}

type Rank interface {
	GetDocId() string
	GetRank() int
}

func Init(r Ranker) {
	once.Do(func() {
		defaultRanker = r
	})
}

func Get() Ranker {
	return defaultRanker
}
