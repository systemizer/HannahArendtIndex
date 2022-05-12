package postgres

import (
	"errors"
	"os"

	"github.com/lib/pq"
	"github.com/systemizer/ArendtArchives/backend/store"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresStore struct {
	db *gorm.DB
}

func New() (store.Store, error) {
	dsn := "host=localhost user=postgres password=postgres dbname=arendt port=5432 sslmode=disable"
	if os.Getenv("DATABASE_URL") != "" {
		dsn = os.Getenv("DATABASE_URL")
	}

	/* 	newLogger := logger.New(
	   		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	   		logger.Config{
	   			SlowThreshold:             time.Second, // Slow SQL threshold
	   			LogLevel:                  logger.Info, // Log level
	   			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
	   			Colorful:                  false,       // Disable color
	   		},
	   	)
	*/
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&store.Collection{}, &store.CollectionItem{})
	if err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) ListCollections() ([]store.Collection, error) {
	var collections []store.Collection
	res := s.db.Order("id").Limit(534).Find(&collections)
	return collections, res.Error
}

func (s *PostgresStore) AddCollection(c *store.Collection) error {
	res := s.db.Create(c)
	return res.Error
}

func (s *PostgresStore) AddCollectionItem(c *store.CollectionItem) error {
	res := s.db.Create(c)
	return res.Error
}

func (s *PostgresStore) ListCollectionItems() ([]store.CollectionItem, error) {
	var collectionItems []store.CollectionItem
	res := s.db.Preload("Collection").Where("id >= ?", 69041).Find(&collectionItems)
	return collectionItems, res.Error
}

func (s *PostgresStore) SaveCollectionItem(c *store.CollectionItem) error {
	tx := s.db.Save(c)
	return tx.Error
}

func (s *PostgresStore) GetCollectionByLocID(locID string) (store.Collection, error) {
	collection := store.Collection{}
	res := s.db.Where("loc_id = ?", locID).First(&collection)
	return collection, res.Error
}

func (s *PostgresStore) GetCollectionItemByLocID(locID string) (store.CollectionItem, error) {
	collectionItem := store.CollectionItem{}
	res := s.db.Where("loc_id = ?", locID).First(&collectionItem)
	return collectionItem, res.Error
}

func (s *PostgresStore) SearchCollectionItems(query string, page int, collections []int) ([]store.CollectionItemSearchResult, error) {
	result := []store.CollectionItemSearchResult{}
	var res *gorm.DB

	// If no query is given,  then simply list collection items
	if query == "" {
		res = s.db.Raw(sqlSearchNoQuery, pq.Array(collections), pq.Array(collections), page*15).Scan(&result)
	} else {
		res = s.db.Raw(sqlSearch, query, pq.Array(collections), pq.Array(collections), page*15).Scan(&result)
	}

	return result, res.Error
}

func (s *PostgresStore) SearchSummary(query string) ([]store.SearchSummaryResult, error) {
	return nil, errors.New("not yet implemented")
}
