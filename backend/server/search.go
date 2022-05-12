package server

import (
	"github.com/gin-gonic/gin"
	"github.com/systemizer/ArendtArchives/backend/store"
)

type SearchParams struct {
	Search      string `json:"search"`
	Collections []int  `json:"collections" binding:"required"`
	Page        int    `json:"page"`
}

type SearchResponse struct {
	Results []store.CollectionItemSearchResult `json:"results"`
}

func (s *Server) searchEndpoint(c *gin.Context) {
	var params SearchParams
	var resp SearchResponse
	err := c.BindJSON(&params)
	if err != nil {
		handleError(c, err)
		return
	}

	cols, err := s.store.SearchCollectionItems(params.Search, params.Page, params.Collections)
	if err != nil {
		handleError(c, err)
		return
	}
	resp.Results = cols

	c.JSON(200, resp)
}
