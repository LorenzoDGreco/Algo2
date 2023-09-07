package Tablero

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	V "Algueiza/Vuelo"
	err "Algueiza/errores"
	TDAHeap "tdas/cola_prioridad"
	TDADict "tdas/diccionario"
	TDAPila "tdas/pila"
)

const (
	ASC     = "asc"
	DESC    = "desc"
	TODO_OK = "OK"
)

type Tupla[T any] struct {
	Clave T
	Valor string
}

func CrearTupla[T any](clave T, dato string) *Tupla[T] {
	tupla := new(Tupla[T])
	tupla.Clave = clave
	tupla.Valor = dato
	return tupla
}

type tablero struct {
	abbPrincipal TDADict.DiccionarioOrdenado[string, infoTablero]
}

type infoTablero struct {
	abbInfo    TDADict.DiccionarioOrdenado[string, V.Vuelo]
	dictViajes TDADict.Diccionario[string, TDAHeap.ColaPrioridad[V.Vuelo]]
}

func CrearTablero() Tablero {
	tabl := new(tablero)
	tabl.abbPrincipal = TDADict.CrearABB[string, infoTablero](func(a, b string) int { return strings.Compare(a, b) })
	return tabl
}

func crearInfoTablero() infoTablero {
	infoTablero := new(infoTablero)
	infoTablero.abbInfo = TDADict.CrearABB[string, V.Vuelo](func(a, b string) int { return strings.Compare(a, b) })
	infoTablero.dictViajes = TDADict.CrearHash[string, TDAHeap.ColaPrioridad[V.Vuelo]]()
	return *infoTablero
}

func (t *tablero) GuardarVuelo(fechaDespegue string, vuelo V.Vuelo) {
	if t.abbPrincipal.Pertenece(fechaDespegue) {
		t.actualizarVuelo(fechaDespegue, vuelo)
		return
	}

	infoTablero := crearInfoTablero()
	Heap := TDAHeap.CrearHeap(func(a, b V.Vuelo) int { return strings.Compare(a.FechaVuelo(), b.FechaVuelo()) })

	Heap.Encolar(vuelo)
	infoTablero.dictViajes.Guardar(vuelo.VerViaje(), Heap)
	infoTablero.abbInfo.Guardar(vuelo.VerNum(), vuelo)
	t.abbPrincipal.Guardar(fechaDespegue, infoTablero)
}

func (t *tablero) actualizarVuelo(fechaDespegue string, vuelo V.Vuelo) {
	infoTab := t.abbPrincipal.Obtener(fechaDespegue)
	infoTab.abbInfo.Guardar(vuelo.VerNum(), vuelo)

	if infoTab.dictViajes.Pertenece(vuelo.VerViaje()) {
		heap := infoTab.dictViajes.Obtener(vuelo.VerViaje())
		heap.Encolar(vuelo)
		infoTab.dictViajes.Guardar(vuelo.VerViaje(), heap)
		t.abbPrincipal.Guardar(fechaDespegue, infoTab)
		return
	}

	heap := TDAHeap.CrearHeap(func(a, b V.Vuelo) int { return strings.Compare(a.FechaVuelo(), b.FechaVuelo()) })
	heap.Encolar(vuelo)
	infoTab.dictViajes.Guardar(vuelo.VerViaje(), heap)
	t.abbPrincipal.Guardar(fechaDespegue, infoTab)
}

func (t *tablero) VerTablero(cantVuelos, modo, desde, hasta string) string {
	cantidadVuelos, _ := strconv.Atoi(cantVuelos)
	errores := erroresTablero(cantidadVuelos, modo, desde, hasta)
	if errores != TODO_OK {
		return errores
	}

	Contador := 0
	if modo == ASC {
		t.abbPrincipal.IterarRango(&desde, &hasta, func(a string, b infoTablero) bool {
			vuelosVistos := printTablero(b.abbInfo, modo, hasta, cantidadVuelos)
			sumarContador(&Contador, vuelosVistos)
			return Contador < cantidadVuelos
		})
		return TODO_OK
	}

	pila := TDAPila.CrearPilaDinamica[TDADict.DiccionarioOrdenado[string, V.Vuelo]]()
	t.abbPrincipal.IterarRango(&desde, &hasta, func(a string, b infoTablero) bool {
		apilar(&pila, b.abbInfo)
		return true
	})
	for !pila.EstaVacia() && Contador < cantidadVuelos {
		arbol := pila.Desapilar()
		vuelosVistos := printTablero(arbol, modo, hasta, cantidadVuelos)
		sumarContador(&Contador, vuelosVistos)
	}

	return TODO_OK
}

func erroresTablero(cantVuelos int, modo, desde, hasta string) string {
	errores := err.ErrorCantidadInvalida{Cantidad: cantVuelos}.Error()
	if errores != err.TODO_OK {
		return errores
	}
	errores = err.ErrorFechaInvalida{Desde: desde, Hasta: hasta}.Error()
	if errores != err.TODO_OK {
		return errores
	}
	errores = err.ErrorModoInvalido{Modo: modo}.Error()
	return errores
}

