package articles

import "time"

// ArticleResponse adalah struktur data untuk response API
type ArticleResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Category  string    `json:"category"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewArticleResponse mengonversi Article menjadi ArticleResponse
func NewArticleResponse(a *Article) *ArticleResponse {
	return &ArticleResponse{
		ID:        a.ID,
		Title:     a.Title,
		Content:   a.Content,
		Category:  a.Category,
		Status:    a.Status,
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
	}
}

// NewArticlesResponse mengonversi slice Article menjadi slice ArticleResponse
func NewArticlesResponse(articles []*Article) []*ArticleResponse {
	res := make([]*ArticleResponse, len(articles))
	for i, a := range articles {
		res[i] = NewArticleResponse(a)
	}
	return res
}
