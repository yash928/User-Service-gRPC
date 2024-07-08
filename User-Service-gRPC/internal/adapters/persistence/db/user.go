package db

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
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

func (u *UserDBImpl) FindUserGivenId(ids []string) ([]user.User, error) {

	var users []user.User

	placeholders := make([]string, len(ids))
	idParams := make([]interface{}, len(ids))
	for i := range ids {
		placeholders[i] = "?"
		idParams[i] = ids[i]
	}
	query := fmt.Sprintf("SELECT id, name, address, city, state, country, pincode, phone_number, marital_status, height FROM users WHERE id IN (%s)", strings.Join(placeholders, ", "))
	log.Print(query, idParams)
	// Execute the query
	rows, err := u.db.Query(query, idParams...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Parse the result

	for rows.Next() {
		var user user.User
		if err := rows.Scan(&user.Id, &user.Name, &user.Address, &user.City, &user.State, &user.Country, &user.Pincode, &user.PhoneNo, &user.MaritalStatus, &user.Height); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (u *UserDBImpl) FindUserByFilter(filter user.Filter) ([]user.User, error) {
	query := "SELECT id, name, address, city, state, country, pincode, phone_number, marital_status, height FROM users WHERE TRUE"
	args := []interface{}{}
	conditions := []string{}

	if filter.Country != "" {
		conditions = append(conditions, "country = ?")
		args = append(args, filter.Country)
	}
	if filter.MaritalStatus != "" {
		conditions = append(conditions, "marital_status = ?")
		args = append(args, filter.MaritalStatus)
	}

	if len(conditions) > 0 {
		query += " AND " + strings.Join(conditions, " AND ")
	}

	// Execute the query
	rows, err := u.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Parse the result
	var users []user.User
	for rows.Next() {
		var user user.User
		if err := rows.Scan(&user.Id, &user.Name, &user.Address, &user.City, &user.State, &user.Country, &user.Pincode, &user.PhoneNo, &user.MaritalStatus, &user.Height); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
