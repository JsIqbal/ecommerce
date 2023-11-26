package repo

import (
	"context"
	"fmt"

	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/jsiqbal/ecommerce/logger"
	"github.com/jsiqbal/ecommerce/service"
)

// db model
type Category struct {
	ID        string         `db:"id"`
	Name      string         `db:"name"`
	ParentID  sql.NullString `db:"parent_id"`
	Sequence  sql.NullString `db:"sequence"`
	StatusID  int            `db:"status_id"`
	CreatedAt int64          `db:"created_at"`
}

type CategoryRepo interface {
	service.CategoryRepo
}

type categoryRepo struct {
	db *sqlx.DB
}

func NewCategoryRepo(db *sqlx.DB) CategoryRepo {
	return &categoryRepo{
		db: db,
	}
}

func (r *categoryRepo) Add(ctx context.Context, ctgry *service.Category) (*service.Category, error) {
	var parentID interface{}
	if ctgry.ParentID != "" {
		parentID = ctgry.ParentID
	} else {
		parentID = nil
	}

	var sequence interface{}
	if ctgry.Sequence != "" {
		sequence = ctgry.Sequence
	} else {
		sequence = nil
	}

	var newCtgry Category
	err := r.db.QueryRowContext(ctx,
		"INSERT INTO categories (name, parent_id, sequence, status_id, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING id, name, parent_id, sequence, status_id, created_at",
		ctgry.Name, parentID, sequence, ctgry.StatusID, ctgry.CreatedAt,
	).Scan(&newCtgry.ID, &newCtgry.Name, &newCtgry.ParentID, &newCtgry.Sequence, &newCtgry.StatusID, &newCtgry.CreatedAt)
	if err != nil {
		return nil, err
	}

	logger.Info(ctx, "db category", newCtgry)

	return &service.Category{
		ID:        newCtgry.ID,
		Name:      newCtgry.Name,
		ParentID:  newCtgry.ParentID.String,
		Sequence:  newCtgry.Sequence.String,
		StatusID:  newCtgry.StatusID,
		CreatedAt: newCtgry.CreatedAt,
	}, nil
}

func (r *categoryRepo) GetItemByID(ctx context.Context, ctgryID string) (*service.Category, error) {
	var ctgry Category

	err := r.db.Get(&ctgry, "SELECT id, name, parent_id, sequence, status_id, created_at FROM categories WHERE id = $1", ctgryID)
	if err == sql.ErrNoRows {
		// No category found
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	logger.Info(ctx, "category", ctgry)

	return &service.Category{
		ID:        ctgry.ID,
		Name:      ctgry.Name,
		ParentID:  ctgry.ParentID.String,
		Sequence:  ctgry.Sequence.String,
		StatusID:  ctgry.StatusID,
		CreatedAt: ctgry.CreatedAt,
	}, nil
}

func (r *categoryRepo) GetItems(ctx context.Context, page int64, limit int64) (*service.CategoryResult, error) {
	// calculate offset based on page and limit for pagination
	offset := (page - 1) * limit

	// fetch brands and total count
	var dbctgries []Category

	query := fmt.Sprintf("SELECT * FROM categories ORDER BY created_at DESC OFFSET %d LIMIT %d", offset, limit)
	err := r.db.SelectContext(ctx, &dbctgries, query)
	if err != nil {
		return nil, err
	}

	var totalCount int64
	err = r.db.GetContext(ctx, &totalCount, "SELECT COUNT(*) FROM categories")
	if err != nil {
		return nil, err
	}

	var ctries []service.Category
	for _, dbCtgry := range dbctgries {
		ctries = append(ctries, service.Category{
			ID:        dbCtgry.ID,
			Name:      dbCtgry.Name,
			ParentID:  dbCtgry.ParentID.String,
			Sequence:  dbCtgry.Sequence.String,
			StatusID:  dbCtgry.StatusID,
			CreatedAt: dbCtgry.CreatedAt,
		})
	}

	// return the result
	result := &service.CategoryResult{
		Categories: ctries,
		Total:      totalCount,
		Page:       page,
		Limit:      limit,
	}

	return result, nil
}

func (r *categoryRepo) UpdateItemByID(ctx context.Context, ctgryID string, ctgry *service.Category) error {
	var parentID interface{}
	if ctgry.ParentID != "" {
		parentID = ctgry.ParentID
	} else {
		parentID = nil
	}

	var sequence interface{}
	if ctgry.Sequence != "" {
		sequence = ctgry.Sequence
	} else {
		sequence = nil
	}

	_, err := r.db.Exec(
		"UPDATE categories SET name = $1, parent_id = $2, sequence = $3, status_id = $4 WHERE id = $5",
		ctgry.Name, parentID, sequence, ctgry.StatusID, ctgryID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *categoryRepo) DeleteItemByID(ctx context.Context, ctgryID string) error {
	_, err := r.db.Exec("DELETE FROM categories WHERE id = $1", ctgryID)
	if err != nil {
		return err
	}

	return nil
}
