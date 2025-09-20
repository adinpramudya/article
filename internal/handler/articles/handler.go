package articles

import (
	"article/internal/models/articles"

	"github.com/gin-gonic/gin"
)

type service interface{
	CreateArticle(article *articles.Article) error
	GetAllArticles(limit, offset int) ([] articles.Article,error)
	GetArticleByID(id uint)(*articles.Article, error)
	UpdateArticle(id uint, updated *articles.Article) error
	DeleteArticle(id uint) error
}

type handler struct{
	*gin.Engine
	service service
}

func NewHandler(api *gin.Engine, service service) *handler{
	return &handler{
		api,
		service,
	}
}
func (h *handler) RegisterRoute() {
	// Group utama: /api/v1
	api := h.Group("/api/v1")

	// Subgroup: /api/v1/articles
	route := api.Group("/articles")

	route.POST("/", h.CreateArticle)        // Create
	route.GET("/", h.GetArticles)           // Read all (with paging: ?limit=10&offset=0)
	route.GET("/:id", h.GetArticleByID)     // Read by ID
	route.PUT("/:id", h.UpdateArticle)      // Full update
	route.PATCH("/:id", h.UpdateArticle)    // Partial update
	route.DELETE("/:id", h.DeleteArticle)   // Delete
}
