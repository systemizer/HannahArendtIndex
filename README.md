## To regenerate sqlite database

```
make dump-data

# Then manually...
# 1. Remove any "SET" queries
# 2. Remove any lines with references to pg_catalog
# 3. Add BEGIN; at the beginning and END; at the end

# Next generate the table (if not already generated) by starting the server
# with the sqlite db configured.

# Then, go into the sqlite db and import the data
sqlite3
> .open gorm.db
> .read dump.sql

# Next, generate the search table
> CREATE VIRTUAL TABLE collection_items_fts USING fts5(id UNINDEXED, ocr_result);
> INSERT INTO collection_items_fts(id, ocr_result) SELECT id, ocr_result from collection_items;
> CREATE INDEX idx_collection_items_collection_id ON collection_items(collection_id);
```