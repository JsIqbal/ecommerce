package repo

import (
	"context"
	"fmt"

	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/jsiqbal/ecommerce/service"
)

// DB Models
type Supplier struct {
	ID                 string `db:"id"`
	Name               string `db:"name"`
	Email              string `db:"email"`
	Phone              string `db:"phone"`
	StatusID           int    `db:"status_id"`
	IsVerifiedSupplier bool   `db:"is_verified_supplier"`
	CreatedAt          int64  `db:"created_at"`
}

type SupplierRepo interface {
	service.SupplierRepo
}

type supplierRepo struct {
	db *sqlx.DB
}

func NewSupplierRepo(db *sqlx.DB) SupplierRepo {
	return &supplierRepo{
		db: db,
	}
}

func (r *supplierRepo) Add(ctx context.Context, spplr *service.Supplier) (*service.Supplier, error) {
	var newSpplr Supplier
	err := r.db.QueryRowContext(ctx,
		"INSERT INTO suppliers (name, email, phone, status_id, is_verified_Supplier, created_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *",
		spplr.Name, spplr.Email, spplr.Phone, spplr.StatusID, spplr.IsVerifiedSupplier, spplr.CreatedAt,
	).Scan(&newSpplr.ID, &newSpplr.Name, &newSpplr.Email, &newSpplr.Phone, &newSpplr.StatusID, &newSpplr.IsVerifiedSupplier, &newSpplr.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &service.Supplier{
		ID:                 newSpplr.ID,
		Name:               newSpplr.Name,
		Email:              newSpplr.Email,
		Phone:              newSpplr.Phone,
		StatusID:           newSpplr.StatusID,
		IsVerifiedSupplier: newSpplr.IsVerifiedSupplier,
		CreatedAt:          newSpplr.CreatedAt,
	}, nil
}

func (r *supplierRepo) GetItemByID(ctx context.Context, spplrID string) (*service.Supplier, error) {
	var spplr Supplier

	err := r.db.Get(&spplr, "SELECT id, name, email, phone, status_id, is_verified_supplier, created_at FROM suppliers WHERE id = $1", spplrID)
	if err == sql.ErrNoRows {
		// No product found
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &service.Supplier{
		ID:                 spplr.ID,
		Name:               spplr.Name,
		Email:              spplr.Email,
		Phone:              spplr.Phone,
		StatusID:           spplr.StatusID,
		IsVerifiedSupplier: spplr.IsVerifiedSupplier,
		CreatedAt:          spplr.CreatedAt,
	}, nil
}

func (r *supplierRepo) GetItems(ctx context.Context, page int64, limit int64) (*service.SupplierResult, error) {
	// calculate offset based on page and limit for pagination
	offset := (page - 1) * limit

	// fetch brands and total count
	var dbSpplrs []Supplier

	query := fmt.Sprintf("SELECT * FROM suppliers ORDER BY created_at DESC OFFSET %d LIMIT %d", offset, limit)
	err := r.db.SelectContext(ctx, &dbSpplrs, query)
	if err != nil {
		return nil, err
	}

	var totalCount int64
	err = r.db.GetContext(ctx, &totalCount, "SELECT COUNT(*) FROM suppliers")
	if err != nil {
		return nil, err
	}

	var spplrs []service.Supplier
	for _, dbSpplr := range dbSpplrs {
		spplrs = append(spplrs, service.Supplier{
			ID:                 dbSpplr.ID,
			Name:               dbSpplr.Name,
			Email:              dbSpplr.Email,
			Phone:              dbSpplr.Phone,
			StatusID:           dbSpplr.StatusID,
			IsVerifiedSupplier: dbSpplr.IsVerifiedSupplier,
			CreatedAt:          dbSpplr.CreatedAt,
		})
	}

	// return the result
	result := &service.SupplierResult{
		Suppliers: spplrs,
		Total:     totalCount,
		Page:      page,
		Limit:     limit,
	}

	return result, nil
}

func (r *supplierRepo) UpdateItemByID(ctx context.Context, spplrID string, spplr *service.Supplier) error {
	_, err := r.db.Exec(
		"UPDATE suppliers SET name = $1, email = $2, phone = $3, status_id = $4, is_verified_supplier = $5 WHERE id = $6",
		spplr.Name, spplr.Email, spplr.Phone, spplr.StatusID, spplr.IsVerifiedSupplier, spplrID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *supplierRepo) DeleteItemByID(ctx context.Context, spplrID string) error {
	_, err := r.db.Exec("DELETE FROM suppliers WHERE id = $1", spplrID)
	if err != nil {
		return err
	}

	return nil
}
