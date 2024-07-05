package db

import (
	"database/sql"
	"fmt"
	"user-service-grpc/internal/core/user"

	"github.com/google/uuid"
)

type UserDBImpl struct {
	db *sql.DB
}

func (u *UserDBImpl) FindUserById(id uuid.UUID) (*user.User, error) {

	sqlQuery := fmt.Sprintf("SELECT id,name,address,city,state,country,pincode,phone_number,marital_status,height FROM user WHERE id=%#v", id)

	_, err := u.db.Exec(sqlQuery)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
