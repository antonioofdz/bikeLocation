package handlers

import (
	"fmt"

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
		return nil,err
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