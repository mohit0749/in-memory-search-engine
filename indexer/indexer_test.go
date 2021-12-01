package indexer

import (
	"fmt"
	"testSearchEngine/idgenerator/autoincrementor"
	"testSearchEngine/preprocessor/tokenizer"
	"testSearchEngine/preprocessor/tolowercase"
	"testSearchEngine/ranker/occurrance"
	"testSearchEngine/storage/inmemory"
	"testing"
)

func TestNewIndexer(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		storage := inmemory.NewInMemoryStorage()
		ranker := occurrance.NewOccrranceRanker()
		idgen := autoincrementor.NewAutoIncrementGenerator()
		engine := NewIndexer(storage, ranker, idgen)
		engine.AddPreProcessor(tolowercase.NewLowerCasePreprocessor())
		engine.AddPreProcessor(tokenizer.NewTokenizer(" "))
		/*
			Doc1: apple is a fruit
			Doc2: apple apple come on.
			Doc3: oranges are sour
			Doc4: apple is sweet
		*/
		engine.Insert("Fruit", "apple is a fruit")
		engine.Insert("Fruit", "apple apple come on.")
		engine.Insert("Fruit", "oranges are sour")
		engine.Insert("Fruit", "apple is sweet")
		docs, err := engine.Search("Fruit", "apple")
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(docs)

		docs, err = engine.Search("Fruit", "apple sour")
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(docs)
	})
}
