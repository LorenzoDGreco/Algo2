package cola_prioridad

const (
	INICIO_HEAP        = 10
	REDUCCION          = 2
	AMPLIFICACION      = 2
	FACTOR_REDIMENSION = 4
)

type heap[T any] struct {
	arr      []T
	cantidad int
	cmp      func(T, T) int
}

func CrearHeap[T any](FuncCmp func(T, T) int) ColaPrioridad[T] {
	heap := new(heap[T])
	heap.arr = make([]T, INICIO_HEAP)
	heap.cmp = FuncCmp
	return heap
}

func (h *heap[T]) EstaVacia() bool {
	return h.cantidad == 0
}

func (h *heap[T]) Encolar(elem T) {
	h.esRedimensionable()
	h.arr[h.cantidad] = elem
	h.upHeap(h.cantidad)
	h.cantidad++
}

func (h *heap[T]) VerMax() T {
	if h.EstaVacia() {
		panic("La cola esta vacia")
	}
	return h.arr[0]
}

func (h *heap[T]) Desencolar() T {
	elem := h.VerMax()
	h.swap(0, h.cantidad-1)
	h.cantidad--
	h.downHeap(0)
	h.esRedimensionable()
	return elem
}

func (h *heap[T]) Cantidad() int {
	return h.cantidad
}

func (h *heap[T]) upHeap(posHijo int) {
	posPadre := encontrarPosPadre(posHijo)
	if posPadre < 0 {
		return
	}
	comp := h.cmp(h.arr[posPadre], h.arr[posHijo])

	if comp < 0 {
		h.swap(posPadre, posHijo)
		h.upHeap(posPadre)
	}
}

func (h *heap[T]) downHeap(posPadre int) {
	posHijoIzq := encontrarPosHijoIzq(posPadre)
	posHijoDer := encontrarPosHijoDer(posPadre)

	var compIzq, compDer int
	if h.Cantidad() > posHijoIzq {
		compIzq = h.cmp(h.arr[posPadre], h.arr[posHijoIzq])
	}
	if h.Cantidad() > posHijoDer {
		compDer = h.cmp(h.arr[posPadre], h.arr[posHijoDer])
	}

	if compIzq >= 0 && compDer >= 0 {
		return
	} else if compIzq < 0 && posHijoDer >= h.Cantidad() {
		h.swap(posPadre, posHijoIzq)
		h.downHeap(posHijoIzq)
	} else {
		compMax := h.cmp(h.arr[posHijoIzq], h.arr[posHijoDer])

		if compMax >= 0 {
			h.swap(posPadre, posHijoIzq)
			h.downHeap(posHijoIzq)
		} else if compMax < 0 {
			h.swap(posPadre, posHijoDer)
			h.downHeap(posHijoDer)
		}
	}

}

func (h *heap[T]) esRedimensionable() {
	if h.cantidad*FACTOR_REDIMENSION <= h.capacidadMaxima() && h.capacidadMaxima() > INICIO_HEAP {
		h.redimensionarTamanio((h.capacidadMaxima() / REDUCCION))
	} else if h.capacidadMaxima() == h.cantidad {
		if h.cantidad == 0 || h.capacidadMaxima() == 0 {
			h.redimensionarTamanio(INICIO_HEAP)
		} else {
			h.redimensionarTamanio(h.cantidad * AMPLIFICACION)
		}
	}
}

func (h *heap[T]) redimensionarTamanio(tamanio int) {
	nuevoArray := make([]T, tamanio)
	copy(nuevoArray, h.arr)
	h.arr = nuevoArray
}

func (h *heap[T]) capacidadMaxima() int {
	return cap(h.arr)
}

func encontrarPosPadre(posHijo int) int {
	posPadre := (posHijo - 1) / 2
	return posPadre
}

func encontrarPosHijoIzq(posPadre int) int {
	pos := (posPadre * 2) + 1
	return pos
}

func encontrarPosHijoDer(posPadre int) int {
	pos := (posPadre * 2) + 2
	return pos
}

func (h *heap[T]) swap(posPadre, posHijo int) {
	h.arr[posHijo], h.arr[posPadre] = h.arr[posPadre], h.arr[posHijo]
}

func CrearHeapArr[T any](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	heap := new(heap[T])
	heap.cmp = funcion_cmp
	nuevoArray := make([]T, len(arreglo))
	copy(nuevoArray, arreglo)
	heap.arr = nuevoArray
	heap.cantidad = len(arreglo)
	heap.heapify(heap.cantidad)

	return heap
}

func (h *heap[T]) heapify(posicion int) int {
	if posicion < 0 {
		return 0
	}
	h.downHeap(posicion)
	return h.heapify(posicion - 1)
}

func HeapSort[T any](elementos []T, funcion_cmp func(T, T) int) {
	heap := new(heap[T])
	heap.cmp = funcion_cmp
	heap.arr = elementos
	heap.cantidad = len(elementos)
	heap.heapify(heap.cantidad)
	heap.ordenar(heap.cantidad - 1)
	heap.cantidad = len(elementos)

}

func (h *heap[T]) ordenar(posicion int) int {
	if posicion < 0 {
		return 0
	}
	h.swap(0, posicion)
	h.cantidad--
	h.downHeap(0)
	return h.ordenar(posicion - 1)
}
