package main

import (
	"database/sql"
	"io"
	"log"
	"net/http"
	"time"
)

const url string = "https://sedeaplicaciones.minetur.gob.es/ServiciosRESTCarburantes/PreciosCarburantes/EstacionesTerrestres/"
const dbPath string = "db.sqlite3"

// Fetch the data from `url` and return its body as a string
func GetData(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func main() {
	// Measure request and parsing time
	start := time.Now()
	gasStations, GasPrices := GetGasStationsGasPrices(url)
	log.Printf("Request and parsing time: %v", time.Since(start))

	// Measure database time
	beforeDB := time.Now()
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	dbGasStation := NewDBGasStation(db)
	dbGasStation.Migrate()

	for _, station := range gasStations {
		if _, err := dbGasStation.GetByID(int64(station.IDEESS)); err == nil {
			_, err := dbGasStation.Update(int64(station.IDEESS), station)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			_, err := dbGasStation.Create(station)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	dbGasPrice := NewDBGasPrice(db)
	dbGasPrice.Migrate()

	for _, price := range GasPrices {
		if _, err := dbGasPrice.GetIDByStationDate(int64(price.IDEESS), price.Date); err == nil {
			_, err := dbGasPrice.Update(int64(price.IDEESS), price, price.Date)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			_, err := dbGasPrice.Create(price, price.Date)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	log.Printf("Database time: %v", time.Since(beforeDB))
}
