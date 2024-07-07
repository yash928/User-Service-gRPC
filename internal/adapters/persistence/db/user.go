package db

import (
	"database/sql"
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

func (u *UserDBImpl) SaveUser(userDet *user.User) error {

	sqlQuery := `INSERT INTO users
			(id, name, address, city, state, country, pincode, phone_number, marital_status, height)
			VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		`
	_, err := u.db.Exec(sqlQuery, userDet.Id.String(), userDet.Name, userDet.Address, userDet.City, userDet.State, userDet.Country, userDet.Pincode, userDet.PhoneNo, userDet.MaritalStatus, userDet.Height)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserDBImpl) FindUserById(id uuid.UUID) (*user.User, error) {
	var userDet user.User
	sqlQuery := "SELECT id, name, address, city, state, country, pincode, phone_number, marital_status, height FROM users WHERE id=?"

	row := u.db.QueryRow(sqlQuery, id.String())

	// Scan the result into the User struct
	err := row.Scan(
		&userDet.Id,
		&userDet.Name,
		&userDet.Address,
		&userDet.City,
		&userDet.State,
		&userDet.Country,
		&userDet.Pincode,
		&userDet.PhoneNo,
		&userDet.MaritalStatus,
		&userDet.Height,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, user.ErrDocumentNotFound
		}
		return nil, err
	}

	return &userDet, nil

}

func (u *UserDBImpl) FindAllUsers() ([]user.User, error) {

	var users []user.User

	query := `SELECT id, name, address, city, state, country, pincode, phone_no, marital_status, height FROM users`
	rows, err := u.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var userDet user.User
		err := rows.Scan(
			&userDet.Id,
			&userDet.Name,
			&userDet.Address,
			&userDet.City,
			&userDet.State,
			&userDet.Country,
			&userDet.Pincode,
			&userDet.PhoneNo,
			&userDet.MaritalStatus,
			&userDet.Height,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, userDet)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
