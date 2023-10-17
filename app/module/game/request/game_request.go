package request

// Requests & responses for ArticleController & ArticleService
type GameRequest struct {
	Title   string `json:"title" form:"title" validate:"required,max=255"`
	Content string `json:"content" form:"content" validate:"required"`
}
