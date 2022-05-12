package sqlite

import (
	"log"
	"os"
	"time"

	"github.com/systemizer/ArendtArchives/backend/store"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type SqliteStore struct {
	db *gorm.DB
}

func New() (store.Store, error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)

	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&store.Collection{}, &store.CollectionItem{})
	if err != nil {
		return nil, err
	}

	return &SqliteStore{
		db: db,
	}, nil
}

func (s *SqliteStore) ListCollections() ([]store.Collection, error) {
	var collections []store.Collection
	res := s.db.Order("id").Limit(534).Find(&collections)
	return collections, res.Error
}

func (s *SqliteStore) AddCollection(c *store.Collection) error {
	res := s.db.Create(c)
	return res.Error
}

func (s *SqliteStore) AddCollectionItem(c *store.CollectionItem) error {
	res := s.db.Create(c)
	return res.Error
}

func (s *SqliteStore) ListCollectionItems() ([]store.CollectionItem, error) {
	var collectionItems []store.CollectionItem
	res := s.db.Preload("Collection").Find(&collectionItems)
	return collectionItems, res.Error
}

func (s *SqliteStore) SaveCollectionItem(c *store.CollectionItem) error {
	tx := s.db.Save(c)
	return tx.Error
}

func (s *SqliteStore) GetCollectionByLocID(locID string) (store.Collection, error) {
	collection := store.Collection{}
	res := s.db.Where("loc_id = ?", locID).First(&collection)
	return collection, res.Error
}

func (s *SqliteStore) GetCollectionItemByLocID(locID string) (store.CollectionItem, error) {
	collectionItem := store.CollectionItem{}
	res := s.db.Where("loc_id = ?", locID).First(&collectionItem)
	return collectionItem, res.Error
}

func (s *SqliteStore) SearchCollectionItems(query string, page int, collections []int) ([]store.CollectionItemSearchResult, error) {
	result := []store.CollectionItemSearchResult{}
	var res *gorm.DB
	if len(collections) == 0 {
		if query == "" {
			res = s.db.Raw(sqlSearchNoCollectionsNoQuery, page*15).Scan(&result)
		} else {
			res = s.db.Raw(sqlSearchNoCollectionsWithQuery, query, page*15).Scan(&result)
		}
	} else {
		if query == "" {
			res = s.db.Raw(sqlSearchWithCollectionsNoQuery, collections, page*15).Scan(&result)
		} else {
			res = s.db.Raw(sqlSearchWithCollectionsWithQuery, query, collections, page*15).Scan(&result)
		}
	}

	return result, res.Error
}

func (s *SqliteStore) SearchSummary(query string) ([]store.SearchSummaryResult, error) {
	var res *gorm.DB
	result := []store.SearchSummaryResult{}

	if query == "" {
		res = s.db.Raw(sqlSearchCollectionSummaryNoQuery).Scan(&result)
	} else {
		res = s.db.Raw(sqlSearchCollectionSummary, query).Scan(&result)
	}

	return result, res.Error
}
