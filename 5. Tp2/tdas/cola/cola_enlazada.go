package cola

type nodo[T any] struct {
	dato      T
	siguiente *nodo[T]
}

type colaEnlazada[T any] struct {
	primero *nodo[T]
	ultimo  *nodo[T]
}

func CrearColaEnlazada[T any]() Cola[T] {
	cola := new(colaEnlazada[T])

	return cola
}

func crearNodo[T any](dato T) *nodo[T] {
	nodo := new(nodo[T])
	nodo.dato = dato

	return nodo
}

func (c *colaEnlazada[T]) EstaVacia() bool {
	return c.primero == nil
}

func (c *colaEnlazada[T]) VerPrimero() T {
	if c.EstaVacia() {
		panic("La cola esta vacia")
	}
	return c.primero.dato
}

func (c *colaEnlazada[T]) Encolar(dato T) {
	nodo_nuevo := crearNodo(dato)
	if c.EstaVacia() {
		c.primero = nodo_nuevo
		c.ultimo = c.primero
	} else {
		c.ultimo.siguiente = nodo_nuevo
		c.ultimo = c.ultimo.siguiente
	}
}

func (c *colaEnlazada[T]) Desencolar() T {
	datoDesencolado := c.VerPrimero()
	c.primero = c.primero.siguiente
	return datoDesencolado
}
