package procesamientos

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	TDADict "tdas/diccionario"
	grafo "tp3/grafo"
)

const (
	AEROPUERTOS     = 0
	INICIAL         = 0
	CIUDAD          = 0
	FINAL           = 1
	VUELOS          = 1
	COD_AEROPUERTO  = 1
	TIEMPO_PROMEDIO = 2
	PRECIO          = 3
	CANT_VUELOS     = 4
)

func ProcesarArchivos(archivos []string) (grafo.Grafo[string, int], grafo.Grafo[string, int], grafo.Grafo[string, float64], TDADict.Diccionario[string, []string], TDADict.Diccionario[string, []string]) {
	archivo, _ := os.Open(archivos[AEROPUERTOS])

	grafoPrecios := grafo.CrearGrafo[string, int](false, true)
	grafoTiempo := grafo.CrearGrafo[string, int](false, true)
	grafoCantVuelos := grafo.CrearGrafo[string, float64](false, true)
	dictCiudades := TDADict.CrearHash[string, []string]()
	dictCoordenadas := TDADict.CrearHash[string, []string]()

	s := bufio.NewScanner(archivo)
	for s.Scan() {
		linea := s.Text()
		slice := strings.Split(linea, ",")
		grafoPrecios.AgregarVertice(slice[COD_AEROPUERTO])
		grafoTiempo.AgregarVertice(slice[COD_AEROPUERTO])
		grafoCantVuelos.AgregarVertice(slice[COD_AEROPUERTO])
		dictCoordenadas.Guardar(slice[COD_AEROPUERTO], slice[len(slice)-2:])
		if dictCiudades.Pertenece(slice[CIUDAD]) {
			arrAeropuertos := dictCiudades.Obtener(slice[CIUDAD])
			arrAeropuertos = append(arrAeropuertos, slice[COD_AEROPUERTO])
			dictCiudades.Guardar(slice[CIUDAD], arrAeropuertos)

		} else {
			arrAeropuertos := []string{}
			arrAeropuertos = append(arrAeropuertos, slice[COD_AEROPUERTO])
			dictCiudades.Guardar(slice[CIUDAD], arrAeropuertos)
		}
	}
	archivo.Close()
	archivo, _ = os.Open(archivos[VUELOS])
	s = bufio.NewScanner(archivo)
	for s.Scan() {
		linea := s.Text()
		slice := strings.Split(linea, ",")
		tiempoPromedio, _ := strconv.Atoi(slice[TIEMPO_PROMEDIO])
		Precio, _ := strconv.Atoi(slice[PRECIO])
		CantVuelos, _ := strconv.ParseFloat(slice[CANT_VUELOS], 64)
		grafoTiempo.AgregarArista(slice[INICIAL], slice[FINAL], tiempoPromedio)
		grafoPrecios.AgregarArista(slice[INICIAL], slice[FINAL], Precio)
		grafoCantVuelos.AgregarArista(slice[INICIAL], slice[FINAL], 1/CantVuelos)
	}
	return grafoTiempo, grafoPrecios, grafoCantVuelos, dictCiudades, dictCoordenadas
}
