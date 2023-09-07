package grafo

import (
	TDACola "tdas/cola"
	TDAHeap "tdas/cola_prioridad"
	Dict "tdas/diccionario"
)

const (
	NONE         = ""
	INF          = 999999999
	INF_FLOAT    = 1.0
	VERT_INICIAL = 0
	INICIO       = 0
	DESTINO      = 1
)

type tupla[T any] struct {
	clave     T
	distancia int
}

type tupla64[T any] struct {
	clave T
	peso  float64
}

type valorDefault[T any] struct {
	valor T
}

func valorBase[T any]() T {
	base := new(valorDefault[T])
	return base.valor
}

func crearTupla64[T any](clave T, dist float64) tupla64[T] {
	tupla64 := new(tupla64[T])
	tupla64.clave = clave
	tupla64.peso = dist
	return *tupla64
}

func crearTupla[T any](clave T, dist int) tupla[T] {
	tupla := new(tupla[T])
	tupla.clave = clave
	tupla.distancia = dist
	return *tupla
}

func GeneralizadoDFS[K comparable, T any](grafo Grafo[K, T]) Dict.Diccionario[K, K] {
	padres := Dict.CrearHash[K, K]()
	visitados := Dict.CrearHash[K, T]()
	vertices := grafo.ObtenerVertices()
	for _, vertice := range vertices {
		if !visitados.Pertenece(vertice) {
			visitados.Guardar(vertice, valorBase[T]())
			padres.Guardar(vertice, valorBase[K]())
			padres, visitados = DFS(grafo, padres, visitados, vertice)
		}
	}
	return padres
}

func DFS[K comparable, T any](grafo Grafo[K, T], padres Dict.Diccionario[K, K], visitados Dict.Diccionario[K, T], Vertice K) (Dict.Diccionario[K, K], Dict.Diccionario[K, T]) {
	for _, adyacentes := range grafo.ObtenerAdyacentes(Vertice) {
		if !visitados.Pertenece(adyacentes) {
			padres.Guardar(adyacentes, Vertice)
			visitados.Guardar(adyacentes, valorBase[T]())
			padres, visitados = DFS(grafo, padres, visitados, adyacentes)
		}
	}
	return padres, visitados
}

func BFS[K comparable, T any](grafo Grafo[K, T], vertice_inicial K) Dict.Diccionario[K, K] {
	padres := Dict.CrearHash[K, K]()
	visitados := Dict.CrearHash[K, T]()
	cola := TDACola.CrearColaEnlazada[K]()
	cola.Encolar(vertice_inicial)
	padres.Guardar(vertice_inicial, valorBase[K]())
	visitados.Guardar(vertice_inicial, valorBase[T]())
	for !cola.EstaVacia() {
		v := cola.Desencolar()
		adyacentes := grafo.ObtenerAdyacentes(v)
		for i := 0; i < len(adyacentes); i++ {
			if !visitados.Pertenece(adyacentes[i]) {
				visitados.Guardar(adyacentes[i], valorBase[T]())
				padres.Guardar(adyacentes[i], v)
				cola.Encolar(adyacentes[i])
			}
		}
	}
	return padres
}

func dijkstra[K comparable](grafo Grafo[K, int], vertice_inicial K) (Dict.Diccionario[K, K], Dict.Diccionario[K, int]) {
	padres := Dict.CrearHash[K, K]()
	dist := Dict.CrearHash[K, int]()
	vertices := grafo.ObtenerVertices()
	Heap := TDAHeap.CrearHeap(func(a, b tupla[K]) int { return -1 * (a.distancia - b.distancia) })
	for i := 0; i < len(vertices); i++ {
		dist.Guardar(vertices[i], INF)
	}
	tupla := crearTupla(vertice_inicial, 0)
	Heap.Encolar(tupla)
	padres.Guardar(vertice_inicial, valorBase[K]())
	dist.Guardar(vertice_inicial, 0)
	for !Heap.EstaVacia() {
		v := Heap.Desencolar()
		adyacentes := grafo.ObtenerAdyacentes(v.clave)
		for i := 0; i < len(adyacentes); i++ {
			if dist.Obtener(v.clave)+grafo.AristaPeso(v.clave, adyacentes[i]) < dist.Obtener(adyacentes[i]) {
				dist.Guardar(adyacentes[i], dist.Obtener(v.clave)+grafo.AristaPeso(v.clave, adyacentes[i]))
				padres.Guardar(adyacentes[i], v.clave)
				tupla := crearTupla(adyacentes[i], dist.Obtener(adyacentes[i]))
				Heap.Encolar(tupla)
			}
		}
	}
	return padres, dist
}

