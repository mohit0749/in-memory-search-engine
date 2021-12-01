package occurrance

import (
	"sort"
	"testSearchEngine/ranker"
	"testSearchEngine/storage"
)

type occurrance struct {
	countMap map[string]int
}

type documentRank struct {
	rank  int
	docId string
}

func (d documentRank) GetDocId() string {
	return d.docId
}

func (d documentRank) GetRank() int {
	return d.rank
}

func NewOccrranceRanker() *occurrance {
	return &occurrance{countMap: make(map[string]int)}
}

func (o occurrance) RankDocument(docIds []storage.Attributes) ([]ranker.Rank, error) {
	for _, doc := range docIds {
		o.countMap[doc.GetID()]++
	}
	docList := make([]struct {
		key   string
		value int
	}, 0)
	for k, v := range o.countMap {
		docList = append(docList, struct {
			key   string
			value int
		}{key: k, value: v})
	}
	sort.Slice(docList, func(i, j int) bool {
		return docList[i].value > docList[j].value
	})
	rankedDocs := make([]ranker.Rank, 0)
	for i, v := range docList {
		rankedDocs = append(rankedDocs, documentRank{
			rank:  i,
			docId: v.key,
		})
	}
	return rankedDocs, nil
}
