// frameworks/database/code_repository.go

package database

import (
	"database/sql"
	"github.com/MikeMwita/ref-codegen/internal/core/entities"
)

// CodeRepository is a type that implements the code repository interface using a database
type CodeRepository struct {
	db *sql.DB // the database connection
}

// NewCodeRepository creates a new CodeRepository with the given database connection
func NewCodeRepository(db *sql.DB) *CodeRepository {
	return &CodeRepository{
		db: db,
	}
}

// Save saves a code to the database
func (r *CodeRepository) Save(code *entities.Code) error {
	// prepare the SQL statement
	stmt, err := r.db.Prepare("INSERT INTO codes (value, format) VALUES ($1, $2)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	// execute the SQL statement with the code values
	_, err = stmt.Exec(code.Value, code.Format)
	if err != nil {
		return err
	}
	return nil
}

// GetByValue gets a code from the database by the value
func (r *CodeRepository) GetByValue(value string) (*entities.Code, error) {
	// prepare the SQL statement
	stmt, err := r.db.Prepare("SELECT value, format FROM codes WHERE value = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	// execute the SQL statement with the value
	row := stmt.QueryRow(value)
	// scan the row into variables
	var value, format string
	err = row.Scan(&value, &format)
	if err != nil {
		return nil, err
	}
	// create a new Code entity from the variables
	code, err := entities.NewCode(value, format)
	if err != nil {
		return nil, err
	}
	return code, nil
}