func dijkstraFloat64[K comparable](grafo Grafo[K, float64], vertice_inicial K) (Dict.Diccionario[K, K], Dict.Diccionario[K, float64]) {
	padres := Dict.CrearHash[K, K]()
	dist := Dict.CrearHash[K, float64]()
	vertices := grafo.ObtenerVertices()
	Heap := TDAHeap.CrearHeap(func(a, b tupla64[K]) int {
		if a.peso-b.peso < 0 {
			return 1
		} else if a.peso-b.peso > 0 {
			return -1
		} else {
			return 0
		}
	})

	dist = guardarVerticesDict(dist, vertices, INF_FLOAT)

	tupla := crearTupla64(vertice_inicial, 0.0)
	Heap.Encolar(tupla)
	padres.Guardar(vertice_inicial, valorBase[K]())
	dist.Guardar(vertice_inicial, 0)
	for !Heap.EstaVacia() {
		v := Heap.Desencolar()
		adyacentes := grafo.ObtenerAdyacentes(v.clave)
		for _, w := range adyacentes {
			if dist.Obtener(v.clave)+grafo.AristaPeso(v.clave, w) < dist.Obtener(w) {
				dist.Guardar(w, dist.Obtener(v.clave)+grafo.AristaPeso(v.clave, w))
				padres.Guardar(w, v.clave)
				tupla := crearTupla64(w, dist.Obtener(w))
				Heap.Encolar(tupla)
			}
		}
	}
	return padres, dist
}

func CaminosMinimos[K comparable](grafo Grafo[K, int], vertice_inicial K) (Dict.Diccionario[K, K], Dict.Diccionario[K, int]) {
	if grafo.Pesado() {
		padres, distancia := dijkstra(grafo, vertice_inicial)
		return padres, distancia
	}
	padres := BFS(grafo, vertice_inicial)
	return padres, nil
}

func Centralidad[K comparable](grafo Grafo[K, float64]) Dict.Diccionario[K, int] {
	centr := Dict.CrearHash[K, int]()
	vertices := grafo.ObtenerVertices()
	centr = guardarVerticesDict(centr, vertices, 0)

	for _, v := range vertices {
		centrAux := Dict.CrearHash[K, int]()

		padres, dist := dijkstraFloat64(grafo, v)
		centrAux = guardarVerticesDict(centrAux, vertices, 0)
		VerticesOrdenados := ordenarVertices(grafo, dist)

		for _, w := range VerticesOrdenados {
			if padres.Obtener(w) == valorBase[K]() {
				continue
			}
			centrAux.Guardar(padres.Obtener(w), 1+centrAux.Obtener(w)+centrAux.Obtener(padres.Obtener(w)))
		}
		for _, w := range vertices {
			if w == v {
				continue
			}
			centr.Guardar(w, centr.Obtener(w)+centrAux.Obtener(w))
		}
	}
	return centr
}

func ordenarVertices[K comparable](grafo Grafo[K, float64], dist Dict.Diccionario[K, float64]) []K {
	vertices := grafo.ObtenerVertices()
	arrOrdenado := []K{}
	Heap := TDAHeap.CrearHeap(func(a, b tupla64[K]) int {
		if a.peso-b.peso > 0 {
			return 1
		} else if a.peso-b.peso < 0 {
			return -1
		} else {
			return 0
		}
	})
	for i := 0; i < len(vertices); i++ {
		tupla := crearTupla64(vertices[i], dist.Obtener(vertices[i]))
		Heap.Encolar(tupla)
	}
	for !Heap.EstaVacia() {
		arrOrdenado = append(arrOrdenado, Heap.Desencolar().clave)
	}
	return arrOrdenado
}

