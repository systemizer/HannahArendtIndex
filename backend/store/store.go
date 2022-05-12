package store

import (
	"time"

	"gorm.io/gorm"
)

type Collection struct {
	gorm.Model      `json:"-"`
	ID              uint             `gorm:"primarykey" json:"id"`
	Title           string           `json:"title" gorm:"index"`
	LOCID           string           `json:"locId" gorm:"unique"`
	CollectionItems []CollectionItem `json:"-"`
}

type CollectionItem struct {
	gorm.Model      `json:"-"`
	LOCID           string `json:"locId" gorm:"unique"`
	ImageURL        string `json:"imageUrl"`
	CollectionIndex int    `json:"index"`

	CollectionID int
	Collection   Collection

	// Populated Fields
	OCRResult   *string
	CustomTitle *string    `json:"customTitle"`
	CustomDate  *time.Time `json:"customDate"`

	// Full Text Search
	// TSV *string `gorm:"type:tsvector;index:,type:gin"`
	TSV *string
}

type CollectionItemSearchResult struct {
	ID             uint
	LOCID          string `json:"locId"`
	ImageURL       string `json:"imageUrl"`
	OCRResult      string
	Headline       string `json:"headline"`
	CollectionName string `json:"collectionName"`
}

type SearchSummaryResult struct {
	CollectionID   int    `json:"collectionId"`
	CollectionName string `json:"collectionName"`
	Count          int    `json:"count"`
}

type Store interface {
	ListCollections() ([]Collection, error)
	ListCollectionItems() ([]CollectionItem, error)
	AddCollection(*Collection) error
	GetCollectionByLocID(locID string) (Collection, error)
	GetCollectionItemByLocID(locID string) (CollectionItem, error)

	SaveCollectionItem(*CollectionItem) error
	AddCollectionItem(*CollectionItem) error

	SearchCollectionItems(query string, page int, collections []int) ([]CollectionItemSearchResult, error)
	SearchSummary(query string) ([]SearchSummaryResult, error)
}
