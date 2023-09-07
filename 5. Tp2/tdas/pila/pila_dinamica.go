package pila

const vacio = 0
const inicioPila = 10
const doble = 2
const mitad = 2
const tamanioMax = 4

/* Definición del struct pila proporcionado por la cátedra. */
type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func CrearPilaDinamica[T any]() Pila[T] {
	/*
		PRE: Recibe un tipo que quiera armar su pila
		POST: Retorna la pila con sus valores asignados
	*/
	pila := pilaDinamica[T]{}

	pila.datos = make([]T, inicioPila)
	pila.cantidad = vacio

	return &pila
}

func (pila *pilaDinamica[T]) Apilar(elem T) {
	/*
		PRE: Recibe un unico elemento tipo generico
		POST: Lo apila al ultimo lugar de la pila
	*/
	pila.esRedimensionable()
	pila.datos[pila.cantidad] = elem
	pila.cantidad++

}

func (pila *pilaDinamica[T]) Desapilar() T {
	/*
		POST: Desapila el ultimo elemento y lo retorna, si está vacia levanta un panic
	*/
	if pila.EstaVacia() {
		panic("La pila esta vacia")
	}
	pila.cantidad--
	elem := pila.datos[pila.cantidad]
	pila.esRedimensionable()
	return elem
}

func (pila *pilaDinamica[T]) EstaVacia() bool {
	/*
		POST: Si la pila está vacia devuelve True, en caso contrario devuelve False
	*/
	return pila.cantidad == vacio
}

func (pila *pilaDinamica[T]) VerTope() T {
	/*
		POST: Devuelve el ultimo elemento de la pila, si no hay elementos levanta un panic
	*/
	if pila.EstaVacia() {
		panic("La pila esta vacia")
	}
	return pila.datos[pila.cantidad-1]
}

func (pila *pilaDinamica[T]) capacidadMaximo() int {
	/*
		POST: Devuelve la Capacidad máxima de elementos
	*/
	return cap(pila.datos)
}

func (pila *pilaDinamica[T]) esRedimensionable() {
	/*
		POST: Si es redimensionable llama a redimensionarTamanio sino no hace nada
	*/
	if pila.cantidad*tamanioMax <= pila.capacidadMaximo() && pila.capacidadMaximo() > inicioPila {
		pila.redimensionarTamanio((pila.capacidadMaximo() / mitad))
	} else if pila.capacidadMaximo() == pila.cantidad {
		pila.redimensionarTamanio(pila.cantidad * doble)
	}
}

func (pila *pilaDinamica[T]) redimensionarTamanio(tamanio int) { //Desventajas al programar en espaniol F
	/*
		PRE: Recibe un numero entero del tamino al cual modificar
		POST: Devuelve una lista extendida por el tamanio con la informacion copiada del anterior array
	*/
	nuevoArray := make([]T, tamanio)
	copy(nuevoArray, pila.datos)
	pila.datos = nuevoArray
}
