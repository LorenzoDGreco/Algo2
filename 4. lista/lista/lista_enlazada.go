package lista

type listaEnlazada[T any] struct {
	primero *nodo[T]
	ultimo  *nodo[T]
	largo   int
}

type nodo[T any] struct {
	dato      T
	siguiente *nodo[T]
}

type iterador[T any] struct {
	lista    *listaEnlazada[T]
	actual   *nodo[T]
	anterior *nodo[T]
}

func CrearListaEnlazada[T any]() Lista[T] {
	lista := new(listaEnlazada[T])
	return lista
}

func crearNodo[T any](dato T) *nodo[T] {
	nodo := new(nodo[T])
	nodo.dato = dato

	return nodo
}

func (l *listaEnlazada[T]) EstaVacia() bool {
	return l.largo == 0
}

func (l *listaEnlazada[T]) VerPrimero() T {
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}
	return l.primero.dato
}

func (l *listaEnlazada[T]) VerUltimo() T {
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}
	return l.ultimo.dato
}

func (l *listaEnlazada[T]) InsertarPrimero(dato T) {
	nodo := crearNodo(dato)
	if l.EstaVacia() {
		l.ultimo = nodo
	} else {
		nodo.siguiente = l.primero
	}
	l.primero = nodo
	l.largo++
}

func (l *listaEnlazada[T]) InsertarUltimo(dato T) {
	nodo := crearNodo(dato)
	if l.EstaVacia() {
		l.primero = nodo
	}
	if l.ultimo != nil {
		l.ultimo.siguiente = nodo
	}
	l.ultimo = nodo
	l.largo++
}

func (l *listaEnlazada[T]) Largo() int {
	return l.largo
}

func (l *listaEnlazada[T]) BorrarPrimero() T {
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}
	if l.largo == 1 {
		l.ultimo = nil
	}
	datoDeslistado := l.primero.dato
	l.primero = l.primero.siguiente
	l.largo--
	return datoDeslistado
}

func (l *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	actual := l.primero
	continuar := true
	contador := 0
	for continuar && contador < l.largo {
		continuar = visitar(actual.dato)
		actual = actual.siguiente
		contador++
	}
}

func (l *listaEnlazada[T]) Iterador() IteradorLista[T] {
	iterador := new(iterador[T])
	iterador.lista = l
	iterador.actual = l.primero
	iterador.anterior = l.primero
	return iterador
}

func (i *iterador[T]) HaySiguiente() bool {
	return i.actual != nil
}

func (i *iterador[T]) VerActual() T {
	if !i.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return i.actual.dato
}

func (i *iterador[T]) Siguiente() {
	if !i.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	i.anterior = i.actual
	i.actual = i.actual.siguiente
}

func (i *iterador[T]) Insertar(dato T) {
	nuevoNodo := crearNodo(dato)
	if i.lista.largo == 0 {
		i.lista.primero, i.lista.ultimo, i.anterior = nuevoNodo, nuevoNodo, nuevoNodo
	} else {
		if !i.HaySiguiente() {
			i.anterior.siguiente = nuevoNodo
			i.lista.ultimo = nuevoNodo
		} else {
			nuevoNodo.siguiente = i.actual
			if i.anterior != i.actual {
				i.anterior.siguiente = nuevoNodo
			} else {
				i.lista.primero = nuevoNodo
				i.anterior = nuevoNodo
			}
		}
	}
	i.actual = nuevoNodo
	i.lista.largo++
}

func (i *iterador[T]) Borrar() T {
	if !i.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	dato := i.actual.dato
	if i.actual == i.lista.primero {
		i.lista.primero = i.actual.siguiente
	}
	if i.actual == i.lista.ultimo {
		i.lista.ultimo = i.anterior
	}
	i.anterior.siguiente = i.actual.siguiente
	i.actual = i.actual.siguiente
	i.lista.largo--
	return dato
}