func ObtenerMST[K comparable](grafo Grafo[K, int]) Grafo[K, int] {
	visitados := Dict.CrearHash[K, int]()
	vertice := grafo.ObtenerVerticeAleatorio()
	visitados.Guardar(vertice, 0)

	arbol := crearArbol(grafo)
	heap := crearHeap(grafo, vertice)

	for !heap.EstaVacia() {
		tupla := heap.Desencolar()
		viaje, peso := tupla.clave, tupla.distancia
		if visitados.Pertenece(viaje[DESTINO]) {
			continue
		}
		arbol.AgregarArista(viaje[INICIO], viaje[DESTINO], peso)
		visitados.Guardar(viaje[DESTINO], 0)
		adyacentes := grafo.ObtenerAdyacentes(viaje[DESTINO])
		for i := 0; i < len(adyacentes); i++ {
			arr := []K{viaje[DESTINO], adyacentes[i]}
			tupla := crearTupla(arr, grafo.AristaPeso(viaje[DESTINO], adyacentes[i]))
			heap.Encolar(tupla)
		}
	}
	return arbol
}

func crearArbol[K comparable](grafo Grafo[K, int]) Grafo[K, int] {
	arbol := CrearGrafo[K, int](false, true)
	vertices := grafo.ObtenerVertices()

	for i := 0; i < len(vertices); i++ {
		arbol.AgregarVertice(vertices[i])
	}
	return arbol
}

func crearHeap[K comparable](grafo Grafo[K, int], vertice K) TDAHeap.ColaPrioridad[tupla[[]K]] {
	heap := TDAHeap.CrearHeap(func(a, b tupla[[]K]) int { return -1 * (a.distancia - b.distancia) })
	adyacentes := grafo.ObtenerAdyacentes(vertice)

	for i := 0; i < len(adyacentes); i++ {
		arr := []K{vertice, adyacentes[i]}
		tupla := crearTupla(arr, grafo.AristaPeso(vertice, adyacentes[i]))
		heap.Encolar(tupla)
	}
	return heap
}

func Topologico[K comparable](grafo Grafo[K, int]) []K {
	vertices := grafo.ObtenerVertices()
	resultado := []K{}

	grados := crearGrados(grafo, vertices)

	cola := encolarPrimerosVertices(grados, vertices)

	for !cola.EstaVacia() {
		vertice := cola.Desencolar()
		resultado = append(resultado, vertice)
		adyacentes := grafo.ObtenerAdyacentes(vertice)
		for i := 0; i < len(adyacentes); i++ {
			grados.Guardar(adyacentes[i], grados.Obtener(adyacentes[i])-1)
			if grados.Obtener(adyacentes[i]) == 0 {
				cola.Encolar(adyacentes[i])
			}
		}
	}
	return resultado
}

func crearGrados[K comparable](grafo Grafo[K, int], vertices []K) Dict.Diccionario[K, int] {
	grados := Dict.CrearHash[K, int]()

	grados = guardarVerticesDict(grados, vertices, 0)

	for i := 0; i < len(vertices); i++ {
		adyacentes := grafo.ObtenerAdyacentes(vertices[i])
		for i := 0; i < len(adyacentes); i++ {
			grados.Guardar(adyacentes[i], grados.Obtener(adyacentes[i])+1)
		}
	}
	return grados
}

func encolarPrimerosVertices[K comparable](grados Dict.Diccionario[K, int], vertices []K) TDACola.Cola[K] {
	cola := TDACola.CrearColaEnlazada[K]()
	for i := 0; i < len(vertices); i++ {
		if grados.Obtener(vertices[i]) == 0 {
			cola.Encolar(vertices[i])
		}
	}
	return cola
}

func guardarVerticesDict[K comparable, T any](dict Dict.Diccionario[K, T], vertices []K, peso T) Dict.Diccionario[K, T] {
	for i := 0; i < len(vertices); i++ {
		dict.Guardar(vertices[i], peso)
	}
	return dict
}
