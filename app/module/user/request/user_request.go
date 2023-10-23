package request

// Requests & responses for ArticleController & ArticleService
type UserRequest struct {
	Name string `json:"name" form:"name" validate:"required,max=255"`
	Icon string `json:"Icon" form:"Icon" validate:"required"`
}
