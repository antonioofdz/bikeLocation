package handlers

import (
	"errors"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/antonioofdz/personalprojectdra/pkg/database"
	"github.com/antonioofdz/personalprojectdra/pkg/models"
)

// TRAER EL LISTADO DE BICICLETAS PARA PINTARLAS EN EL MAPA
func getListBikes() ([]models.BikeTest, error) {
	/*sqlListBikes := `SELECT b.id, b.model, bl.address, bl.lat,
		bl.lon, ub.booked, ub.dateRent, ub.dateReturn
	FROM bike b
	INNER JOIN bike_location bl ON b.id = bl.id_bike
	INNER JOIN user_bike ub ON ub.id_bike = b.id`
	*/

	sqlListBikes := `SELECT  b.id, b.model, bl.address, bl.lat, 
									bl.lon, IFNULL(ub.booked, false)
								FROM bike b 
								INNER JOIN bike_location bl ON b.id = bl.id_bike 
								LEFT JOIN user_bike ub ON ub.id_bike = b.id `

	db, err := database.Open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var bikes []models.BikeTest
	rows, err := db.Query(sqlListBikes)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var b models.BikeTest
		//if err := rows.Scan(&b.Id, &b.Model, &b.Address, &b.Lat, &b.Lon, &b.Booked, &b.DateRent, &b.DateReturn); err != nil {
		if err := rows.Scan(&b.Id, &b.Model, &b.Address, &b.Lat, &b.Lon, &b.Booked); err != nil {
			continue
		}
		b.Address, err = GetAddress(b.Lat, b.Lon)
		if err != nil {
			b.Address = ""
		}
		bikes = append(bikes, b)
	}

	return bikes, nil
}

// METODO PARA RESERVAR UNA BICICLETA
func bookBike(bookBikeModel *models.BookBike) error {
	sqlBookBike := `INSERT INTO user_bike(booked, dateRent, dateReturn, id_user, id_bike) 
						VALUES (true, NOW(), null, ?, ?); `

	db, err := database.Open()
	if err != nil {
		fmt.Println("hola2", err)
	}
	defer db.Close()

	stmt, err := db.Prepare(sqlBookBike)
	if err != nil {
		log.Fatal("Cannot prepare DB statement", err)
	}

	if _, err := stmt.Exec(bookBikeModel.IdUser, bookBikeModel.IdBike); err != nil {
		log.Fatal("Cannot run insert statement", err)
	}

	return nil
}

// METODO PARA RESERVAR UNA BICICLETA
func getBike(id int64) (*models.BikeTest, error) {
	sqlListBikes := `SELECT b.id, b.model, bl.address, bl.lat, 
						bl.lon, IFNULL(ub.booked, false)
					FROM bike b 
					LEFT JOIN bike_location bl ON b.id = bl.id_bike 
					LEFT JOIN user_bike ub ON ub.id_bike = b.id 
					WHERE b.id = ? LIMIT 1`

	db, err := database.Open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	fmt.Println("ID BIKE : ", id)
	var b models.BikeTest
	err = db.QueryRow(sqlListBikes, id).Scan(&b.Id, &b.Model, &b.Address, &b.Lat, &b.Lon, &b.Booked)
	if err != nil {
		fmt.Println("ERR ", err)
		return nil, err
	}

	b.Address, err = GetAddress(b.Lat, b.Lon)
	if err != nil {
		b.Address = ""
	}

	return &b, nil

}

// METODO PARA CANCELAR UNA RESERVA
func endBookBike(bookBikeModel *models.BookBike) error {
	sqlEndBookBike := `UPDATE user_bike SET booked = false, dateReturn=NOW() 
							WHERE id_user=? and id_bike=?;`

	db, err := database.Open()
	if err != nil {
		fmt.Println("hola2", err)
	}
	defer db.Close()

	stmt, err := db.Prepare(sqlEndBookBike)
	if err != nil {
		log.Fatal("Cannot prepare DB statement", err)
	}

	if _, err := stmt.Exec(bookBikeModel.IdUser, bookBikeModel.IdBike); err != nil {
		log.Fatal("Cannot run insert statement", err)
	}

	return nil
}

func UpdateBike(bike *models.BikeFull) error {
	sql := `UPDATE bike set model=? where id=?`
	db, err := database.Open()
	if err != nil {
		return errors.New("error opening the database")
	}
	defer db.Close()

	dbPrepare, err := db.Prepare(sql)
	if err != nil {
		return errors.New("error executing query")
	}

	_, err = dbPrepare.Exec(&bike.Model, &bike.Id)
	if err != nil {
		err = errors.New("CANNOT PREPARE STATMENT")
		return err
	}

	sql = `UPDATE bike_location set lat=?, lon=?, address=? WHERE id_bike=?`
	dbPrepareLocation, err := db.Prepare(sql)
	if err != nil {
		return errors.New("error executing query")
	}

	_, err = dbPrepareLocation.Exec(&bike.Lat, &bike.Lon, &bike.Address, &bike.Id)
	if err != nil {
		err = errors.New("CANNOT PREPARE STATMENT")
		return err
	}

	return nil
}

func DeleteBike(id int64) error {
	sql := `DELETE FROM bike_location WHERE id_bike=?`
	db, err := database.Open()
	if err != nil {
		return errors.New("error opening the database")
	}
	defer db.Close()

	dbPrepare, err := db.Prepare(sql)
	if err != nil {
		return errors.New("error executing query")
	}

	_, err = dbPrepare.Exec(id)
	if err != nil {
		err = errors.New("CANNOT PREPARE STATMENT")
		return err
	}

	sql = `delete from bike where id=?`
	dbPrepareLocation, err := db.Prepare(sql)
	if err != nil {
		return errors.New("error executing query")
	}

	_, err = dbPrepareLocation.Exec(id)
	if err != nil {
		err = errors.New("CANNOT PREPARE STATMENT")
		return err
	}

	return nil
}

func InsertBike(bike *models.BikeTest) error {
	sql := `INSERT INTO bike (model) VALUES (?);`
	db, err := database.Open()
	if err != nil {
		return errors.New("error opening the database")
	}
	defer db.Close()

	dbPrepare, err := db.Prepare(sql)
	if err != nil {
		return errors.New("error executing query")
	}

	data, err := dbPrepare.Exec(&bike.Model)
	if err != nil {
		err = errors.New("CANNOT PREPARE STATMENT")
		return err
	}

	id, err := data.LastInsertId()
	if err != nil {
		err = errors.New("ERROR GETTING LAST ID INSERT")
		return err
	}

	sql = `INSERT INTO bike_location (lat, lon, address, id_bike) VALUES (?, ?, ?, ?);`
	dbPrepareLocation, err := db.Prepare(sql)
	if err != nil {
		return errors.New("error executing query")
	}

	_, err = dbPrepareLocation.Exec(&bike.Lat, &bike.Lon, &bike.Address, id)
	if err != nil {
		err = errors.New("CANNOT PREPARE STATMENT")
		return err
	}

	return nil
}
