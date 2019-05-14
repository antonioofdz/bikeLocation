package handlers

import (
	"github.com/antonioofdz/personalprojectdra/pkg/database"
	_ "github.com/go-sql-driver/mysql"
)

func InitDBHandler() error {
	return database.InitDB()
}
