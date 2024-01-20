package database

import (
	"database/sql"
	"github.com/MikeMwita/ref-codegen/internal/core/entities"
)

type CodeRepository struct {
	db *sql.DB
}

func (r *CodeRepository) Save(code *entities.Code) error {
	stmt, err := r.db.Prepare("INSERT INTO codes (value, format) VALUES ($1, $2)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(code.Value, code.Format)
	if err != nil {
		return err
	}
	return nil
}

// GetByValue gets a code from the database by the value

func (r *CodeRepository) GetByValue(value string) (*entities.Code, error) {
	stmt, err := r.db.Prepare("SELECT value, format FROM codes WHERE value = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(value)
	value, format := "", ""
	err = row.Scan(&value, &format)
	if err != nil {
		return nil, err
	}
	code, err := entities.NewCode(value, format)
	if err != nil {
		return nil, err
	}
	return code, nil
}

func NewCodeRepository(db *sql.DB) *CodeRepository {
	return &CodeRepository{
		db: db,
	}
}
