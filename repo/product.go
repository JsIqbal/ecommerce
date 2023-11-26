package repo

import (
	"context"
	"fmt"
	"strings"

	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/jsiqbal/ecommerce/logger"
	"github.com/jsiqbal/ecommerce/service"
	"github.com/jsiqbal/ecommerce/util"
	"github.com/lib/pq"
)

// DB models
type Product struct {
	ID             string         `db:"id"`
	Name           string         `db:"name"`
	Description    string         `db:"description"`
	Specifications sql.NullString `db:"specifications"`
	BrandID        string         `db:"brand_id"`
	CategoryID     string         `db:"category_id"`
	SupplierID     string         `db:"supplier_id"`
	UnitPrice      float64        `db:"unit_price"`
	DiscountPrice  float64        `db:"discount_price"`
	Tags           pq.StringArray `db:"tags"`
	StatusID       int            `db:"status_id"`
	CreatedAt      int64          `db:"created_at"`
}

type ProductStock struct {
	ID            string `db:"id"`
	ProductID     string `db:"product_id"`
	StockQuantity int64  `db:"stock_quantity"`
	UpdatedAt     int64  `db:"updated_at"`
}

type ProductRepo interface {
	service.ProductRepo
}

type productRepo struct {
	db *sqlx.DB
}

func NewProductRepo(db *sqlx.DB) ProductRepo {
	return &productRepo{
		db: db,
	}
}

func (r *productRepo) Add(ctx context.Context, product *service.Product) (*service.Product, error) {
	var newProduct Product
	err := r.db.QueryRowContext(ctx,
		`INSERT INTO products (
			name, 
			description, 
			specifications, 
			brand_id, 
			category_id, 
			supplier_id, 
			unit_price, 
			discount_price, 
			tags, 
			status_id, 
			created_at
		) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		RETURNING id, name, description, specifications, brand_id, category_id, supplier_id, unit_price, discount_price, tags, status_id, created_at`,
		product.Name,
		product.Description,
		product.Specifications,
		product.Brand.ID,
		product.Category.ID,
		product.Supplier.ID,
		product.UnitPrice,
		product.DiscountPrice,
		pq.Array(product.Tags),
		product.StatusID,
		product.CreatedAt,
	).Scan(
		&newProduct.ID,
		&newProduct.Name,
		&newProduct.Description,
		&newProduct.Specifications,
		&newProduct.BrandID,
		&newProduct.CategoryID,
		&newProduct.SupplierID,
		&newProduct.UnitPrice,
		&newProduct.DiscountPrice,
		&newProduct.Tags,
		&newProduct.StatusID,
		&newProduct.CreatedAt)
	if err != nil {
		logger.Error(ctx, "can not create product", err)
		return nil, err
	}

	// insert product stock
	_, err = r.db.ExecContext(ctx,
		"INSERT INTO product_stocks (product_id, stock_quantity, updated_at) VALUES ($1, $2, $3)",
		newProduct.ID, product.ProductStock.StockQuantity, util.GetCurrentTimestamp(),
	)
	if err != nil {
		logger.Error(ctx, "can not create product stock", err)
		return nil, err
	}

	createdProduct, err := r.formatProduct(ctx, &newProduct)
	if err != nil {
		logger.Error(ctx, "can not aggregate product info", err)
		return nil, err
	}

	return createdProduct, nil
}

