package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Open() (*sql.DB, error) {
	db, err := sql.Open("mysql", "dev-user:dev-password@/ermakina")
	if err != nil {
		return db, err
	}

	return db, nil
}

func getUserCredentials(userCredentials *UserDBCredentials) (*UserDB, error) {
	sqlGetUserCredentials := "SELECT id, name, surname, email, token FROM Users where username=@p1 and password=p2"
	db, err := Open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var userDB UserDB
	err = db.QueryRow(sqlGetUserCredentials, &userCredentials.Username, &userCredentials.Password).Scan(&userDB.Id, &userDB.Name, &userDB.Surname, &userDB.Email)
	if err != nil {
		return nil, err
	}

	return &userDB, nil
}
