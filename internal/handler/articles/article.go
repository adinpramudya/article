package articles

import (
	"article/internal/models/articles"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Create Article godoc
// @Summary Create a new article
// @Description Create a new article with title, content, category, status
// @Tags Articles
// @Accept json
// @Produce json
// @Param article body articles.ArticleRequest true "Article Request"
// @Success 201 {object} articles.ArticleResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /articles/ [post]
func (h *handler) CreateArticle(c *gin.Context) {
	var req articles.ArticleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// mapping ke model
	article := articles.Article{
		Title:    req.Title,
		Content:  req.Content,
		Category: req.Category,
		Status:   req.Status,
	}

	if err := h.service.CreateArticle(&article); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, article)
}


// Get All Articles godoc
// @Summary List all articles
// @Description Get paginated list of articles
// @Tags Articles
// @Produce json
// @Param limit query int false "Limit per page"
// @Param offset query int false "Offset"
// @Success 200 {array} articles.ArticleResponse
// @Failure 500 {object} map[string]string
// @Router /articles/ [get]
func (h *handler) GetArticles(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	articlesList, err := h.service.GetAllArticles(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Konversi []Article -> []*Article
	articlesPtr := make([]*articles.Article, len(articlesList))
	for i := range articlesList {
		articlesPtr[i] = &articlesList[i]
	}

	c.JSON(http.StatusOK, articles.NewArticlesResponse(articlesPtr))
}


// Get Article by ID godoc
// @Summary Get article by ID
// @Description Get article by ID
// @Tags Articles
// @Produce json
// @Param id path int true "Article ID"
// @Success 200 {object} articles.ArticleResponse
// @Failure 404 {object} map[string]string
// @Router /articles/{id} [get]
func (h *handler) GetArticleByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	article, err := h.service.GetArticleByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "article not found"})
		return
	}

	c.JSON(http.StatusOK, articles.NewArticleResponse(article))
}

// Update Article godoc
// @Summary Update an article
// @Description Update article fields: title, content, category, status
// @Tags Articles
// @Accept json
// @Produce json
// @Param id path int true "Article ID"
// @Param article body articles.ArticleRequest true "Article Request"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /articles/{id} [put]
func (h *handler) UpdateArticle(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	var article articles.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.UpdateArticle(uint(id), &article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Article updated successfully"})
}

// Delete Article godoc
// @Summary Delete an article
// @Description Delete an article by ID
// @Tags Articles
// @Produce json
// @Param id path int true "Article ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /articles/{id} [delete]
func (h *handler) DeleteArticle(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	if err := h.service.DeleteArticle(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Article deleted successfully"})
}