func (r *productRepo) GetItemByID(ctx context.Context, productID string) (*service.Product, error) {
	var dbProduct Product

	err := r.db.Get(&dbProduct, "SELECT * FROM products WHERE id = $1", productID)
	if err == sql.ErrNoRows {
		// No product found
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	formattedProduct, err := r.formatProduct(ctx, &dbProduct)
	if err != nil {
		return nil, err
	}

	return formattedProduct, nil
}

func (r *productRepo) GetItems(ctx context.Context, filterParams service.FilterProductsParams) (*service.ProductResult, error) {
	// check max price
	if filterParams.MaxPrice == 0 {
		filterParams.MaxPrice = service.MAX_INF
	}

	// calculate offset based on page and limit for pagination
	if filterParams.Page == 0 {
		filterParams.Page = 1
	}

	offset := (filterParams.Page - 1) * filterParams.Limit

	// Start building the SQL query
	query := "SELECT * FROM products"

	// Add filters to the query
	query += generateFilterConditions(filterParams)

	// Add pagination
	query += fmt.Sprintf(" ORDER BY unit_price ASC OFFSET %d LIMIT %d", offset, filterParams.Limit)

	logger.Info(ctx, "Full query", query)

	// Fetch products and total count
	var dbProducts []Product
	err := r.db.Select(&dbProducts, query)
	if err != nil {
		return nil, err
	}

	// Fetch total count
	totalCount, err := r.getTotalProductCount(ctx, filterParams)
	if err != nil {
		return nil, err
	}

	var products []service.Product
	for _, dbProduct := range dbProducts {
		products = append(products, service.Product{
			ID:             dbProduct.ID,
			Name:           dbProduct.Name,
			Description:    dbProduct.Description,
			Specifications: dbProduct.Specifications.String,
			UnitPrice:      dbProduct.UnitPrice,
			DiscountPrice:  dbProduct.DiscountPrice,
			Tags:           dbProduct.Tags,
			StatusID:       dbProduct.StatusID,
			CreatedAt:      dbProduct.CreatedAt,
		})
	}

	result := &service.ProductResult{
		Products: products,
		Total:    totalCount,
		Page:     filterParams.Page,
		Limit:    filterParams.Limit,
	}

	return result, nil
}

// func (r *productRepo) UpdateItemByID(ctx context.Context, productID string, product *service.Product) error {
// 	_, err := r.db.Exec(
// 		`UPDATE products
// 		SET
// 			name = $1,
// 			description = $2,
// 			specifications = $3,
// 			brand_id = $4,
// 			category_id = $5,
// 			supplier_id = $6,
// 			unit_price = $7,
// 			discount_price = $8,
// 			tags = $9,
// 			status_id = $10,
// 		WHERE id = $11`,
// 		product.Name,
// 		product.Description,
// 		product.Specifications,
// 		product.Brand.ID,
// 		product.Category.ID,
// 		product.Supplier.ID,
// 		product.UnitPrice,
// 		product.DiscountPrice,
// 		product.Tags,
// 		product.StatusID,
// 	)
// 	if err != nil {
// 		return err
// 	}

//		return nil
//	}
func (r *productRepo) UpdateItemByID(ctx context.Context, productID string, product *service.Product) error {
	// Convert the tags slice to a comma-separated string
	tagsString := strings.Join(product.Tags, ",")

	_, err := r.db.Exec(
		`UPDATE products 
        SET 
            name = $1, 
            description = $2, 
            specifications = $3, 
            brand_id = $4, 
            category_id = $5, 
            supplier_id = $6, 
            unit_price = $7, 
            discount_price = $8, 
            tags = $9, 
            status_id = $10 
        WHERE id = $11`,
		product.Name,
		product.Description,
		product.Specifications,
		product.Brand.ID,
		product.Category.ID,
		product.Supplier.ID,
		product.UnitPrice,
		product.DiscountPrice,
		tagsString, // Use the concatenated string instead of the slice
		product.StatusID,
		productID, // Assuming productID is the last parameter in the WHERE clause
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *productRepo) DeleteItemByID(ctx context.Context, productId string) error {
	_, err := r.db.Exec("DELETE FROM product_stocks WHERE product_id = $1", productId)
	if err != nil {
		return err
	}

	_, err = r.db.Exec("DELETE FROM products WHERE id = $1", productId)
	if err != nil {
		return err
	}

	return nil
}

func (r *productRepo) GetProductStock(ctx context.Context, productID string) (*service.ProductStock, error) {
	var productStock ProductStock

	err := r.db.Get(&productStock, "SELECT * FROM product_stocks WHERE product_id = $1", productID)
	if err == sql.ErrNoRows {
		// No product found
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &service.ProductStock{
		ID:            productStock.ProductID,
		ProductID:     productStock.ProductID,
		StockQuantity: productStock.StockQuantity,
		UpdatedAt:     productStock.UpdatedAt,
	}, nil
}

func (r *productRepo) formatProduct(ctx context.Context, dbProduct *Product) (*service.Product, error) {
	logger.Info(ctx, "db product format", dbProduct)
	createdProduct := &service.Product{
		ID:             dbProduct.ID,
		Name:           dbProduct.Name,
		Description:    dbProduct.Description,
		Specifications: dbProduct.Specifications.String,
		UnitPrice:      dbProduct.UnitPrice,
		DiscountPrice:  dbProduct.DiscountPrice,
		Tags:           dbProduct.Tags,
		StatusID:       dbProduct.StatusID,
		CreatedAt:      dbProduct.CreatedAt,
	}

	// fetch brand, then aggregate with product
	var brand Brand
	err := r.db.Get(&brand, "SELECT * FROM brands WHERE id = $1", dbProduct.BrandID)
	if err != nil {
		logger.Error(ctx, "can not get brand", err)
		return nil, err
	}

	createdProduct.Brand = service.Brand{
		ID:        brand.ID,
		Name:      brand.Name,
		StatusID:  brand.StatusID,
		CreatedAt: brand.CreatedAt,
	}

	// fetch category, then aggregate with product
	var ctgry Category
	err = r.db.Get(&ctgry, "SELECT * FROM categories WHERE id = $1", dbProduct.CategoryID)
	if err != nil {
		logger.Error(ctx, "can not get category", err)
		return nil, err
	}

	createdProduct.Category = service.Category{
		ID:        ctgry.ID,
		Name:      ctgry.Name,
		ParentID:  ctgry.ParentID.String,
		Sequence:  ctgry.Sequence.String,
		StatusID:  ctgry.StatusID,
		CreatedAt: ctgry.CreatedAt,
	}

	// fetch category, then aggregate with product
	var spplr Supplier
	err = r.db.Get(&spplr, "SELECT * FROM suppliers WHERE id = $1", dbProduct.SupplierID)
	if err != nil {
		logger.Error(ctx, "can not get supplier", err)
		return nil, err
	}

	createdProduct.Supplier = service.Supplier{
		ID:                 spplr.ID,
		Name:               spplr.Name,
		Email:              spplr.Email,
		Phone:              spplr.Phone,
		IsVerifiedSupplier: spplr.IsVerifiedSupplier,
		StatusID:           spplr.StatusID,
		CreatedAt:          spplr.CreatedAt,
	}

	// fetch product stock, then aggregate with product
	productStock, err := r.GetProductStock(ctx, dbProduct.ID)
	if err != nil {
		logger.Error(ctx, "can not get product stock", err)
		return nil, err
	}

	createdProduct.ProductStock = *productStock

	return createdProduct, nil
}

func (r *productRepo) getTotalProductCount(ctx context.Context, filterParams service.FilterProductsParams) (int64, error) {
	var totalCount int64

	// Start building the SQL query
	query := "SELECT COUNT(*) FROM products"

	// Add filters to the query

	query += generateFilterConditions(filterParams)

	// Execute the query to get total count
	err := r.db.GetContext(ctx, &totalCount, query)
	if err != nil {
		return 0, err
	}

	return totalCount, nil
}

// generateFilterConditions generates the SQL conditions based on filter parameters
func generateFilterConditions(filterParams service.FilterProductsParams) string {
	// Implement your logic to generate conditions based on filter parameters
	// For example, you can check if filterParams contains non-zero values and construct conditions accordingly.
	// Return the generated conditions as a string.
	whereClause := fmt.Sprintf(" WHERE status_id = 1 AND unit_price >= %f AND unit_price <= %f", filterParams.MinPrice, filterParams.MaxPrice)

	if filterParams.Name != "" {
		whereClause += fmt.Sprintf(" AND name = '%s'", filterParams.Name)
	}

	if len(filterParams.BrandIDs) > 0 {
		brandConditions := make([]string, len(filterParams.BrandIDs))
		for i, brandID := range filterParams.BrandIDs {
			brandConditions[i] = fmt.Sprintf("brand_id = '%s'", brandID)
		}

		whereClause += fmt.Sprintf(" AND (%s)", strings.Join(brandConditions, " OR "))
	}

	if len(filterParams.CategoryID) > 0 {
		whereClause += fmt.Sprintf(" AND category_id = '%s'", filterParams.CategoryID)
	}

	if len(filterParams.SupplierID) > 0 {
		whereClause += fmt.Sprintf(" AND supplier_id = '%s'", filterParams.SupplierID)
	}

	if filterParams.IsVerifiedSupplier {
		whereClause += fmt.Sprintf(" AND is_verified_supplier = '%v'", filterParams.IsVerifiedSupplier)
	}

	return whereClause
}
