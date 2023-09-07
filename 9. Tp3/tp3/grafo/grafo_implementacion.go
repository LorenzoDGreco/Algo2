package grafo

import (
	Dict "tdas/diccionario"
)

type grafos[K comparable, T any] struct {
	vertices Dict.Diccionario[K, Dict.Diccionario[K, T]]
	dirigido bool
	peso     bool
}

func CrearGrafo[K comparable, T any](esDirigido, aristaPeso bool) Grafo[K, T] {
	grafo := new(grafos[K, T])
	grafo.vertices = Dict.CrearHash[K, Dict.Diccionario[K, T]]()
	grafo.dirigido = esDirigido
	grafo.peso = aristaPeso
	return grafo
}

func (g grafos[K, T]) EsDirigido() bool {
	return g.EsDirigido()
}

func (g grafos[K, T]) Pesado() bool {
	return g.peso
}

func (g *grafos[K, T]) AgregarVertice(vertice K) {
	dictVertices := Dict.CrearHash[K, T]()
	g.vertices.Guardar(vertice, dictVertices)
}

func (g *grafos[K, T]) EliminarVertice(vertice K) {
	if g.vertices.Pertenece(vertice) {
		iterador := g.vertices.Iterador()
		for iterador.HaySiguiente() {
			_, dictaux := iterador.VerActual()
			if dictaux.Pertenece(vertice) {
				dictaux.Borrar(vertice)
			}
			iterador.Siguiente()
		}
		g.vertices.Borrar(vertice)
	}
}

func (g *grafos[K, T]) AgregarArista(inicial, final K, peso T) {
	dictAdyacentes := g.vertices.Obtener(inicial)
	dictAdyacentes.Guardar(final, peso)
	if !g.dirigido {
		dictAdyacentes2 := g.vertices.Obtener(final)
		dictAdyacentes2.Guardar(inicial, peso)
	}
}

func (g grafos[K, T]) HayArista(inicial, final K) bool {
	dictAdyacentes := g.vertices.Obtener(inicial)
	return dictAdyacentes.Pertenece(final)
}

func (g *grafos[K, T]) EliminarArista(inicial, final K) T {
	if !g.HayArista(inicial, final) {
		panic("No existe arista entre estos 2 vertices")
	}
	dictAdyacentes := g.vertices.Obtener(inicial)
	peso := dictAdyacentes.Borrar(final)
	if !g.dirigido {
		dictAdyacentes2 := g.vertices.Obtener(final)
		dictAdyacentes2.Borrar(inicial)
	}
	return peso
}

func (g grafos[K, T]) AristaPeso(inicial, final K) T {
	if !g.HayArista(inicial, final) {
		panic("No existe arista entre estos 2 vertices")
	}
	dictAdyacentes := g.vertices.Obtener(inicial)
	return dictAdyacentes.Obtener(final)
}

func (g grafos[K, T]) ObtenerVertices() []K {
	iterDiccionario := g.vertices.Iterador()
	arrVertices := []K{}
	for iterDiccionario.HaySiguiente() {
		vertice, _ := iterDiccionario.VerActual()
		arrVertices = append(arrVertices, vertice)
		iterDiccionario.Siguiente()
	}
	return arrVertices
}

func (g grafos[K, T]) ObtenerAdyacentes(vertice K) []K {
	iterAdyacentes := g.vertices.Obtener(vertice).Iterador()
	arrVerticesAdyacentes := []K{}
	for iterAdyacentes.HaySiguiente() {
		vertice, _ := iterAdyacentes.VerActual()
		arrVerticesAdyacentes = append(arrVerticesAdyacentes, vertice)
		iterAdyacentes.Siguiente()
	}
	return arrVerticesAdyacentes
}

func (g grafos[K, T]) ObtenerVerticeAleatorio() K {
	iterDiccionario := g.vertices.Iterador()
	vertice_aleatorio, _ := iterDiccionario.VerActual()
	return vertice_aleatorio
}
