package staticstore

import (
	"embed"
	"encoding/json"

	"github.com/systemizer/ArendtArchives/backend/store"
)

//go:embed hannah-arendt-papers.json
var f embed.FS

type dataFile struct {
	Content struct {
		Results []struct {
			ID    string `json:"id"`
			Title string `json:"title"`
		} `json:"results"`
	} `json:"content"`
}

type StaticStore struct {
	collections []store.Collection
}

func New() (*StaticStore, error) {

	data, err := f.ReadFile("hannah-arendt-papers.json")
	if err != nil {
		return nil, err
	}

	var df dataFile
	err = json.Unmarshal(data, &df)
	if err != nil {
		return nil, err
	}

	cols := []store.Collection{}
	for _, d := range df.Content.Results {
		cols = append(cols, store.Collection{
			Title: d.Title,
			LOCID: d.ID,
		})
	}

	return &StaticStore{
		collections: cols,
	}, nil
}

func (s *StaticStore) ListCollections() ([]store.Collection, error) {
	return s.collections, nil
}

func (s *StaticStore) AddCollection(c *store.Collection) error {
	return nil
}

func (s *StaticStore) AddCollectionItem(c *store.CollectionItem) error {
	return nil
}

func (s *StaticStore) ListCollectionItems() ([]store.CollectionItem, error) {
	return []store.CollectionItem{}, nil
}

func (s *StaticStore) SaveCollectionItem(c *store.CollectionItem) error {
	return nil
}
