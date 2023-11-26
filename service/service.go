package service

import (
	"context"

	"github.com/jsiqbal/ecommerce/util"
)

const (
	MAX_INF = 1000000000000000
)

type service struct {
	brandRepo        BrandRepo
	ctgryRepo        CategoryRepo
	spplrRepo        SupplierRepo
	productRepo      ProductRepo
	productStockRepo ProductStockRepo
}

func NewService(
	brandRepo BrandRepo,
	ctgryRepo CategoryRepo,
	spplrRepo SupplierRepo,
	productRepo ProductRepo,
) Service {
	return &service{
		brandRepo:   brandRepo,
		ctgryRepo:   ctgryRepo,
		spplrRepo:   spplrRepo,
		productRepo: productRepo,
	}
}

func (s *service) Response(ctx context.Context, description string, data interface{}) *ResponseData {
	return &ResponseData{
		Timestamp:   util.GetCurrentTimestamp(),
		Description: description,
		Data:        data,
	}
}

//----------------BRAND----------------

func (s *service) AddBrand(ctx context.Context, brand *Brand) (*Brand, error) {
	newBrand, err := s.brandRepo.Add(ctx, brand)
	if err != nil {
		return nil, err
	}

	return newBrand, nil
}

func (s *service) GetBrand(ctx context.Context, brandID string) (*Brand, error) {
	brand, err := s.brandRepo.GetItemByID(ctx, brandID)
	if err != nil {
		return nil, err
	}

	return brand, nil
}

func (s *service) GetBrands(ctx context.Context, page, limit int64) (*BrandResult, error) {
	result, err := s.brandRepo.GetItems(ctx, page, limit)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *service) UpdateBrand(ctx context.Context, brandID string, brand *Brand) error {
	err := s.brandRepo.UpdateItemByID(ctx, brandID, brand)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) DeleteBrand(ctx context.Context, brandID string) error {
	err := s.brandRepo.DeleteItemByID(ctx, brandID)
	if err != nil {
		return err
	}

	return nil
}

//----------------CATEGORY----------------

func (s *service) AddCategory(ctx context.Context, ctgry *Category) (*Category, error) {
	ctgry, err := s.ctgryRepo.Add(ctx, ctgry)
	if err != nil {
		return nil, err
	}

	return ctgry, nil
}

func (s *service) GetCategory(ctx context.Context, ctgryID string) (*Category, error) {
	ctgry, err := s.ctgryRepo.GetItemByID(ctx, ctgryID)
	if err != nil {
		return nil, err
	}

	return ctgry, nil
}

func (s *service) GetCategories(ctx context.Context, page, limit int64) (*CategoryResult, error) {
	result, err := s.ctgryRepo.GetItems(ctx, page, limit)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *service) UpdateCategory(ctx context.Context, ctgryID string, ctgry *Category) error {
	err := s.ctgryRepo.UpdateItemByID(ctx, ctgryID, ctgry)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) DeleteCategory(ctx context.Context, ctgryID string) error {
	err := s.ctgryRepo.DeleteItemByID(ctx, ctgryID)
	if err != nil {
		return err
	}

	return nil
}

//----------------SUPPLIER----------------

func (s *service) AddSupplier(ctx context.Context, spplr *Supplier) (*Supplier, error) {
	newSpplr, err := s.spplrRepo.Add(ctx, spplr)
	if err != nil {
		return nil, err
	}

	return newSpplr, nil
}

func (s *service) GetSupplier(ctx context.Context, spplrID string) (*Supplier, error) {
	spllr, err := s.spplrRepo.GetItemByID(ctx, spplrID)
	if err != nil {
		return nil, err
	}

	return spllr, nil
}

func (s *service) GetSuppliers(ctx context.Context, page, limit int64) (*SupplierResult, error) {
	result, err := s.spplrRepo.GetItems(ctx, page, limit)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *service) UpdateSupplier(ctx context.Context, spplrID string, spplr *Supplier) error {
	err := s.spplrRepo.UpdateItemByID(ctx, spplrID, spplr)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) DeleteSupplier(ctx context.Context, spplrID string) error {
	err := s.spplrRepo.DeleteItemByID(ctx, spplrID)
	if err != nil {
		return err
	}

	return nil
}

//----------------PRODUCT----------------

func (s *service) AddProduct(ctx context.Context, product *Product) (*Product, error) {
	product, err := s.productRepo.Add(ctx, product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *service) GetProduct(ctx context.Context, productID string) (*Product, error) {
	product, err := s.productRepo.GetItemByID(ctx, productID)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *service) GetProducts(ctx context.Context, filterParams FilterProductsParams) (*ProductResult, error) {
	itemsResult, err := s.productRepo.GetItems(ctx, filterParams)
	if err != nil {
		return nil, err
	}

	if itemsResult == nil {
		return nil, nil
	}

	var products []Product

	for _, product := range itemsResult.Products {
		p, err := s.GetProduct(ctx, product.ID)
		if err != nil {
			return nil, err
		}

		if p == nil {
			continue
		}

		products = append(products, *p)
	}

	itemsResult.Products = products
	return itemsResult, nil
}

func (s *service) UpdateProduct(ctx context.Context, productID string, product *Product) error {
	err := s.productRepo.UpdateItemByID(ctx, productID, product)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) DeleteProduct(ctx context.Context, productID string) error {
	err := s.productRepo.DeleteItemByID(ctx, productID)
	if err != nil {
		return err
	}

	return nil
}
