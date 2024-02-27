package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

var url string = "https://sedeaplicaciones.minetur.gob.es/ServiciosRESTCarburantes/PreciosCarburantes/EstacionesTerrestres/"

// Fetch the data from `url` and return its body as a string
func getData(url string) (string, error) {
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

// API response
type Response struct {
	Fecha             string          `json:"Fecha"`
	ListaGasolineras  []RawGasStation `json:"ListaEESSPrecio"`
	Nota              string          `json:"Nota"`
	ResultadoConsulta string          `json:"ResultadoConsulta"`
}

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
	IDEESS    int
	Rotulo    string
	Direccion string
	Localidad string
	IDMun     int
	Municipio string
	IDProv    int
	Provincia string
	IDCCAA    int
	CP        string
	Latitud   float64
	Longitud  float64
	Horario   string
	Margen    string
	Remisión  string
	TipoVenta string

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

func getRawGasStations(url string) []GasStation {
	body, err := getData(url)
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
	for _, rawStation := range gasStations.ListaGasolineras {
		IDEESS, _ := strconv.Atoi(rawStation.IDEESS)
		IDMun, _ := strconv.Atoi(rawStation.IDMun)
		IDProv, _ := strconv.Atoi(rawStation.IDProv)
		IDCCAA, _ := strconv.Atoi(rawStation.IDCCAA)
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
	}

	return stations
}

func main() {
	gasStations := getRawGasStations(url)

	for _, station := range gasStations {
		fmt.Println(station)
	}
}
