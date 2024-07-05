package db

import (
	"database/sql"
	"fmt"
	"user-service-grpc/internal/adapters/ports"
	"user-service-grpc/internal/core/user"

	"github.com/google/uuid"
)

type UserDBImpl struct {
	db *sql.DB
}

func NewUserDb(db *sql.DB) ports.UserDB {
	return &UserDBImpl{
		db: db,
	}
}

func (u *UserDBImpl) FindUserById(id uuid.UUID) (*user.User, error) {

	sqlQuery := fmt.Sprintf("SELECT id,name,address,city,state,country,pincode,phone_number,marital_status,height FROM user WHERE id=%#v", id)

	_, err := u.db.Exec(sqlQuery)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
