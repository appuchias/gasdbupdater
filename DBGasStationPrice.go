package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type DBGasPrice SQLiteRepository

func NewDBGasPrice(db *sql.DB) *DBGasPrice {
	return (*DBGasPrice)(NewSQLiteRepository(db))
}

// Create the gas_stationprice table in the database
func (r *DBGasPrice) Migrate() error {
	_, err := r.db.Exec(`
		CREATE TABLE IF NOT EXISTS "gas_stationprice" (
			"id"	integer NOT NULL,
			"price_g95e5"	decimal,
			"price_g98e5"	decimal,
			"price_glp"	decimal,
			"price_goa"	decimal,
			"station_id"	integer NOT NULL,
			"date"	date NOT NULL,
			FOREIGN KEY("station_id") REFERENCES "gas_station"("id_eess") DEFERRABLE INITIALLY DEFERRED,
			CONSTRAINT "unique_station_date_combination" UNIQUE("station_id","date"),
			PRIMARY KEY("id" AUTOINCREMENT)
		)
	`)
	return err
}

// Insert a new gas price into the database
func (r *DBGasPrice) Create(price GasPrice, fecha string) (*GasPrice, error) {
	_, err := r.db.Exec(`
		INSERT INTO gas_stationprice (
			price_g95e5, price_g98e5, price_glp, price_goa, station_id, date
		) VALUES (?, ?, ?, ?, ?, ?)
	`, price.PrecioGasolina95E5, price.PrecioGasolina98E5, price.PrecioGLP, price.PrecioGasoleoA, price.IDEESS, fecha)

	if err != nil {
		return nil, err
	}

	return &price, nil
}

// Insert many gas prices into the database
func (r *DBGasPrice) CreateMany(prices []GasPrice, fecha string) ([]GasPrice, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare(`
		INSERT INTO gas_stationprice (
			price_g95e5, price_g98e5, price_glp, price_goa, station_id, date
		) VALUES (?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	for _, price := range prices {
		_, err := stmt.Exec(price.PrecioGasolina95E5, price.PrecioGasolina98E5, price.PrecioGLP, price.PrecioGasoleoA, price.IDEESS, fecha)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	tx.Commit()
	return prices, nil
}

// Get all gas prices from the database
func (r *DBGasPrice) All() ([]GasPrice, error) {
	rows, err := r.db.Query(`
		SELECT price_g95e5, price_g98e5, price_glp, price_goa, station_id, date
		FROM gas_stationprice
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var prices []GasPrice
	for rows.Next() {
		var price GasPrice
		err := rows.Scan(&price.PrecioGasolina95E5, &price.PrecioGasolina98E5, &price.PrecioGLP, &price.PrecioGasoleoA, &price.IDEESS)
		if err != nil {
			return nil, err
		}

		prices = append(prices, price)
	}

	return prices, nil
}

// Get all gas prices IDs from the database by date
func (r *DBGasPrice) AllStationIDsByDate(fecha string) ([]int64, error) {
	rows, err := r.db.Query(`
		SELECT station_id
		FROM gas_stationprice
		WHERE date = ?
	`, fecha)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ids []int64
	for rows.Next() {
		var id int64
		err := rows.Scan(&id)
		if err != nil {
			return nil, err
		}

		ids = append(ids, id)
	}

	return ids, nil
}

// Get a gas price from the database by ID
func (r *DBGasPrice) GetByID(id int64) (*GasPrice, error) {
	row := r.db.QueryRow(`
		SELECT price_g95e5, price_g98e5, price_glp, price_goa, station_id, date
		FROM gas_stationprice
		WHERE id = ?
	`, id)

	var price GasPrice
	err := row.Scan(&price.PrecioGasolina95E5, &price.PrecioGasolina98E5, &price.PrecioGLP, &price.PrecioGasoleoA, &price.IDEESS)
	if err != nil {
		return nil, err
	}

	return &price, nil
}

// Get all gas prices from the database by station ID
func (r *DBGasPrice) GetByStationID(station_id int64) ([]GasPrice, error) {
	rows, err := r.db.Query(`
		SELECT price_g95e5, price_g98e5, price_glp, price_goa, station_id, date
		FROM gas_stationprice
		WHERE station_id = ?
	`, station_id)
	if err != nil {
		return nil, err
	}

	var prices []GasPrice
	for rows.Next() {
		var price GasPrice
		err := rows.Scan(&price.PrecioGasolina95E5, &price.PrecioGasolina98E5, &price.PrecioGLP, &price.PrecioGasoleoA, &price.IDEESS)
		if err != nil {
			return nil, err
		}

		prices = append(prices, price)
	}

	return prices, nil
}

// Get a gas price from the database by station ID and date
func (r *DBGasPrice) GetIDByStationDate(id int64, fecha string) (int64, error) {
	row := r.db.QueryRow(`
		SELECT id
		FROM gas_stationprice
		WHERE station_id = ? AND date = ?
	`, id, fecha)

	var priceID int64
	err := row.Scan(&priceID)
	if err != nil {
		return 0, err
	}

	return priceID, nil
}

// Update a gas price in the database
func (r *DBGasPrice) Update(id int64, updated GasPrice, fecha string) (*GasPrice, error) {
	_, err := r.db.Exec(`
		UPDATE gas_stationprice
		SET price_g95e5 = ?, price_g98e5 = ?, price_glp = ?, price_goa = ?, station_id = ?, date = ?
		WHERE id = ?
	`, updated.PrecioGasolina95E5, updated.PrecioGasolina98E5, updated.PrecioGLP, updated.PrecioGasoleoA, updated.IDEESS, fecha, id)

	if err != nil {
		return nil, err
	}

	return &updated, nil
}

// Delete a gas price from the database
func (r *DBGasPrice) Delete(id int64) error {
	_, err := r.db.Exec(`
		DELETE FROM gas_stationprice
		WHERE id = ?
	`, id)

	return err
}
