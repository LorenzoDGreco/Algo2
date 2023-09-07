package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	TDAPila "tdas/pila"
	op "tp3/operaciones"
	proc "tp3/procesamientos"
)

const (
	ACCION       = 0
	INICIO       = 0
	FINAL        = 1
	TIPOCAMINO   = 1
	RUTA_ARCHIVO = 1
	AEROPUERTOS  = 1
	VIAJE        = 1
	CANTIDAD     = 1
	MAX_SPLIT    = 2
	VIAJES       = 1
)

func main() {
	parametros := os.Args
	if len(parametros) != 3 {
		panic("ERROR: NO SE INGRESARON LOS PARAMETROS CORRECTOS")
	}

	stdin := os.Stdin
	scan := bufio.NewScanner(stdin)

	grafoTiempo, grafoPrecios, grafoCantVuelos, DictCiudades, dictCoordenadas := proc.ProcesarArchivos(parametros[AEROPUERTOS:])
	pilaMinima := TDAPila.CrearPilaDinamica[string]()

	for scan.Scan() {
		linea := scan.Text()
		comandos := strings.Split(linea, " ")

		switch comandos[ACCION] {

		case "camino_mas":
			comandos = strings.SplitN(linea, ",", MAX_SPLIT)
			viajes := strings.Split(comandos[VIAJES], ",")
			comandos = strings.Split(comandos[ACCION], " ")

			if comandos[TIPOCAMINO] == "barato" {
				pilaMinima = op.ObtenerCaminoMinimo(grafoPrecios, DictCiudades, viajes[INICIO], viajes[FINAL], true)
			} else if comandos[TIPOCAMINO] == "rapido" {
				pilaMinima = op.ObtenerCaminoMinimo(grafoTiempo, DictCiudades, viajes[INICIO], viajes[FINAL], true)
			}

		case "camino_escalas":
			comandos = strings.SplitN(linea, " ", MAX_SPLIT)
			viajes := strings.Split(comandos[VIAJES], ",")

			pilaMinima = op.MenorEscalas(grafoTiempo, DictCiudades, viajes[INICIO], viajes[FINAL])

		case "centralidad":
			Kint, _ := strconv.Atoi(comandos[CANTIDAD])
			op.Centralidad(grafoCantVuelos, Kint)

		case "nueva_aerolinea":
			op.NuevaAerolinea(grafoPrecios, grafoTiempo, grafoCantVuelos, comandos[RUTA_ARCHIVO])

		case "itinerario":
			op.ItinerarioCultural(grafoTiempo, DictCiudades, comandos[RUTA_ARCHIVO])

		case "exportar_kml":
			op.ExportarKml(pilaMinima, dictCoordenadas, comandos[RUTA_ARCHIVO])
		}
	}
}
