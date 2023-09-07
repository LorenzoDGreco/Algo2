package procesamientos

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	T "Algueiza/Tablero"
	V "Algueiza/Vuelo"
	err "Algueiza/errores"

	TDAHeap "tdas/cola_prioridad"
	TDADict "tdas/diccionario"
)

const (
	NUM_VUELO      = 0
	PRIORIDAD      = 5
	FECHA_DESPEGUE = 6
	DELAY          = 7
	TODO_OK        = "OK"
)

type procesamientos struct {
	tablero   T.Tablero
	DicVuelos TDADict.Diccionario[string, V.Vuelo]
}

func CrearProcesamientos() *procesamientos {
	procesa := new(procesamientos)

	procesa.DicVuelos = TDADict.CrearHash[string, V.Vuelo]()
	procesa.tablero = T.CrearTablero()
	return procesa
}

func (p *procesamientos) ProcesarArchivoVuelos(URLArchivo string) string {
	archivo, errorArchivo := os.Open(URLArchivo)

	errores := err.ErrorURLArchivo{ErrorArchivo: errorArchivo}.Error()
	if errores != TODO_OK {
		return errores
	}

	s := bufio.NewScanner(archivo)
	for s.Scan() {
		linea := s.Text()
		slice := strings.Split(linea, ",")
		dictVuelo := TDADict.CrearHash[int, string]()

		slice[PRIORIDAD] = borrarCeros(slice[PRIORIDAD])
		slice[DELAY] = borrarCeros(slice[DELAY])

		for i := 0; i < 10; i++ {
			dictVuelo.Guardar(i, slice[i])
		}

		Vuelo := V.CrearVuelo(dictVuelo)

		if p.DicVuelos.Pertenece(Vuelo.VerNum()) {
			p.tablero.BorrarUnitario(p.DicVuelos.Obtener(Vuelo.VerNum()))
			p.DicVuelos.Borrar(Vuelo.VerNum())
		}
		p.tablero.GuardarVuelo(slice[FECHA_DESPEGUE], Vuelo)
		p.DicVuelos.Guardar(slice[NUM_VUELO], Vuelo)
	}
	archivo.Close()
	return TODO_OK
}

func borrarCeros(numero string) string {
	if numero == "" || numero == "0" {
		return "0"
	}

	esNegativo := false
	if strings.HasPrefix(numero, "-") { //https://pkg.go.dev/strings#HasPrefix
		esNegativo = true
		numero = strings.TrimPrefix(numero, "-") //https://pkg.go.dev/strings#TrimPrefix
	}

	for i, digito := range numero {
		if string(digito) != "0" {
			if esNegativo {
				return "-" + numero[i:]
			}
			return numero[i:]
		}
	}
	//Si se encuentra una cadena con 2 o más 0 retornará un único 0
	return "0"
}

func (p procesamientos) VerVuelo(numVuelo string) (string, string) {
	errores := err.ErrorNoHayVuelo{HayVuelo: p.DicVuelos.Pertenece(numVuelo)}.Error()
	if errores != TODO_OK {
		return "", errores
	}
	return p.DicVuelos.Obtener(numVuelo).ImprimirVuelo(), TODO_OK
}

func (p procesamientos) ImprimirTablero(cantVuelos, modo, desde, hasta string) string {
	return p.tablero.VerTablero(cantVuelos, modo, desde, hasta)
}

func iterarDiccionario(dict TDADict.Diccionario[string, V.Vuelo]) []V.Vuelo {
	IterDiccionario := dict.Iterador()
	ArrAux := []V.Vuelo{}
	for IterDiccionario.HaySiguiente() {
		_, vuelo := IterDiccionario.VerActual()
		ArrAux = append(ArrAux, vuelo)
		IterDiccionario.Siguiente()
	}
	return ArrAux
}

func (p procesamientos) PrioridadVuelos(K string) string {
	Kint, _ := strconv.Atoi(K)
	errores := err.ErrorCantidadInvalidaPrioridad{Cantidad: Kint}.Error()
	if errores != TODO_OK {
		return errores
	}
	ArrAux := iterarDiccionario(p.DicVuelos)
	ArrHeap := TDAHeap.CrearHeapArr(ArrAux, func(a, b V.Vuelo) int { return a.VerPrioridad() - b.VerPrioridad() })
	contador := 0

	return imprimirVuelosPrioridad(ArrAux, ArrHeap, contador, Kint)
}

func imprimirVuelosPrioridad(ArrVuelos []V.Vuelo, ArrHeap TDAHeap.ColaPrioridad[V.Vuelo], contador, cantVuelos int) string {
	for !ArrHeap.EstaVacia() && contador < cantVuelos {
		vuelo := ArrHeap.Desencolar()
		heapAux := TDAHeap.CrearHeap(func(a, b V.Vuelo) int { return -1 * strings.Compare(a.VerNum(), b.VerNum()) })
		heapAux.Encolar(vuelo)
		if ArrHeap.EstaVacia() {
			fmt.Fprintf(os.Stdout, "%s\n", heapAux.Desencolar().ImprimirPrioridad())
			return TODO_OK
		}
		if heapAux.VerMax().VerPrioridad() == ArrHeap.VerMax().VerPrioridad() {
			for heapAux.VerMax().VerPrioridad() == ArrHeap.VerMax().VerPrioridad() && !ArrHeap.EstaVacia() {
				vuelo = ArrHeap.Desencolar()
				heapAux.Encolar(vuelo)
			}
		}
		for !heapAux.EstaVacia() && contador < cantVuelos {
			fmt.Fprintf(os.Stdout, "%s\n", heapAux.Desencolar().ImprimirPrioridad())
			contador++
		}
	}
	return TODO_OK
}

func (p *procesamientos) SiguienteVuelo(inicio, destino, fecha string) (string, string) {
	IniDest := inicio + " " + destino
	return p.tablero.SiguienteVuelo(IniDest, fecha)
}

func (p *procesamientos) Borrar(desde, hasta string) string {
	errores := err.ErrorRangoFechasInvalida{Desde: desde, Hasta: hasta}.Error()
	if errores != TODO_OK {
		return errores
	}

	VuelosBorrados := p.tablero.Borrar(desde, hasta)
	for i := 0; i < len(VuelosBorrados); i++ {
		numVuelo := VuelosBorrados[i].VerNum()
		vueloImprimir := p.DicVuelos.Borrar(numVuelo)
		fmt.Fprintf(os.Stdout, "%s\n", vueloImprimir.ImprimirVuelo())
	}
	return TODO_OK
}
