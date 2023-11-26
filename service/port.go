package service

import (
	"context"
)

type BrandRepo interface {
	Add(ctx context.Context, brand *Brand) (*Brand, error)
	GetItemByID(ctx context.Context, brandID string) (*Brand, error)
	GetItems(ctx context.Context, page int64, limit int64) (*BrandResult, error)
	UpdateItemByID(ctx context.Context, brandID string, brand *Brand) error
	DeleteItemByID(ctx context.Context, brandID string) error
}

type CategoryRepo interface {
	Add(ctx context.Context, ctgry *Category) (*Category, error)
	GetItemByID(ctx context.Context, ctgryID string) (*Category, error)
	GetItems(ctx context.Context, page int64, limit int64) (*CategoryResult, error)
	UpdateItemByID(ctx context.Context, ctgryID string, ctgry *Category) error
	DeleteItemByID(ctx context.Context, ctgryID string) error
}

type SupplierRepo interface {
	Add(ctx context.Context, spplr *Supplier) (*Supplier, error)
	GetItemByID(ctx context.Context, spplrID string) (*Supplier, error)
	GetItems(ctx context.Context, page int64, limit int64) (*SupplierResult, error)
	UpdateItemByID(ctx context.Context, spplrID string, spplr *Supplier) error
	DeleteItemByID(ctx context.Context, spplrID string) error
}

type ProductRepo interface {
	Add(ctx context.Context, product *Product) (*Product, error)
	GetItemByID(ctx context.Context, productID string) (*Product, error)
	GetItems(ctx context.Context, filterParams FilterProductsParams) (*ProductResult, error)
	UpdateItemByID(ctx context.Context, productID string, product *Product) error
	DeleteItemByID(ctx context.Context, productID string) error
}

type ProductStockRepo interface {
	Add(ctx context.Context, productStock *ProductStock) (*ProductStock, error)
	GetItemByProductID(ctx context.Context, productID string) (*ProductStock, error)
	UpdateItemByID(ctx context.Context, productStockID string, product *ProductStock) error
	DeleteItemByID(ctx context.Context, productStockID string) error
}

type Service interface {
	Response(ctx context.Context, description string, data interface{}) *ResponseData

	AddBrand(ctx context.Context, brand *Brand) (*Brand, error)
	GetBrand(ctx context.Context, brandID string) (*Brand, error)
	GetBrands(ctx context.Context, page, limit int64) (*BrandResult, error)
	UpdateBrand(ctx context.Context, brandID string, brand *Brand) error
	DeleteBrand(ctx context.Context, brandID string) error

	AddCategory(ctx context.Context, ctgry *Category) (*Category, error)
	GetCategory(ctx context.Context, ctgryID string) (*Category, error)
	GetCategories(ctx context.Context, page, limit int64) (*CategoryResult, error)
	UpdateCategory(ctx context.Context, ctgryID string, ctgry *Category) error
	DeleteCategory(ctx context.Context, ctgryID string) error

	AddSupplier(ctx context.Context, spplr *Supplier) (*Supplier, error)
	GetSupplier(ctx context.Context, spplrID string) (*Supplier, error)
	GetSuppliers(ctx context.Context, page, limit int64) (*SupplierResult, error)
	UpdateSupplier(ctx context.Context, spplrID string, spplr *Supplier) error
	DeleteSupplier(ctx context.Context, spplrID string) error

	AddProduct(ctx context.Context, product *Product) (*Product, error)
	GetProduct(ctx context.Context, productID string) (*Product, error)
	GetProducts(ctx context.Context, filterParams FilterProductsParams) (*ProductResult, error)
	UpdateProduct(ctx context.Context, productID string, product *Product) error
	DeleteProduct(ctx context.Context, productID string) error
}
