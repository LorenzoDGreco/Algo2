package grafo

type Grafo[K comparable, T any] interface {
	AgregarVertice(K)

	EliminarVertice(K)

	HayArista(K, K) bool

	AgregarArista(K, K, T)

	EliminarArista(K, K) T

	AristaPeso(K, K) T

	ObtenerVertices() []K

	ObtenerAdyacentes(K) []K

	ObtenerVerticeAleatorio() K

	EsDirigido() bool

	Pesado() bool
}
