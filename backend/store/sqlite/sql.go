package sqlite

const sqlSearchCollectionSummary = `
SELECT
  collections.title as collection_name, collections.id as collection_id,  count(*) as count
FROM
  collection_items_fts(?) INNER JOIN collection_items ON collection_items_fts.id = collection_items.id
  INNER JOIN collections ON collections.id = collection_items.collection_id
GROUP BY collection_name
ORDER BY count DESC;
`

const sqlSearchCollectionSummaryNoQuery = `
SELECT
  collections.title as collection_name, collections.id as collection_id,  count(*) as count
FROM
  collection_items INNER JOIN collections ON collections.id = collection_items.collection_id
GROUP BY collection_name
ORDER BY count DESC;
`

const sqlSearchWithCollectionsWithQuery = `
SELECT 
  collection_items.id, image_url, collection_items.loc_id, collection_items.ocr_result, snippet(collection_items_fts, 1, '', '', '', 40) as headline, collections.title as collection_name 
FROM 
  collection_items_fts(?) INNER JOIN collection_items ON collection_items_fts.id = collection_items.id 
  INNER JOIN collections ON collections.id = collection_items.collection_id
WHERE collections.id IN ?
ORDER BY rank LIMIT 15 OFFSET ?;
`

const sqlSearchWithCollectionsNoQuery = `
SELECT 
  collection_items.id, image_url, collection_items.loc_id, collection_items.ocr_result, substr(ocr_result, 1, 256) as headline, collections.title as collection_name 
FROM 
  collection_items INNER JOIN collections ON collection_items.collection_id = collections.id 
WHERE collections.id IN ?
LIMIT 15 OFFSET ?;
`

const sqlSearchNoCollectionsWithQuery = `
SELECT 
  collection_items.id, image_url, collection_items.loc_id, collection_items.ocr_result, snippet(collection_items_fts, 1, '', '', '', 64) as headline, collections.title as collection_name 
FROM 
  collection_items_fts(?) INNER JOIN collection_items ON collection_items_fts.id = collection_items.id 
  INNER JOIN collections ON collections.id = collection_items.collection_id
ORDER BY rank LIMIT 15 OFFSET ?;
`

const sqlSearchNoCollectionsNoQuery = `
SELECT 
  collection_items.id, image_url, collection_items.loc_id, collection_items.ocr_result, substr(ocr_result, 1, 256) as headline, collections.title as collection_name 
FROM 
  collection_items INNER JOIN collections ON collection_items.collection_id = collections.id 
LIMIT 15 OFFSET ?;
`
