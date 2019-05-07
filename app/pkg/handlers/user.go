package handlers

import (
	"fmt"
	"errors"

	_ "github.com/go-sql-driver/mysql"

	"github.com/antonioofdz/personalProjectDra/pkg/models"
	"github.com/antonioofdz/personalProjectDra/pkg/database"
)

func getUserCredentials(userCredentials *models.UserDBCredentials) (*models.UserDBToken, error) {
	sqlGetUserCredentials := `SELECT users.token 
							FROM users 
									inner join users_credentials on users_credentials.id_user = users.id 
							WHERE users_credentials.user=? and users_credentials.password=?`
	db, err := database.Open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var userDB models.UserDBToken
	err = db.QueryRow(sqlGetUserCredentials, &userCredentials.Username, &userCredentials.Password).Scan(&userDB.Token)
	if err != nil {
		return nil, err
	}

	return &userDB, nil
}

func getUserUserByToken(token string) (*models.UserDB, error) {
	fmt.Println(token)
	sqlGetUserCredentials :=    `SELECT users.id, users.name, 
									 users.surname, users.email, users.token 
								FROM users WHERE token=?`
	db, err := database.Open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var userDB models.UserDB
	err = db.QueryRow(sqlGetUserCredentials, token).Scan(&userDB.Id, &userDB.Name, &userDB.Surname, &userDB.Email, &userDB.Token)
	if err != nil {
		return nil, err
	}

	return &userDB, nil
}

func signInUser(signInUser *models.SignInUserDB) error {
	sqlCheckIfExistsUser := `SELECT count(*) from users_credentials where user = ?`
	if !checkValidPassword(signInUser.Password) {
		err := errors.New("INVALID PASSWORD")
		return err
	}

	db, err := database.Open()
	if err != nil {
		return err
	}
	defer db.Close()

	var countUsers int
	err = db.QueryRow(sqlCheckIfExistsUser, signInUser.Username).Scan(&countUsers)
	if err != nil {
		return err
	}

	if countUsers > 0 {
		err = errors.New("USERNAME ALREADY EXISTS")
		return err
	}

	return insertUser(signInUser)
}

func insertUser(signInUser *models.SignInUserDB) (error) {
	db, err := database.Open()
	if err != nil {
		return err
	}
	defer db.Close()

	sqlInsertUser := `INSERT INTO users (name, surname, email, token) VALUES (?, ?, ?, UUID());`
	stmt, err := db.Prepare(sqlInsertUser)
    if err != nil {
        err = errors.New("CANNOT PREPARE STATMENT")
		return err
	}
	
	data, err := stmt.Exec(signInUser.Username, signInUser.Surname, signInUser.Email);
    if err != nil {
        err = errors.New("CANNOT PREPARE STATMENT")
		return err
	}
	signInUser.Id, err = data.LastInsertId()
	if err != nil {
		err = errors.New("ERROR GETTING LAST ID INSERT")
		return err
	}

	sqlInsertUserCredentials := `INSERT INTO users_credentials (user, password, id_user) VALUES (?, ?, ?);`
	stmt, err = db.Prepare(sqlInsertUserCredentials)
    if err != nil {
        err = errors.New("CANNOT PREPARE STATMENT INSERT USER CREDENTIALS")
		return err
	}
	
	data, err = stmt.Exec(signInUser.Username, signInUser.Password, signInUser.Id);
    if err != nil {
        err = errors.New("CANNOT PREPARE STATMENT INSERT USER CREDENTIALS")
		return err
    }

	return nil
}