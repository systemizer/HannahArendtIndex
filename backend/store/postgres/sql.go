package postgres

const sqlSearch = `
SELECT 
  collection_items.id, image_url, collection_items.loc_id, ocr_result, ts_headline(ocr_result, q, 'StartSel = "", StopSel = "", MaxWords=200, MinWords=100') as headline, ts_rank(tsv, q) as rank, collections.title as collection_name
FROM 
  collection_items LEFT JOIN collections ON (collections.id = collection_items.collection_id), plainto_tsquery(?) q 
WHERE tsv @@ q AND (collections.id = ANY(?::integer[]) OR cardinality(?::integer[]) = 0)
ORDER BY rank DESC OFFSET ? LIMIT 15
`

const sqlSearchNoQuery = `
SELECT 
  collection_items.id, image_url, collection_items.loc_id, ocr_result, substring(ocr_result from 0 for 500) as headline, collections.title as collection_name
FROM 
  collection_items LEFT JOIN collections ON (collections.id = collection_items.collection_id) 
WHERE  (collections.id = ANY(?::integer[]) OR cardinality(?::integer[]) = 0) AND tsv IS NOT NULL
ORDER BY id OFFSET ? LIMIT 15
`
