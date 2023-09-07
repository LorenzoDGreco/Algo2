package diccionario

import (
	"fmt"
	TDALista "tdas/lista"
)

/*
De aca sacamos la funcion de hashing
https://pkg.go.dev/hash/crc64#Table
https://cs.opensource.google/go/go/+/refs/tags/go1.20.4:src/hash/crc64/crc64.go
*/

const (
	CAPACIDAD_INICIAL int     = 10
	EXTENSION         int     = 2
	INDICE_EXTENSION  float32 = 3
	INDICE_REDUCCION  float32 = 0.5
	NO_HAY_ELEMENTOS  int     = 0
)

func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

type claveValor[K comparable, T any] struct {
	clave K
	valor T
}

type hashStruct[K comparable, T any] struct {
	arreglo  []TDALista.Lista[claveValor[K, T]]
	tabla    *Table
	cantidad int
}

type iterador[K comparable, T any] struct {
	dict          *hashStruct[K, T]
	iteradorLista TDALista.IteradorLista[claveValor[K, T]]
	posicion      int
	contador      int
}

func crearClaveValor[K comparable, T any](clave K, valor T) claveValor[K, T] {
	claveValor := new(claveValor[K, T])
	claveValor.clave = clave
	claveValor.valor = valor
	return *claveValor
}

func (h *hashStruct[K, T]) redimensionar() {
	if float32(h.cantidad/cap(h.arreglo)) > INDICE_EXTENSION {
		arregloNuevo := make([]TDALista.Lista[claveValor[K, T]], cap(h.arreglo)*EXTENSION)
		h.nuevoHash(arregloNuevo)

	} else if h.cantidad < cap(h.arreglo)/4 && cap(h.arreglo) > CAPACIDAD_INICIAL {
		var arregloNuevo []TDALista.Lista[claveValor[K, T]]

		if h.cantidad > CAPACIDAD_INICIAL {
			arregloNuevo = make([]TDALista.Lista[claveValor[K, T]], cap(h.arreglo)/3)
		} else {
			arregloNuevo = make([]TDALista.Lista[claveValor[K, T]], CAPACIDAD_INICIAL)
		}

		h.nuevoHash(arregloNuevo)

	}
}

func (h *hashStruct[K, T]) nuevoHash(arregloNuevo []TDALista.Lista[claveValor[K, T]]) {
	for i := 0; i < cap(h.arreglo); i++ {
		if h.arreglo[i] != nil {
			lista := h.arreglo[i]
			if !lista.EstaVacia() {
				iterador := lista.Iterador()
				for iterador.HaySiguiente() {
					posicionNueva := obtenerPosicion(iterador.VerActual().clave, arregloNuevo, h.tabla)
					if arregloNuevo[posicionNueva] == nil {
						arregloNuevo[posicionNueva] = TDALista.CrearListaEnlazada[claveValor[K, T]]()
					}
					arregloNuevo[posicionNueva].InsertarUltimo(iterador.VerActual())
					iterador.Siguiente()
				}
			}
		}
	}
	h.arreglo = arregloNuevo
}

func CrearHash[K comparable, T any]() Diccionario[K, T] {
	hash := new(hashStruct[K, T])
	hash.arreglo = make([]TDALista.Lista[claveValor[K, T]], CAPACIDAD_INICIAL)
	hash.tabla = makeTable(ECMA)
	return hash
}

func obtenerPosicion[K comparable, T any](clave K, arreglo []TDALista.Lista[claveValor[K, T]], tabla *Table) int {
	Bytes := convertirABytes(clave)
	var posicion int
	if cap(arreglo) != 0 {
		posicion = int(Checksum(Bytes, tabla) % uint64(cap(arreglo)))
	}
	return posicion
}

func iterarLista[K comparable, T any](clave K, lista TDALista.Lista[claveValor[K, T]]) (TDALista.IteradorLista[claveValor[K, T]], bool) {
	iterador := lista.Iterador()
	seEncuentra := false
	for iterador.HaySiguiente() && !seEncuentra {
		if iterador.VerActual().clave == clave {
			seEncuentra = true
		} else {
			iterador.Siguiente()
		}
	}
	return iterador, seEncuentra
}

