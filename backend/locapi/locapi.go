package locapi

type Collection struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Resources []struct {
		Search string `json:"search"`
	} `json:"resources"`
}

type CollectionItem struct {
	ID        string   `json:"id"`
	Index     int      `json:"index"`
	ImageURLs []string `json:"image_url"`
}

type APICollectionResponse struct {
	Content struct {
		Pagination struct {
			Next string `json:"next,omitempty"`
		} `json:"pagination"`
		Results []Collection `json:"results"`
	} `json:"content"`
}

type APICollectionItemResponse struct {
	Content struct {
		Pagination struct {
			Next string `json:"next,omitempty"`
		} `json:"pagination"`
		Results []CollectionItem `json:"results"`
	} `json:"content"`
}

type Client interface {
	ListCollections() ([]Collection, error)
	ListCollectionItems(c Collection) ([]CollectionItem, error)
}
