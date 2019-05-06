
package models

type UserDBCredentials struct {
	Username string
	Password string
}

type UserDBToken struct {
	Token string
}

type UserDB struct {
	UserBasic
	UserDBToken
	Name    string
	Surname string
	Email   string
}

type UserBasic struct {
	Id int
}