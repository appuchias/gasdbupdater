package main

import (
	"encoding/json"
	"log"
	"regexp"
	"strconv"
	"time"
)

// Gas station (only strings)
type RawGasStation struct {
	IDEESS    string `json:"IDEESS"`
	Rotulo    string `json:"Rótulo"`
	Direccion string `json:"Dirección"`
	Localidad string `json:"Localidad"`
	IDMun     string `json:"IDMunicipio"`
	Municipio string `json:"Municipio"`
	IDProv    string `json:"IDProvincia"`
	Provincia string `json:"Provincia"`
	IDCCAA    string `json:"IDCCAA"`
	CP        string `json:"C.P."`
	Latitud   string `json:"Latitud"`
	Longitud  string `json:"Longitud (WGS84)"`
	Horario   string `json:"Horario"`
	Margen    string `json:"Margen"`
	Remisión  string `json:"Remisión"`
	TipoVenta string `json:"Tipo Venta"`

	PrecioGasoleoA            string `json:"Precio Gasoleo A"`
	PrecioGasoleoB            string `json:"Precio Gasoleo B"`
	PrecioGasoleoPremium      string `json:"Precio Gasoleo Premium"`
	PrecioGasolina95E10       string `json:"Precio Gasolina 95 E10"`
	PrecioGasolina95E5        string `json:"Precio Gasolina 95 E5"`
	PrecioGasolina95E5Premium string `json:"Precio Gasolina 95 E5 Premium"`
	PrecioGasolina98E10       string `json:"Precio Gasolina 98 E10"`
	PrecioGasolina98E5        string `json:"Precio Gasolina 98 E5"`
	PrecioGasolina98E5Premium string `json:"Precio Gasolina 98 E5 Premium"`
	PrecioGLP                 string `json:"Precio Gases licuados del petróleo"`
	PrecioGNC                 string `json:"Precio Gas Natural Comprimido"`
	PrecioGNL                 string `json:"Precio Gas Natural Licuado"`
	PrecioHidrogeno           string `json:"Precio Hidrogeno"`
}

// Gas station
type GasStation struct {
	IDEESS    int64
	Rotulo    string
	Direccion string
	Localidad string
	IDMun     int64
	Municipio string
	IDProv    int64
	Provincia string
	IDCCAA    int64
	CP        string
	Latitud   float64
	Longitud  float64
	Horario   string
	Margen    string
	Remisión  string
	TipoVenta string
}

type GasPrice struct {
	IDEESS int64
	Date   string

	PrecioGasoleoA            float64
	PrecioGasoleoB            float64
	PrecioGasoleoPremium      float64
	PrecioGasolina95E10       float64
	PrecioGasolina95E5        float64
	PrecioGasolina95E5Premium float64
	PrecioGasolina98E10       float64
	PrecioGasolina98E5        float64
	PrecioGasolina98E5Premium float64
	PrecioGLP                 float64
	PrecioGNC                 float64
	PrecioGNL                 float64
	PrecioHidrogeno           float64
}

func getRawGasStations(url string) ([]GasStation, []GasPrice) {
	body, err := GetData(url)
	if err != nil {
		log.Fatal(err)
	}

	// Fix float separators (, -> .)
	body = regexp.MustCompile(`(\d+),(\d+)`).ReplaceAllString(body, `$1.$2`)

	// Decode JSON into structs
	var gasStations Response
	err = json.Unmarshal([]byte(body), &gasStations)
	if err != nil {
		log.Fatal(err)
	}

	// Convert RawGasStation to GasStation
	var stations []GasStation
	var prices []GasPrice
	for _, rawStation := range gasStations.ListaGasolineras {
		IDEESS, _ := strconv.ParseInt(rawStation.IDEESS, 10, 64)
		IDMun, _ := strconv.ParseInt(rawStation.IDMun, 10, 64)
		IDProv, _ := strconv.ParseInt(rawStation.IDProv, 10, 64)
		IDCCAA, _ := strconv.ParseInt(rawStation.IDCCAA, 10, 64)
		Latitud, _ := strconv.ParseFloat(rawStation.Latitud, 64)
		Longitud, _ := strconv.ParseFloat(rawStation.Longitud, 64)

		PrecioGasoleoA, _ := strconv.ParseFloat(rawStation.PrecioGasoleoA, 64)
		PrecioGasoleoB, _ := strconv.ParseFloat(rawStation.PrecioGasoleoB, 64)
		PrecioGasoleoPremium, _ := strconv.ParseFloat(rawStation.PrecioGasoleoPremium, 64)
		PrecioGasolina95E10, _ := strconv.ParseFloat(rawStation.PrecioGasolina95E10, 64)
		PrecioGasolina95E5, _ := strconv.ParseFloat(rawStation.PrecioGasolina95E5, 64)
		PrecioGasolina95E5Premium, _ := strconv.ParseFloat(rawStation.PrecioGasolina95E5Premium, 64)
		PrecioGasolina98E10, _ := strconv.ParseFloat(rawStation.PrecioGasolina98E10, 64)
		PrecioGasolina98E5, _ := strconv.ParseFloat(rawStation.PrecioGasolina98E5, 64)
		PrecioGasolina98E5Premium, _ := strconv.ParseFloat(rawStation.PrecioGasolina98E5Premium, 64)
		PrecioGLP, _ := strconv.ParseFloat(rawStation.PrecioGLP, 64)
		PrecioGNC, _ := strconv.ParseFloat(rawStation.PrecioGNC, 64)
		PrecioGNL, _ := strconv.ParseFloat(rawStation.PrecioGNL, 64)
		PrecioHidrogeno, _ := strconv.ParseFloat(rawStation.PrecioHidrogeno, 64)

		station := GasStation{
			IDEESS:    IDEESS,
			Rotulo:    rawStation.Rotulo,
			Direccion: rawStation.Direccion,
			Localidad: rawStation.Localidad,
			IDMun:     IDMun,
			Municipio: rawStation.Municipio,
			IDProv:    IDProv,
			Provincia: rawStation.Provincia,
			IDCCAA:    IDCCAA,
			CP:        rawStation.CP,
			Latitud:   Latitud,
			Longitud:  Longitud,
			Horario:   rawStation.Horario,
			Margen:    rawStation.Margen,
			Remisión:  rawStation.Remisión,
			TipoVenta: rawStation.TipoVenta,
		}
		price := GasPrice{
			IDEESS: IDEESS,
			Date:   time.Now().Format(time.DateOnly),

			PrecioGasoleoA:            PrecioGasoleoA,
			PrecioGasoleoB:            PrecioGasoleoB,
			PrecioGasoleoPremium:      PrecioGasoleoPremium,
			PrecioGasolina95E10:       PrecioGasolina95E10,
			PrecioGasolina95E5:        PrecioGasolina95E5,
			PrecioGasolina95E5Premium: PrecioGasolina95E5Premium,
			PrecioGasolina98E10:       PrecioGasolina98E10,
			PrecioGasolina98E5:        PrecioGasolina98E5,
			PrecioGasolina98E5Premium: PrecioGasolina98E5Premium,
			PrecioGLP:                 PrecioGLP,
			PrecioGNC:                 PrecioGNC,
			PrecioGNL:                 PrecioGNL,
			PrecioHidrogeno:           PrecioHidrogeno,
		}

		stations = append(stations, station)
		prices = append(prices, price)
	}

	return stations, prices
}
