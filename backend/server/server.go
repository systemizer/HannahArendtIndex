package server

import (
	"time"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"github.com/systemizer/ArendtArchives/backend/store"
	"github.com/systemizer/ArendtArchives/backend/store/sqlite"
)

type Server struct {
	store  store.Store
	router *gin.Engine
	cache  *cache.Cache
}

func New() (*Server, error) {
	se := &Server{}

	// Initialize Store
	s, err := sqlite.New()
	//s, err := postgres.New()
	if err != nil {
		return nil, err
	}
	se.store = s

	// Initialize Cache
	se.cache = cache.New(time.Minute*5, time.Minute*10)

	// Initialize Router
	r := gin.Default()

	// Set to no-cache on index page
	r.Use(func() gin.HandlerFunc {
		return func(c *gin.Context) {
			if c.Request.URL.Path == "/" {
				c.Writer.Header().Set("Cache-Control", "no-cache")
			}
		}
	}())

	r.Use(static.Serve("/", static.LocalFile("/web", true)))

	v1 := r.Group("/v1")
	{
		v1.GET("/health", healthEndpoint)
		v1.POST("/search", se.searchEndpoint)
		v1.POST("/summary", se.summaryEndpoint)

	}
	se.router = r

	return se, nil
}

func (s *Server) Start() {
	s.router.Run()

}

func handleError(c *gin.Context, err error) {
	c.Error(err)
	c.Abort()
}
