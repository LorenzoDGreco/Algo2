package operaciones

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"

	Heap "tdas/cola_prioridad"
	Dict "tdas/diccionario"
	Pila "tdas/pila"
	grafo "tp3/grafo"
)

const (
	PRIORITARIO       = 0
	LATITUD           = 0
	PRIMERO           = 0
	MENOS_PRIORITARIO = 1
	SIN_PESO          = 1
	LONGITUD          = 1
)

type tupla struct {
	numCmp  int
	origen  string
	destino string
}

func crearTupla(peso int, origen, destino string) tupla {
	tupla := new(tupla)
	tupla.numCmp = peso
	tupla.origen = origen
	tupla.destino = destino
	return *tupla
}

func obtenerViajeMin(disTotales []tupla) (string, string) {
	min := minimo(disTotales)
	for i := 0; i < len(disTotales); i++ {
		if disTotales[i].numCmp == min {
			return disTotales[i].origen, disTotales[i].destino
		}
	}
	return disTotales[0].origen, disTotales[0].destino
}

func minimo(arreglo []tupla) int {
	if len(arreglo) == 1 {
		return arreglo[0].numCmp
	}
	restante := minimo(arreglo[1:])
	return min(restante, arreglo[0].numCmp)
}

func min(a, b int) int {
	if a >= b {
		return b
	}
	return a
}

func ObtenerCaminoMinimo(grafoVuelos grafo.Grafo[string, int], dictCiudades Dict.Diccionario[string, []string],
	origen, destino string, imprimir bool) Pila.Pila[string] {

	aeropuertosOrigen := dictCiudades.Obtener(origen)
	aeropuertosDestino := dictCiudades.Obtener(destino)
	DictPadres := Dict.CrearHash[string, Dict.Diccionario[string, string]]()
	distanciasTotales := []tupla{}
	for i := 0; i < len(aeropuertosOrigen); i++ {
		desde := aeropuertosOrigen[i]
		padres, dist := grafo.CaminosMinimos(grafoVuelos, desde)
		DictPadres.Guardar(desde, padres)
		for j := 0; j < len(aeropuertosDestino); j++ {
			hasta := aeropuertosDestino[j]
			tupla := crearTupla(dist.Obtener(hasta), desde, hasta)
			distanciasTotales = append(distanciasTotales, tupla)
		}
	}
	viajeInicial, viajeFinal := obtenerViajeMin(distanciasTotales)
	_, pilaMinima := MostrarViaje(viajeInicial, viajeFinal, DictPadres.Obtener(viajeInicial), imprimir)
	return pilaMinima
}

func MostrarViaje(viajeInicial, viajeFinal string, DictPadres Dict.Diccionario[string, string],
	imprimir bool) (string, Pila.Pila[string]) {

	pilaMinima := Pila.CrearPilaDinamica[string]()
	OrdenFinal := []string{}
	OrdenTxt := ""

	for viajeFinal != viajeInicial {
		OrdenFinal = append(OrdenFinal, viajeFinal)
		pilaMinima.Apilar(viajeFinal)
		viajeFinal = DictPadres.Obtener(viajeFinal)
	}
	OrdenFinal = append(OrdenFinal, viajeFinal)
	pilaMinima.Apilar(viajeFinal)

	for i := len(OrdenFinal) - 1; i > 0; i-- {
		OrdenTxt += OrdenFinal[i] + " -> "
	}
	OrdenTxt += OrdenFinal[PRIMERO]

	if imprimir {
		fmt.Fprintf(os.Stdout, "%s\n", OrdenTxt)
	}
	return OrdenTxt, pilaMinima
}

func MenorEscalas(grafoVuelos grafo.Grafo[string, int], dictCiudades Dict.Diccionario[string, []string],
	origen, destino string) Pila.Pila[string] {

	aeropuertosOrigen := dictCiudades.Obtener(origen)
	aeropuertosDestino := dictCiudades.Obtener(destino)
	DictPadres := Dict.CrearHash[string, Dict.Diccionario[string, string]]()
	distanciasTotales := []tupla{}
	for i := 0; i < len(aeropuertosOrigen); i++ {
		desde := aeropuertosOrigen[i]
		padres := grafo.BFS(grafoVuelos, desde)
		DictPadres.Guardar(desde, padres)
		for j := 0; j < len(aeropuertosDestino); j++ {
			hasta := aeropuertosDestino[j]
			tupla := crearTupla(obtenerRadio(padres, desde, hasta), desde, hasta)
			distanciasTotales = append(distanciasTotales, tupla)
		}
	}
	viajeInicial, viajeFinal := obtenerViajeMin(distanciasTotales)
	_, pilaMinima := MostrarViaje(viajeInicial, viajeFinal, DictPadres.Obtener(viajeInicial), true)
	return pilaMinima
}

func obtenerRadio(padres Dict.Diccionario[string, string], desde, hasta string) int {
	radio := 0
	actual := hasta
	for actual != desde {
		radio++
		actual = padres.Obtener(actual)
	}
	return radio
}

