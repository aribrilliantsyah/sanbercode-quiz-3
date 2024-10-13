package schemas

type CreateCategory struct {
	Name string `json:"name" binding:"required"`
	// CreatedBy string `json:"created_by" binding:"required"`
}

type UpdateCategory struct {
	Name string `json:"name" binding:"required"`
	// ModifiedBy string `json:"modified_by" binding:"required"`
}