func (h *hashStruct[K, T]) Guardar(clave K, dato T) {

	claveDato := crearClaveValor(clave, dato)
	posicion := obtenerPosicion(clave, h.arreglo, h.tabla)
	if h.arreglo[posicion] == nil {
		h.arreglo[posicion] = TDALista.CrearListaEnlazada[claveValor[K, T]]()
	}
	iterador, seEncuentra := iterarLista(clave, h.arreglo[posicion])
	if seEncuentra {
		iterador.Borrar()
		h.arreglo[posicion].InsertarPrimero(claveDato)
	} else {
		h.arreglo[posicion].InsertarUltimo(claveDato)
		h.cantidad++
	}
	h.redimensionar()
}

func (h *hashStruct[K, T]) Pertenece(clave K) bool {
	posicion := obtenerPosicion(clave, h.arreglo, h.tabla)
	seEncuentra := false
	if h.arreglo[posicion] == nil {
		return seEncuentra
	}
	_, seEncuentra = iterarLista(clave, h.arreglo[posicion])
	return seEncuentra
}

func (h *hashStruct[K, T]) Obtener(clave K) T {
	posicion := obtenerPosicion(clave, h.arreglo, h.tabla)
	if h.arreglo[posicion] == nil {
		panic("La clave no pertenece al diccionario")
	}
	iterador, seEncuentra := iterarLista(clave, h.arreglo[posicion])
	if !seEncuentra {
		panic("La clave no pertenece al diccionario")
	} else {
		return iterador.VerActual().valor
	}
}

func (h *hashStruct[K, T]) Borrar(clave K) T {
	posicion := obtenerPosicion(clave, h.arreglo, h.tabla)
	var valor T
	if h.arreglo[posicion] != nil {

		iterador, seEncuentra := iterarLista(clave, h.arreglo[posicion])
		if !seEncuentra {
			panic("La clave no pertenece al diccionario")
		}
		valor = iterador.VerActual().valor
		iterador.Borrar()
		h.cantidad--
		h.redimensionar()
	} else {
		panic("La clave no pertenece al diccionario")
	}
	return valor
}

func (h *hashStruct[K, T]) Cantidad() int {
	return h.cantidad
}

func (h *hashStruct[K, T]) Iterar(visitar func(clave K, dato T) bool) {
	actual := h.arreglo
	continuar := true
	var contador int
	for continuar && contador < cap(h.arreglo) {
		if actual[contador] != nil {
			iterador := actual[contador].Iterador()
			for iterador.HaySiguiente() && continuar {
				continuar = visitar(iterador.VerActual().clave, iterador.VerActual().valor)
				iterador.Siguiente()
			}
		}
		contador++
	}
}

func (h *hashStruct[K, T]) Iterador() IterDiccionario[K, T] {
	iterador := new(iterador[K, T])
	iterador.dict = h
	if h.arreglo[iterador.posicion] == nil {
		h.arreglo[iterador.posicion] = TDALista.CrearListaEnlazada[claveValor[K, T]]()
	}
	iterador.iteradorLista = h.arreglo[iterador.posicion].Iterador()
	return iterador
}

func (i *iterador[K, T]) HaySiguiente() bool {
	if i.dict.cantidad == NO_HAY_ELEMENTOS {
		return false
	}
	return !(i.contador == i.dict.cantidad)

}

func (i *iterador[K, T]) VerActual() (K, T) {
	if !i.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	if !i.iteradorLista.HaySiguiente() {
		i.Siguiente()
	}
	dato := i.iteradorLista.VerActual()
	return dato.clave, dato.valor
}

func (i *iterador[K, T]) Siguiente() {
	if !i.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	if i.iteradorLista.HaySiguiente() {
		i.iteradorLista.Siguiente()
		i.contador++
	} else {
		i.avanzarLista()
	}
}

func (i *iterador[K, T]) avanzarLista() {
	i.posicion++
	for (i.dict.arreglo[i.posicion] == nil || i.dict.arreglo[i.posicion].EstaVacia()) && i.posicion != len(i.dict.arreglo) {
		i.posicion++

	}
	i.iteradorLista = i.dict.arreglo[i.posicion].Iterador()

}
