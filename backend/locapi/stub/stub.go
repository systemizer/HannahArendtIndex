package stub

import (
	"embed"
	"encoding/json"
	"fmt"

	"github.com/systemizer/ArendtArchives/backend/locapi"
)

type StubClient struct {
	collections     []locapi.Collection
	collectionItems []locapi.CollectionItem
}

//go:embed hannah-arendt-papers.json
var f embed.FS

//go:embed collection-item.json
var f2 embed.FS

func NewStubClient() (locapi.Client, error) {
	data, err := f.ReadFile("hannah-arendt-papers.json")
	if err != nil {
		return nil, err
	}

	var res locapi.APICollectionResponse
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	data2, err := f2.ReadFile("collection-item.json")
	if err != nil {
		return nil, err
	}

	var res2 locapi.APICollectionItemResponse
	err = json.Unmarshal(data2, &res2)
	if err != nil {
		return nil, err
	}

	if res2.Content.Pagination.Next == "" {
		fmt.Println("It's empty string")
	}

	return &StubClient{
		collections:     res.Content.Results,
		collectionItems: res2.Content.Results,
	}, nil
}

func (s *StubClient) ListCollections() ([]locapi.Collection, error) {
	return s.collections, nil
}

func (s *StubClient) ListCollectionItems(_ locapi.Collection) ([]locapi.CollectionItem, error) {
	return s.collectionItems, nil
}
