package sync

import (
	"fmt"
	"time"

	"github.com/systemizer/ArendtArchives/backend/locapi"
	"github.com/systemizer/ArendtArchives/backend/locapi/api"
	"github.com/systemizer/ArendtArchives/backend/store"
	"github.com/systemizer/ArendtArchives/backend/store/sqlite"
	"gorm.io/gorm"
)

type Sync struct {
	store  store.Store
	client locapi.Client
}

func (s *Sync) Run() error {

	col, err := s.client.ListCollections()
	if err != nil {
		return err
	}

	for _, c := range col {
		time.Sleep(time.Second * 1)
		collection, err := s.store.GetCollectionByLocID(c.ID)
		if err == gorm.ErrRecordNotFound {
			collection = store.Collection{
				Title: c.Title,
				LOCID: c.ID,
			}
			err2 := s.store.AddCollection(&collection)
			if err2 != nil {
				return err
			}
		} else if err != nil {
			return err
		}

		cItems, err := s.client.ListCollectionItems(c)
		if err != nil {
			fmt.Printf("Failed to fetch collection items for %s: %v\n", c.ID, err)
			continue
		}

		for _, ci := range cItems {
			_, err := s.store.GetCollectionItemByLocID(ci.ID)
			if err == gorm.ErrRecordNotFound {
				err2 := s.store.AddCollectionItem(&store.CollectionItem{
					CollectionID:    int(collection.ID),
					Collection:      collection,
					LOCID:           ci.ID,
					CollectionIndex: ci.Index,
					ImageURL:        ci.ImageURLs[len(ci.ImageURLs)-1],
				})
				if err2 != nil {
					fmt.Printf("Failed to add collection item: %v\n", err)
					return err
				}
			} else if err != nil {
				return err
			}
		}

	}

	return nil
}

func New() (*Sync, error) {
	s, err := sqlite.New()
	if err != nil {
		return nil, err
	}

	return &Sync{
		store:  s,
		client: api.NewAPIClient(),
	}, err
}
