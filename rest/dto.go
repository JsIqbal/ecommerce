package rest

/////////////////////// brand dtos //////////////////////

type createBrandReq struct {
	Name     string `json:"name" binding:"required,min=2,max=50"`
	StatusID int    `json:"status_id" binding:"required,validStatusID"`
}

type getBrandReq struct {
	ID string `uri:"id" binding:"required"`
}

type getBrandsReq struct {
	Page  int64 `form:"page" binding:"required,min=1"`
	Limit int64 `form:"limit" binding:"required,min=1,max=100"`
}

type updateBrandReq struct {
	Name     string `json:"name" binding:"required,min=2,max=50"`
	StatusID int    `json:"status_id" binding:"required"`
}

type deleteBrandReq struct {
	ID string `uri:"id" binding:"required"`
}

/////////////////////// category dtos //////////////////////

type createCategoryReq struct {
	Name     string `json:"name" binding:"required,min=2,max=50"`
	ParentID string `json:"parent_id"`
	StatusID int    `json:"status_id" binding:"required,validStatusID"`
}

type getCategoryReq struct {
	ID string `uri:"id" binding:"required"`
}

type getCategoriesReq struct {
	Page  int64 `form:"page" binding:"required,min=1"`
	Limit int64 `form:"limit" binding:"required,min=1,max=100"`
}

type updateCategoryReq struct {
	Name     string `json:"name" binding:"required,min=2,max=50"`
	StatusID int    `json:"status_id"`
}

type deleteCategoryReq struct {
	ID string `uri:"id" binding:"required"`
}

//////////////////////// supplier dtos /////////////////////////

type createSupplierReq struct {
	Name               string `json:"name" binding:"required,min=2,max=50"`
	Email              string `json:"email" binding:"required,email"`
	Phone              string `json:"phone" binding:"required,validPhone"`
	StatusID           int    `json:"status_id" binding:"required,validStatusID"`
	IsVerifiedSupplier bool   `json:"is_verified_supplier" binding:"required"`
}

type getSupplierReq struct {
	ID string `uri:"id" binding:"required"`
}

type getSuppliersReq struct {
	Page  int64 `form:"page" binding:"required,min=1"`
	Limit int64 `form:"limit" binding:"required,min=1,max=100"`
}

type updateSupplierReq struct {
	Name               string `json:"name" binding:"required,min=2,max=50"`
	Email              string `json:"email" binding:"required,email"`
	Phone              string `json:"phone" binding:"required,validPhone"`
	StatusID           int    `json:"status_id" binding:"required,validStatusID"`
	IsVerifiedSupplier bool   `json:"is_verified_supplier" binding:"required"`
}

type deleteSupplierReq struct {
	ID string `uri:"id" binding:"required"`
}

//////////////////////////////// product dtos //////////////////////////////////

type createProductReq struct {
	Name           string   `json:"name" binding:"required,min=2,max=50"`
	Description    string   `json:"description" binding:"required,min=2,max=500"`
	Specifications string   `json:"specifications" binding:"min=0,max=500"`
	BrandID        string   `json:"brand_id" binding:"required"`
	CategoryID     string   `json:"category_id" binding:"required"`
	SupplierID     string   `json:"supplier_id" binding:"required"`
	UnitPrice      float64  `json:"unit_price" binding:"required,min=0"`
	DiscountPrice  float64  `json:"discount_price" binding:"required,min=0"`
	Tags           []string `json:"tags" binding:"required"`
	StatusID       int      `json:"status_id" binding:"required,validStatusID"`
	StockQuantity  int64    `json:"stock_quantity" binding:"required,min=1"`
}

type getProductReq struct {
	ID string `uri:"id" binding:"required"`
}

type getProductsReq struct {
	Name       string   `form:"name"`
	MinPrice   float64  `form:"min_price" binding:"min=0"`
	MaxPrice   float64  `form:"max_price" binding:"min=0"`
	BrandIDs   []string `form:"brand_ids"`
	CategoryID string   `form:"category_id"`
	SupplierID string   `form:"supplier_id"`
	Page       int64    `form:"Page"`
	Limit      int64    `form:"limit" binding:"required,min=1,max=100"`
}

type updateProductReq struct {
	Name           string   `json:"name" binding:"required,min=2,max=50"`
	Description    string   `json:"description" binding:"required,min=2,max=500"`
	Specifications string   `json:"specifications" binding:"min=0,max=500"`
	BrandID        string   `json:"brand_id" binding:"required"`
	CategoryID     string   `json:"category_id" binding:"required"`
	SupplierID     string   `json:"supplier_id" binding:"required"`
	UnitPrice      float64  `json:"unit_price" binding:"required,min=0"`
	DiscountPrice  float64  `json:"discount_price" binding:"required,min=0"`
	Tags           []string `json:"tags" binding:"required"`
	StatusID       int      `json:"status_id" binding:"required,validStatusID"`
	StockQuantity  int64    `json:"stock_quantity" binding:"required,min=1"`
}

type deleteProductReq struct {
	ID string `uri:"id" binding:"required"`
}
