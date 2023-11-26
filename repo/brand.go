package repo

import (
	"context"
	"fmt"
	"log"

	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/jsiqbal/ecommerce/logger"
	"github.com/jsiqbal/ecommerce/service"
)

// model
type Brand struct {
	ID        string `db:"id"`
	Name      string `db:"name"`
	StatusID  int    `db:"status_id"`
	CreatedAt int64  `db:"created_at"`
}

type BrandRepo interface {
	service.BrandRepo
}

type brandRepo struct {
	db *sqlx.DB
}

func NewBrandRepo(db *sqlx.DB) BrandRepo {
	return &brandRepo{
		db: db,
	}
}

func (r *brandRepo) Add(ctx context.Context, brand *service.Brand) (*service.Brand, error) {
	var newBrand Brand
	err := r.db.QueryRowContext(ctx,
		"INSERT INTO brands (name, status_id, created_at) VALUES ($1, $2, $3) RETURNING id, name, status_id, created_at",
		brand.Name, brand.StatusID, brand.CreatedAt,
	).Scan(&newBrand.ID, &newBrand.Name, &newBrand.StatusID, &newBrand.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &service.Brand{
		ID:        newBrand.ID,
		Name:      newBrand.Name,
		StatusID:  newBrand.StatusID,
		CreatedAt: newBrand.CreatedAt,
	}, nil
}

func (r *brandRepo) GetItemByID(ctx context.Context, brandID string) (*service.Brand, error) {
	log.Println("hello")
	var brand Brand

	err := r.db.Get(&brand, "SELECT * FROM brands WHERE id = $1", brandID)
	if err == sql.ErrNoRows {
		// No product found
		logger.Error(ctx, "cannot find brand", err)
		return nil, nil
	} else if err != nil {
		logger.Error(ctx, "error in finding brand", err)
		return nil, err
	}

	logger.Info(ctx, "found brand", brand)

	return &service.Brand{
		ID:        brand.ID,
		Name:      brand.Name,
		StatusID:  brand.StatusID,
		CreatedAt: brand.CreatedAt,
	}, nil
}

func (r *brandRepo) GetItems(ctx context.Context, page int64, limit int64) (*service.BrandResult, error) {
	// calculate offset based on page and limit for pagination
	offset := (page - 1) * limit

	// fetch brands and total count
	var dbBrands []Brand

	query := fmt.Sprintf("SELECT * FROM brands ORDER BY created_at DESC OFFSET %d LIMIT %d", offset, limit)
	err := r.db.SelectContext(ctx, &dbBrands, query)
	if err != nil {
		return nil, err
	}

	var totalCount int64
	err = r.db.GetContext(ctx, &totalCount, "SELECT COUNT(*) FROM brands")
	if err != nil {
		return nil, err
	}

	var brands []service.Brand
	for _, dbBrand := range dbBrands {
		brands = append(brands, service.Brand{
			ID:        dbBrand.ID,
			Name:      dbBrand.Name,
			StatusID:  dbBrand.StatusID,
			CreatedAt: dbBrand.CreatedAt,
		})
	}

	// return the result
	result := &service.BrandResult{
		Brands: brands,
		Total:  totalCount,
		Page:   page,
		Limit:  limit,
	}

	return result, nil
}

func (r *brandRepo) UpdateItemByID(ctx context.Context, brandID string, brand *service.Brand) error {
	_, err := r.db.Exec(
		"UPDATE brands SET name = $1, status_id = $2 WHERE id = $3",
		brand.Name, brand.StatusID, brandID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *brandRepo) DeleteItemByID(ctx context.Context, brandID string) error {
	_, err := r.db.Exec("DELETE FROM brands WHERE id = $1", brandID)
	if err != nil {
		return err
	}

	return nil
}
