package repositories

import (
	"fmt"
	"strings"

	"github.com/funukonta/shopping/internal/models"
	"github.com/jmoiron/sqlx"
)

type AuthRepo interface {
	Login(*models.Login) error
}

type authRepo struct {
	db *sqlx.DB
}

func NewAuthRepo(db *sqlx.DB) AuthRepo {
	return &authRepo{db: db}
}

func (r *authRepo) Login(login *models.Login) error {
	query := `select email,password from customers where email=$1`
	dataDb := new(models.Login)
	err := r.db.Get(dataDb, query, login.Email)
	if err != nil {
		if strings.Contains(err.Error(), "no rows") {
			return fmt.Errorf("email tidak ditemukan")
		}
		return err
	}

	if dataDb.Password != login.Password {
		return fmt.Errorf("password salah")
	}

	return nil
}
