package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type DBGasStation SQLiteRepository

func NewDBGasStation(db *sql.DB) *DBGasStation {
	return (*DBGasStation)(NewSQLiteRepository(db))
}

func (r *DBGasStation) Migrate() error {
	_, err := r.db.Exec(`
		CREATE TABLE IF NOT EXISTS "gas_station" (
			"id_eess"	integer PRIMARY KEY NOT NULL,
			"company"	varchar(128) NOT NULL,
			"address"	varchar(128) NOT NULL,
			"schedule"	varchar(64) NOT NULL,
			"latitude"	varchar(10) NOT NULL,
			"longitude"	varchar(10) NOT NULL,
			"locality_id"	integer NOT NULL,
			"postal_code"	integer NOT NULL,
			"province_id"	integer NOT NULL,
			FOREIGN KEY("locality_id") REFERENCES "gas_locality"("id_mun") DEFERRABLE INITIALLY DEFERRED,
			FOREIGN KEY("province_id") REFERENCES "gas_province"("id_prov") DEFERRABLE INITIALLY DEFERRED
		)
	`)

	return err
}

func (r *DBGasStation) Create(station GasStation) (*GasStation, error) {
	_, err := r.db.Exec(`
		INSERT INTO gas_station (
			id_eess, company, address, schedule, latitude, longitude, locality_id, postal_code, province_id
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`, station.IDEESS, station.Rotulo, station.Direccion, station.Horario, station.Latitud, station.Longitud, station.IDMun, station.CP, station.IDProv)

	if err != nil {
		return nil, err
	}

	return &station, nil
}

func (r *DBGasStation) All() ([]GasStation, error) {
	rows, err := r.db.Query(`
		SELECT id_eess, company, address, schedule, latitude, longitude, locality_id, postal_code, province_id
		FROM gas_station
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stations []GasStation
	for rows.Next() {
		var station GasStation
		err := rows.Scan(&station.IDEESS, &station.Rotulo, &station.Direccion, &station.Horario, &station.Latitud, &station.Longitud, &station.IDMun, &station.CP, &station.IDProv)
		if err != nil {
			return nil, err
		}

		stations = append(stations, station)
	}

	return stations, nil
}

func (r *DBGasStation) GetByID(id int64) (*GasStation, error) {
	row := r.db.QueryRow(`
		SELECT id_eess, company, address, schedule, latitude, longitude, locality_id, postal_code, province_id
		FROM gas_station
		WHERE id_eess = ?
	`, id)

	var station GasStation
	err := row.Scan(&station.IDEESS, &station.Rotulo, &station.Direccion, &station.Horario, &station.Latitud, &station.Longitud, &station.IDMun, &station.CP, &station.IDProv)
	if err != nil {
		return nil, err
	}

	return &station, nil
}

func (r *DBGasStation) Update(id int64, updated GasStation) (*GasStation, error) {
	_, err := r.db.Exec(`
		UPDATE gas_station
		SET company = ?, address = ?, schedule = ?, latitude = ?, longitude = ?, locality_id = ?, postal_code = ?, province_id = ?
		WHERE id_eess = ?
	`, updated.Rotulo, updated.Direccion, updated.Horario, updated.Latitud, updated.Longitud, updated.IDMun, updated.CP, updated.IDProv, id)

	if err != nil {
		return nil, err
	}

	return &updated, nil
}

func (r *DBGasStation) Delete(id int64) error {
	_, err := r.db.Exec(`
		DELETE FROM gas_station
		WHERE id_eess = ?
	`, id)

	return err
}
