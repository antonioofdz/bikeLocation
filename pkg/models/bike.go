package models

import (
	"github.com/go-sql-driver/mysql"
)

type UserBikeDB struct {
	UserBook int
	Booked bool
	DateReturn mysql.NullTime
	DateRent mysql.NullTime
}

type BikeLocationDB struct {
	Lat int
	Lon int
	Address string
}

type BikeDB struct {
	Id      int
	Model    string
}

type BikeFull struct {
	BikeDB
	BikeLocationDB
	UserBikeDB
}

type BookBike struct {
	IdBike int
	IdUser int
}