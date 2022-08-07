package dto

type CreateArticleDTO struct {
	Author string `json:"author" form:"author" binding:"required,min=1"`
	Title  string `json:"title" form:"title" binding:"required,min=1"`
	Body   string `json:"body" form:"body" binding:"required,min=1"`
}

type GetArticleDTO struct {
	Author string `json:"author" form:"author"`
	Search string `json:"search" form:"search"`
}
