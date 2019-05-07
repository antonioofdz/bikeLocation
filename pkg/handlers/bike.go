package handlers

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/antonioofdz/personalProjectDra/pkg/models"
	"github.com/antonioofdz/personalProjectDra/pkg/database"
)

// TRAER EL LISTADO DE BICICLETAS PARA PINTARLAS EN EL MAPA
func getListBikes() ([] models.BikeFull, error) {
	sqlListBikes := `SELECT b.id, b.model, bl.address, bl.lat, 
									bl.lon, ub.booked, ub.dateRent, ub.dateReturn 
								FROM bike b 
								INNER JOIN bike_location bl ON b.id = bl.id_bike 
								INNER JOIN user_bike ub ON ub.id_bike = b.id`

	db, err := database.Open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var bikes [] models.BikeFull
	rows, err := db.Query(sqlListBikes)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var b models.BikeFull
		if err := rows.Scan(&b.Id, &b.Model, &b.Address, &b.Lat, &b.Lon, &b.Booked, &b.DateRent, &b.DateReturn ); err != nil {
			continue
		}
		bikes = append(bikes, b)
	}
	
	return bikes, nil
}

// METODO PARA RESERVAR UNA BICICLETA
func bookBike(bookBikeModel *models.BookBike) (error) {
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

// METODO PARA CANCELAR UNA RESERVA
func endBookBike(bookBikeModel *models.BookBike) (error) {
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