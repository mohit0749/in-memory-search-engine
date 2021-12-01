package storage

import (
	"sync"
	"time"
)

var defaultStorage Storage
var once sync.Once

type Attributes interface {
	GetID() string
	GetCreatedTime() *time.Time
}

type Storage interface {
	Add(datasetName string, doc []string, docText, id string) error
	Search(datasetName string, words []string) ([]Attributes, error)
}

func Init(s Storage) {
	once.Do(func() {
		defaultStorage = s
	})
}

func Get() Storage {
	return defaultStorage
}