func sumarContador(contador *int, Suma int) {
	*contador = *contador + Suma
}

func apilar(pila *TDAPila.Pila[TDADict.DiccionarioOrdenado[string, V.Vuelo]], subArbol TDADict.DiccionarioOrdenado[string, V.Vuelo]) {
	pilaAux := *pila
	pilaAux.Apilar(subArbol)
}

func printTablero(b TDADict.Diccionario[string, V.Vuelo], modo, fechaMax string, cantVuelos int) int {
	contador := 0
	IteradorArbol := b.Iterador()

	if modo == ASC {
		return printTableroASC(contador, cantVuelos, IteradorArbol)
	}

	return printTableroDES(contador, cantVuelos, IteradorArbol)

}

func printTableroASC(contador, cantVuelos int, IteradorArbol TDADict.IterDiccionario[string, V.Vuelo]) int {
	for IteradorArbol.HaySiguiente() && contador < cantVuelos {
		_, Vuelo := IteradorArbol.VerActual()
		fmt.Fprintf(os.Stdout, "%s\n", Vuelo.ImprimirTablero())
		contador++
		IteradorArbol.Siguiente()
	}
	return contador
}

func printTableroDES(contador, cantVuelos int, IteradorArbol TDADict.IterDiccionario[string, V.Vuelo]) int {
	pila := TDAPila.CrearPilaDinamica[V.Vuelo]()
	for IteradorArbol.HaySiguiente() {
		_, Vuelo := IteradorArbol.VerActual()
		pila.Apilar(Vuelo)
		IteradorArbol.Siguiente()
	}
	for !pila.EstaVacia() && contador < cantVuelos {
		Vuelo := pila.Desapilar()
		fmt.Fprintf(os.Stdout, "%s\n", Vuelo.ImprimirTablero())
		contador++
	}
	return contador
}

func (t *tablero) BorrarUnitario(vuelo V.Vuelo) {
	infoTab := t.abbPrincipal.Obtener(vuelo.FechaVuelo())
	infoTab.eliminarViajeViejo(vuelo)
	infoTab.abbInfo.Borrar(vuelo.VerNum())
	t.abbPrincipal.Guardar(vuelo.FechaVuelo(), infoTab)
}

func (t *infoTablero) eliminarViajeViejo(vuelo V.Vuelo) {

	if !t.dictViajes.Pertenece(vuelo.VerViaje()) {
		return
	}
	heap := t.dictViajes.Obtener(vuelo.VerViaje())
	arrAux := []V.Vuelo{}
	for !heap.EstaVacia() {
		arrAux = append(arrAux, heap.Desencolar())
	}
	for i := 0; i < len(arrAux); i++ {
		if arrAux[i].VerNum() != vuelo.VerNum() {
			heap.Encolar(arrAux[i])
		}
	}
	t.dictViajes.Guardar(vuelo.VerViaje(), heap)
}

func (t *tablero) Borrar(desde, hasta string) []V.Vuelo {
	iterador := t.abbPrincipal.IteradorRango(&desde, &hasta)
	arrClaves := []string{}
	arrVuelos := []V.Vuelo{}

	for iterador.HaySiguiente() {
		clave, infotablero := iterador.VerActual()
		arrClaves = append(arrClaves, clave)
		iteradorAbbInfo := infotablero.abbInfo.Iterador()
		for iteradorAbbInfo.HaySiguiente() {
			_, vuelo := iteradorAbbInfo.VerActual()
			arrVuelos = append(arrVuelos, vuelo)
			iteradorAbbInfo.Siguiente()
		}
		iterador.Siguiente()
	}
	for i := 0; i < len(arrClaves); i++ {
		t.abbPrincipal.Borrar(arrClaves[i])
	}
	return arrVuelos
}

func (t *tablero) SiguienteVuelo(viaje, fecha string) (string, string) {
	iterador := t.abbPrincipal.IteradorRango(&fecha, nil)
	Viaje := strings.Split(viaje, " ")
	if !iterador.HaySiguiente() {
		return err.ErrorNoHayVuelosNuevos{NoHayVuelo: true, Desde: Viaje[0], Hasta: Viaje[1], Fecha: fecha}.Error(), TODO_OK
	}
	for iterador.HaySiguiente() {
		_, infoTablero := iterador.VerActual()
		dicViajes := infoTablero.dictViajes
		if dicViajes.Pertenece(viaje) {
			if !dicViajes.Obtener(viaje).EstaVacia() {
				return dicViajes.Obtener(viaje).VerMax().ImprimirVuelo(), TODO_OK
			}
		}
		iterador.Siguiente()
	}
	return err.ErrorNoHayVuelosNuevos{NoHayVuelo: true, Desde: Viaje[0], Hasta: Viaje[1], Fecha: fecha}.Error(), TODO_OK
}
