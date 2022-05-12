package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"github.com/systemizer/ArendtArchives/backend/store"
)

type SummaryParams struct {
	Search string `json:"search"`
}

type SummaryResponse struct {
	Summary []store.SearchSummaryResult `json:"summary"`
}

func (s *Server) summaryEndpoint(c *gin.Context) {
	var resp SummaryResponse
	var params SummaryParams
	var summary []store.SearchSummaryResult

	err := c.BindJSON(&params)
	if err != nil {
		handleError(c, err)
		return
	}

	cacheKey := params.Search
	cached, found := s.cache.Get(cacheKey)
	if found {
		fmt.Println("using cache!")
		summary = cached.([]store.SearchSummaryResult)
	} else {
		fmt.Println("cache missed")
		summary, err = s.store.SearchSummary(params.Search)
		if err != nil {
			c.Error(err)
			c.Abort()
		}

		s.cache.Add(cacheKey, summary, cache.DefaultExpiration)
	}

	resp.Summary = summary

	c.JSON(200, resp)
}
