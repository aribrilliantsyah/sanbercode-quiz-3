package schemas

import "time"

type CategoryData struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	CreatedBy  string    `json:"created_by"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedBy string    `json:"modified_by"`
	ModifiedAt time.Time `json:"modified_at"`
}

type UserData struct {
	ID         int64     `json:"id"`
	Username   string    `json:"username"`
	CreatedBy  string    `json:"created_by"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedBy string    `json:"modified_by"`
	ModifiedAt time.Time `json:"modified_at"`
}

type BookData struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ImageUrl    string    `json:"image_url"`
	ReleaseYear int32     `json:"release_year"`
	Price       int32     `json:"price"`
	TotalPage   int32     `json:"total_page"`
	Thickness   string    `json:"thickness"`
	CategoryID  int64     `json:"category_id"`
	CreatedBy   string    `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
	ModifiedBy  string    `json:"modified_by"`
	ModifiedAt  time.Time `json:"modified_at"`
}
