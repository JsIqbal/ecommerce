package service

type Supplier struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	Email              string `json:"email"`
	Phone              string `json:"phone"`
	StatusID           int    `json:"status_id"`
	IsVerifiedSupplier bool   `json:"is_verified_supplier"`
	CreatedAt          int64  `json:"created_at"`
}

type SupplierResult struct {
	Suppliers []Supplier `json:"suppliers"`
	Total     int64      `json:"total"`
	Page      int64      `json:"page"`
	Limit     int64      `json:"limit"`
}
