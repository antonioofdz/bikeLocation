package models

import (
	"github.com/go-sql-driver/mysql"
)

type UserBikeDB struct {
	UserBook   int
	Booked     bool
	DateReturn mysql.NullTime
	DateRent   mysql.NullTime
}

type BikeLocationDB struct {
	Lat     float64
	Lon     float64
	Address string
}

type BikeId struct {
	Id int64
}
type BikeDB struct {
	BikeId
	Model string
}

type BikeFull struct {
	BikeDB
	BikeLocationDB
	UserBikeDB
}

type BikeTest struct {
	BikeDB
	BikeLocationDB
	Booked bool
}

type BookBike struct {
	IdBike int
	IdUser int
}
