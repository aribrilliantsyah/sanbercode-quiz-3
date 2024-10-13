package schemas

type CreateBook struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description,omitempty"`
	ImageUrl    string `json:"image_url,omitempty"`
	ReleaseYear int32  `json:"release_year,omitempty"`
	Price       int32  `json:"price,omitempty"`
	TotalPage   int32  `json:"total_page,omitempty"`
	CategoryID  int64  `json:"category_id" binding:"required"`
}

type UpdateBook struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description,omitempty"`
	ImageUrl    string `json:"image_url,omitempty"`
	ReleaseYear int32  `json:"release_year,omitempty"`
	Price       int32  `json:"price,omitempty"`
	TotalPage   int32  `json:"total_page,omitempty"`
	CategoryID  int64  `json:"category_id" binding:"required"`
}
