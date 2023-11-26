package service

type Product struct {
	ID             string       `json:"id"`
	Name           string       `json:"name"`
	Description    string       `json:"description"`
	Specifications string       `json:"specifications"`
	Brand          Brand        `json:"brand"`
	Category       Category     `json:"category"`
	Supplier       Supplier     `json:"supplier"`
	UnitPrice      float64      `json:"unit_price"`
	DiscountPrice  float64      `json:"discount_price"`
	Tags           []string     `json:"tags"`
	StatusID       int          `json:"status_id"`
	CreatedAt      int64        `json:"created_at"`
	ProductStock   ProductStock `json:"product_stock"`
}

type ProductStock struct {
	ID            string `json:"id,omitempty"`
	ProductID     string `json:"product_id,omitempty"`
	StockQuantity int64  `json:"stock_quantity"`
	UpdatedAt     int64  `json:"updated_at"`
}

type FilterProductsParams struct {
	Name               string   `json:"name"`
	MaxPrice           float64  `json:"max_price"`
	MinPrice           float64  `json:"min_price"`
	BrandIDs           []string `json:"brand_ids"`
	CategoryID         string   `json:"category_id"`
	SupplierID         string   `json:"supplier_id"`
	IsVerifiedSupplier bool     `json:"is_verified_supplier"`
	Page               int64    `json:"page"`
	Limit              int64    `json:"limit"`
}

type ProductResult struct {
	Products []Product `json:"products"`
	Total    int64     `json:"total"`
	Page     int64     `json:"page"`
	Limit    int64     `json:"limit"`
}
