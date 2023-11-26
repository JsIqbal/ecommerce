package service

type Brand struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	StatusID  int    `json:"status_id"`
	CreatedAt int64  `json:"created_at"`
}

type BrandResult struct {
	Brands []Brand `json:"brands"`
	Total  int64   `json:"total"`
	Page   int64   `json:"page"`
	Limit  int64   `json:"limit"`
}
