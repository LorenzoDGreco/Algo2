package cola_prioridad_test

import (
	"strings"
	TDAHeap "tdas/cola_prioridad"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHeapVacio(t *testing.T) {
	heap := TDAHeap.CrearHeap(func(a, b int) int { return a - b })

	require.True(t, heap.EstaVacia())
	require.EqualValues(t, 0, heap.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestEncolarElementos(t *testing.T) {
	heap := TDAHeap.CrearHeap(func(a, b int) int { return a - b })

	heap.Encolar(15)
	require.EqualValues(t, 1, heap.Cantidad())
	require.EqualValues(t, 15, heap.VerMax())
	heap.Encolar(10)
	require.EqualValues(t, 2, heap.Cantidad())
	require.EqualValues(t, 15, heap.VerMax())
	heap.Encolar(7)
	require.EqualValues(t, 3, heap.Cantidad())
	require.EqualValues(t, 15, heap.VerMax())
	heap.Encolar(4)
	require.EqualValues(t, 4, heap.Cantidad())
	require.EqualValues(t, 15, heap.VerMax())
	heap.Encolar(2)
	require.EqualValues(t, 5, heap.Cantidad())
	require.EqualValues(t, 15, heap.VerMax())
	heap.Encolar(1)
	require.EqualValues(t, 6, heap.Cantidad())
	require.EqualValues(t, 15, heap.VerMax())
}

func TestEncolarStrings(t *testing.T) {
	heap := TDAHeap.CrearHeap(func(a, b string) int { return strings.Compare(a, b) })
	//https://www.youtube.com/watch?v=m1H-kuJbT8E&t
	heap.Encolar("776.420")
	heap.Encolar("La recaudacion para esta nueva edicion")
	heap.Encolar("Del superclasico Argentino")
	heap.Encolar("MARTEEEEEEEEEEEEEEEEEEEEEEEEN")
	heap.Encolar("GOOOOOOOOOOOOOOOOOOOOL")
	require.EqualValues(t, 5, heap.Cantidad())
	require.EqualValues(t, "MARTEEEEEEEEEEEEEEEEEEEEEEEEN", heap.VerMax())
}

func TestUpHeap(t *testing.T) {
	heap := TDAHeap.CrearHeap(func(a, b int) int { return a - b })

	heap.Encolar(1)
	require.EqualValues(t, 1, heap.VerMax())
	require.EqualValues(t, 1, heap.Cantidad())
	heap.Encolar(2)
	require.EqualValues(t, 2, heap.VerMax())
	require.EqualValues(t, 2, heap.Cantidad())
	heap.Encolar(4)
	require.EqualValues(t, 4, heap.VerMax())
	require.EqualValues(t, 3, heap.Cantidad())
	heap.Encolar(7)
	require.EqualValues(t, 7, heap.VerMax())
	require.EqualValues(t, 4, heap.Cantidad())
	heap.Encolar(10)
	require.EqualValues(t, 10, heap.VerMax())
	require.EqualValues(t, 5, heap.Cantidad())
	heap.Encolar(15)
	require.EqualValues(t, 15, heap.VerMax())
	require.EqualValues(t, 6, heap.Cantidad())
}

func TestDesencolarElementos(t *testing.T) {
	heap := TDAHeap.CrearHeap(func(a, b int) int { return a - b })

	heap.Encolar(11)
	require.EqualValues(t, 1, heap.Cantidad())
	heap.Encolar(9)
	require.EqualValues(t, 2, heap.Cantidad())
	heap.Encolar(9)
	require.EqualValues(t, 3, heap.Cantidad())
	heap.Encolar(5)
	require.EqualValues(t, 4, heap.Cantidad())
	heap.Encolar(3)
	require.EqualValues(t, 5, heap.Cantidad())
	heap.Encolar(2)
	require.EqualValues(t, 6, heap.Cantidad())
	require.EqualValues(t, 11, heap.VerMax())

	require.EqualValues(t, 11, heap.Desencolar())

	require.EqualValues(t, 5, heap.Cantidad())
	require.EqualValues(t, 9, heap.VerMax())
	require.EqualValues(t, 9, heap.Desencolar())

	require.EqualValues(t, 4, heap.Cantidad())
	require.EqualValues(t, 9, heap.VerMax())
	require.EqualValues(t, 9, heap.Desencolar())

	require.EqualValues(t, 3, heap.Cantidad())
	require.EqualValues(t, 5, heap.VerMax())
	require.EqualValues(t, 5, heap.Desencolar())

	require.EqualValues(t, 2, heap.Cantidad())
	require.EqualValues(t, 3, heap.VerMax())
	require.EqualValues(t, 3, heap.Desencolar())

	require.EqualValues(t, 1, heap.Cantidad())
	require.EqualValues(t, 2, heap.VerMax())
	require.EqualValues(t, 2, heap.Desencolar())

	require.EqualValues(t, 0, heap.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })

}

func TestHeapDown(t *testing.T) {
	heap := TDAHeap.CrearHeap(func(a, b int) int { return a - b })

	heap.Encolar(10)
	heap.Encolar(5)
	heap.Encolar(5)
	heap.Encolar(5)

	heap.Desencolar()
	//--------------
	require.EqualValues(t, 5, heap.Desencolar())
	require.EqualValues(t, 5, heap.Desencolar())
	require.EqualValues(t, 5, heap.Desencolar())
	//--------------

	heap.Encolar(10)
	heap.Encolar(4)
	heap.Encolar(5)
	heap.Encolar(3)
	heap.Encolar(3)
	heap.Encolar(5)

	heap.Desencolar()
	//--------------
	require.EqualValues(t, 5, heap.Desencolar())
	require.EqualValues(t, 5, heap.Desencolar())
	require.EqualValues(t, 4, heap.Desencolar())
	require.EqualValues(t, 3, heap.Desencolar())
	require.EqualValues(t, 3, heap.Desencolar())
	//--------------

}

func TestRedimension(t *testing.T) {
	heap := TDAHeap.CrearHeap(func(a, b int) int { return a - b })

	for i := 0; i <= 1000; i++ {
		heap.Encolar(i)
		require.EqualValues(t, i+1, heap.Cantidad())
		require.EqualValues(t, i, heap.VerMax())
	}
	for i := 1000; i >= 0; i-- {
		require.EqualValues(t, i, heap.Desencolar())
		require.EqualValues(t, i, heap.Cantidad())
	}
}

func TestHeapifyArrVacio(t *testing.T) {
	arr := []int{}
	heap := TDAHeap.CrearHeapArr(arr, func(a, b int) int { return a - b })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
	require.EqualValues(t, 0, heap.Cantidad())

	//Encolamos a un heap vacío
	heap.Encolar(1)
	heap.Encolar(2)
	heap.Encolar(4)
}

func TestHeapify(t *testing.T) {
	arr := []int{5, 3, 4, 10, 6, 7}
	heap := TDAHeap.CrearHeapArr(arr, func(a, b int) int { return a - b })

	require.EqualValues(t, 10, heap.VerMax())
	require.EqualValues(t, 6, heap.Cantidad())
	require.EqualValues(t, 10, heap.Desencolar())
	require.EqualValues(t, 7, heap.Desencolar())
	require.EqualValues(t, 6, heap.Desencolar())
	require.EqualValues(t, 5, heap.Desencolar())
	require.EqualValues(t, 4, heap.Desencolar())
	require.EqualValues(t, 3, heap.Desencolar())
}

func TestHeapSort(t *testing.T) {
	arr := []int{5, 4, 2, 1, 6, 3, 7}
	TDAHeap.HeapSort(arr, func(a, b int) int { return a - b })
	for i := 0; i < len(arr); i++ {
		require.EqualValues(t, i+1, arr[i])
	}
}

func TestHeapSortStrings(t *testing.T) {
	arr := []string{"Pedro", "Ana Maria", "Juan", "Juan Roman Riquelme", "Martin Palermo"}
	TDAHeap.HeapSort(arr, func(a, b string) int { return strings.Compare(a, b) })
	require.EqualValues(t, "Ana Maria", arr[0])
	require.EqualValues(t, "Juan", arr[1])
	require.EqualValues(t, "Juan Roman Riquelme", arr[2])
	require.EqualValues(t, "Martin Palermo", arr[3])
	require.EqualValues(t, "Pedro", arr[4])
}