func Centralidad(grafoVuelos grafo.Grafo[string, float64], cantVuelos int) {
	centr := grafo.Centralidad(grafoVuelos)
	heap := Heap.CrearHeap(func(a, b tupla) int { return a.numCmp - b.numCmp })
	arrFinal := []string{}
	vertices := grafoVuelos.ObtenerVertices()
	for _, v := range vertices {
		centrVertice := centr.Obtener(v)
		tupla := crearTupla(centrVertice, v, "")
		heap.Encolar(tupla)
	}
	for i := 0; i < cantVuelos; i++ {
		arrFinal = append(arrFinal, heap.Desencolar().origen)
	}
	for i := 0; i < len(arrFinal)-1; i++ {
		fmt.Fprintf(os.Stdout, "%s, ", arrFinal[i])
	}
	fmt.Fprintf(os.Stdout, "%s\n", arrFinal[len(arrFinal)-1])
}

func NuevaAerolinea(grafoPrecio, grafoTiempo grafo.Grafo[string, int], grafoCantVuelos grafo.Grafo[string, float64], archivo string) {
	csvFile, _ := os.Create(archivo)
	w := csv.NewWriter(csvFile)
	defer w.Flush()

	mst := grafo.ObtenerMST(grafoPrecio)
	padres := grafo.GeneralizadoDFS(mst)
	for _, vertice := range mst.ObtenerVertices() {
		if padres.Obtener(vertice) == grafo.NONE {
			continue
		}
		valor := strconv.Itoa(mst.AristaPeso(vertice, padres.Obtener(vertice)))
		cantVuelos := strconv.Itoa(int(1 / grafoCantVuelos.AristaPeso(vertice, padres.Obtener(vertice))))
		tiempoPromedio := strconv.Itoa(grafoTiempo.AristaPeso(vertice, padres.Obtener(vertice)))
		linea := []string{vertice, padres.Obtener(vertice), tiempoPromedio, valor, cantVuelos}
		w.Write(linea)
	}
	fmt.Fprintf(os.Stdout, "%s\n", "OK")
}

func ItinerarioCultural(grafoTiempo grafo.Grafo[string, int], DictCiudades Dict.Diccionario[string, []string], archivoURL string) {
	archivo, _ := os.Open(archivoURL)
	defer archivo.Close()

	s := bufio.NewScanner(archivo)
	s.Scan()
	linea := s.Text()
	viajesTotales := strings.Split(linea, ",")

	grafoViajes := grafo.CrearGrafo[string, int](true, false)
	for i := 0; len(viajesTotales) > i; i++ {
		grafoViajes.AgregarVertice(viajesTotales[i])
	}

	for s.Scan() {
		linea = s.Text()
		slice := strings.Split(linea, ",")
		grafoViajes.AgregarArista(slice[PRIORITARIO], slice[MENOS_PRIORITARIO], SIN_PESO)
	}

	recorrido := grafo.Topologico(grafoViajes)
	texto := recorrido[PRIMERO]
	for i := 1; len(recorrido) > i; i++ {
		texto += ", " + recorrido[i]
	}
	fmt.Fprintf(os.Stdout, "%s\n", texto)
	for i := 0; len(recorrido)-1 > i; i++ {
		ObtenerCaminoMinimo(grafoTiempo, DictCiudades, recorrido[i], recorrido[i+1], true)
	}
}

func ExportarKml(pila Pila.Pila[string], dictCoordenadas Dict.Diccionario[string, []string], archivoURL string) {
	kmlFile, _ := os.Create(archivoURL)
	defer kmlFile.Close()

	arrViajes := [][]string{}

	kmlFile.WriteString(
		"<?xml version='1.0' encoding='UTF-8'?>\n" +
			"<kml xmlns='http://earth.google.com/kml/2.1'>\n" +
			"	<Document>\n")

	for !pila.EstaVacia() {
		viajes := pila.Desapilar()
		arrViajes = append(arrViajes, dictCoordenadas.Obtener(viajes))
		kmlFile.WriteString(
			"		<Placemark>\n" +
				"			<name>" + viajes + "</name>\n" +
				"			<Point>\n" +
				"				<coordinates>" + dictCoordenadas.Obtener(viajes)[LONGITUD] + ", " + dictCoordenadas.Obtener(viajes)[LATITUD] + "</coordinates>\n" +
				"			</Point>\n" +
				"		</Placemark>\n")

	}

	for i := 0; len(arrViajes)-1 > i; i++ {
		viajeCoord := arrViajes[i][LONGITUD] + ", " + arrViajes[i][LATITUD] + " " + arrViajes[i+1][LONGITUD] + ", " + arrViajes[i+1][LATITUD]

		kmlFile.WriteString(
			"		<Placemark>\n" +
				"			<LineString>\n" +
				"				<coordinates>" + viajeCoord + "</coordinates>\n" +
				"			</LineString>\n" +
				"		</Placemark>\n")
	}

	kmlFile.WriteString(
		"	</Document>\n" +
			"</kml>")

	fmt.Fprintf(os.Stdout, "%s\n", "OK")
}
