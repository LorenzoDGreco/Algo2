package diccionario_test

import (
	TDADiccionario "tdas/diccionario"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestArbolVacio(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](func(a, b int) int { return a - b })
	require.EqualValues(t, 0, dic.Cantidad())
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(10) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(5) })
}

func TestInsertarUnElemento(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](func(a, b int) int { return a - b })
	dic.Guardar(5, 10)
	require.True(t, dic.Pertenece(5))
	require.EqualValues(t, 1, dic.Cantidad())
}

func TestAgregarElementos(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](func(a, b int) int { return a - b })
	dic.Guardar(5, 10)
	dic.Guardar(0, 0)
	dic.Guardar(2, 90)
	dic.Guardar(200, 11)
	require.True(t, dic.Pertenece(5))
	require.EqualValues(t, 4, dic.Cantidad())
	require.EqualValues(t, 10, dic.Obtener(5))
	dic.Guardar(5, -20)
	require.EqualValues(t, -20, dic.Obtener(5))
}

func TestBorrarUnElemento(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](func(a, b int) int { return a - b })
	dic.Guardar(5, 10)
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, 10, dic.Borrar(5))
	require.EqualValues(t, 0, dic.Cantidad())
}

func TestBorrarVariosElementos(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](func(a, b int) int { return a - b })
	dic.Guardar(4, 2)
	dic.Guardar(2, 2)
	dic.Guardar(8, 1)
	dic.Guardar(6, 1)
	dic.Guardar(5, 1)
	dic.Guardar(3, 1)
	dic.Guardar(1, 1)
	dic.Guardar(9, 1)

	dic.Borrar(1) //Hoja
	require.False(t, dic.Pertenece(1))
	dic.Borrar(6) //Un hijo y confirmo que siga existiendo
	require.False(t, dic.Pertenece(6))
	require.True(t, dic.Pertenece(5))
	dic.Guardar(1, 1) // Preparo para que tenga 2 hijos
	require.True(t, dic.Pertenece(1))
	dic.Borrar(2) //2 hijos y confirmo que sigan existiendo
	require.False(t, dic.Pertenece(2))
	require.True(t, dic.Pertenece(1))
	require.True(t, dic.Pertenece(3))

}

func TestBorrarYAgregarVariosElementos(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](func(a, b int) int { return a - b })
	dic.Guardar(4, 2)
	dic.Guardar(2, 2)
	dic.Guardar(8, 1)
	dic.Guardar(6, 1)
	dic.Guardar(5, 1)
	dic.Guardar(3, 1)
	dic.Guardar(1, 1)
	dic.Guardar(9, 1)

	dic.Borrar(1) //Hoja
	require.False(t, dic.Pertenece(1))
	dic.Borrar(6) //Un hijo y confirmo que siga existiendo
	require.False(t, dic.Pertenece(6))
	require.True(t, dic.Pertenece(5))
	dic.Guardar(1, 1) // Preparo para que tenga 2 hijos
	require.True(t, dic.Pertenece(1))
	dic.Borrar(2) //2 hijos y confirmo que sigan existiendo
	require.False(t, dic.Pertenece(2))
	require.True(t, dic.Pertenece(1))
	require.True(t, dic.Pertenece(3))

	dic.Guardar(2, 6)
	dic.Guardar(11, 2)
	dic.Guardar(15, 24)
	dic.Guardar(6, 1)
	require.True(t, dic.Pertenece(2))
	require.True(t, dic.Pertenece(11))
	require.True(t, dic.Pertenece(15))
	require.True(t, dic.Pertenece(6))

}

func TestBorrarTodoYAgregar(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](func(a, b int) int { return a - b })
	dic.Guardar(4, 2)
	dic.Guardar(2, 2)
	dic.Guardar(8, 1)
	dic.Guardar(6, 1)
	dic.Guardar(5, 1)
	dic.Guardar(3, 1)
	dic.Guardar(1, 1)
	dic.Guardar(9, 1)

	dic.Borrar(1)
	require.False(t, dic.Pertenece(1))
	dic.Borrar(4)
	require.False(t, dic.Pertenece(4))
	dic.Borrar(6)
	require.False(t, dic.Pertenece(6))
	dic.Borrar(3)
	require.False(t, dic.Pertenece(3))
	dic.Borrar(9)
	require.False(t, dic.Pertenece(9))
	dic.Borrar(2)
	require.False(t, dic.Pertenece(2))
	dic.Borrar(8)
	require.False(t, dic.Pertenece(8))
	dic.Borrar(5)
	require.False(t, dic.Pertenece(5))

	dic.Guardar(2, 6)
	dic.Guardar(11, 2)
	dic.Guardar(15, 24)
	dic.Guardar(6, 1)
	require.True(t, dic.Pertenece(2))
	require.True(t, dic.Pertenece(11))
	require.True(t, dic.Pertenece(15))
	require.True(t, dic.Pertenece(6))

}

func TestIteradorInterno(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, *int](func(a, b int) int { return a - b })
	elem1 := 1
	dic.Guardar(4, &elem1)
	elem2 := 2
	dic.Guardar(2, &elem2)
	elem3 := 3
	dic.Guardar(8, &elem3)
	dic.Iterar(func(clave int, dato *int) bool {
		*dato = *dato + 1
		return true
	})

	require.EqualValues(t, 2, *dic.Obtener(4))
	require.EqualValues(t, 3, *dic.Obtener(2))
	require.EqualValues(t, 4, *dic.Obtener(8))
}

func TestIteradorInternoConRangoCompleto(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, *int](func(a, b int) int { return a - b })
	elem1 := 1
	dic.Guardar(4, &elem1)
	elem2 := 2
	dic.Guardar(2, &elem2)
	elem3 := 3
	dic.Guardar(8, &elem3)
	dic.IterarRango(nil, nil, func(clave int, dato *int) bool {
		*dato = *dato + 1
		return true
	})

	require.EqualValues(t, 2, *dic.Obtener(4))
	require.EqualValues(t, 3, *dic.Obtener(2))
	require.EqualValues(t, 4, *dic.Obtener(8))
}

func TestIteradorInternoConRangoFinal(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, *int](func(a, b int) int { return a - b })
	elem1 := 1
	dic.Guardar(4, &elem1)
	elem2 := 2
	dic.Guardar(2, &elem2)
	elem3 := 3
	dic.Guardar(8, &elem3)
	elem4 := 5
	dic.Guardar(6, &elem4)
	elem5 := 10
	dic.Guardar(5, &elem5)
	elem6 := 20
	dic.Guardar(14, &elem6)
	clave := 7
	dic.IterarRango(nil, &clave, func(clave int, dato *int) bool {
		*dato = *dato + 1
		return true
	})

	require.EqualValues(t, 2, *dic.Obtener(4))
	require.EqualValues(t, 3, *dic.Obtener(2))
	require.EqualValues(t, 3, *dic.Obtener(8))
	require.EqualValues(t, 11, *dic.Obtener(5))
	require.EqualValues(t, 6, *dic.Obtener(6))
	require.EqualValues(t, 20, *dic.Obtener(14))
}

func TestIteradorInternoConRangoInicio(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, *int](func(a, b int) int { return a - b })
	elem1 := 1
	dic.Guardar(6, &elem1)
	elem2 := 2
	dic.Guardar(2, &elem2)
	elem3 := 3
	dic.Guardar(8, &elem3)
	elem4 := 0
	dic.Guardar(3, &elem4)
	elem5 := 2
	dic.Guardar(4, &elem5)
	elem6 := 3
	dic.Guardar(7, &elem6)

	clave := 3
	dic.IterarRango(&clave, nil, func(clave int, dato *int) bool {
		*dato = *dato + 1
		return true
	})

	require.EqualValues(t, 3, *dic.Obtener(4))
	require.EqualValues(t, 2, *dic.Obtener(2))
	require.EqualValues(t, 1, *dic.Obtener(3))
	require.EqualValues(t, 2, *dic.Obtener(6))
	require.EqualValues(t, 4, *dic.Obtener(7))
	require.EqualValues(t, 4, *dic.Obtener(8))
}

func TestIteradorConRangoInicioYCondDeCor(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, *int](func(a, b int) int { return a - b })
	elem1 := 1
	dic.Guardar(6, &elem1)
	elem2 := 2
	dic.Guardar(2, &elem2)
	elem3 := 3
	dic.Guardar(8, &elem3)
	elem4 := 0
	dic.Guardar(3, &elem4)
	elem5 := 2
	dic.Guardar(4, &elem5)
	elem6 := 3
	dic.Guardar(7, &elem6)

	clave := 3
	dic.IterarRango(&clave, nil, func(clave int, dato *int) bool {
		*dato = *dato + 1
		if clave == 6 {
			return false
		}
		return true
	})

	require.EqualValues(t, 3, *dic.Obtener(4))
	require.EqualValues(t, 2, *dic.Obtener(2))
	require.EqualValues(t, 1, *dic.Obtener(3))
	require.EqualValues(t, 2, *dic.Obtener(6))
	require.EqualValues(t, 3, *dic.Obtener(7))
	require.EqualValues(t, 3, *dic.Obtener(8))
}

func TestIteradorInternoConRangoFinalYCondDeCorte(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, *int](func(a, b int) int { return a - b })
	elem1 := 1
	dic.Guardar(4, &elem1)
	elem2 := 2
	dic.Guardar(2, &elem2)
	elem3 := 3
	dic.Guardar(8, &elem3)
	elem4 := 5
	dic.Guardar(6, &elem4)
	elem5 := 10
	dic.Guardar(5, &elem5)
	elem6 := 20
	dic.Guardar(14, &elem6)
	clave := 7
	dic.IterarRango(nil, &clave, func(clave int, dato *int) bool {
		*dato = *dato + 1
		if clave == 5 {
			return false
		}
		return true
	})

	require.EqualValues(t, 2, *dic.Obtener(4))
	require.EqualValues(t, 3, *dic.Obtener(2))
	require.EqualValues(t, 3, *dic.Obtener(8))
	require.EqualValues(t, 11, *dic.Obtener(5))
	require.EqualValues(t, 5, *dic.Obtener(6))
	require.EqualValues(t, 20, *dic.Obtener(14))
}

func TestIteradorInternoConRangoInicioYFin(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, *int](func(a, b int) int { return a - b })
	elem1 := 1
	dic.Guardar(6, &elem1)
	elem2 := 2
	dic.Guardar(2, &elem2)
	elem3 := 3
	dic.Guardar(8, &elem3)
	elem4 := 0
	dic.Guardar(3, &elem4)
	elem5 := 2
	dic.Guardar(4, &elem5)
	elem6 := 3
	dic.Guardar(7, &elem6)

	clave := 3
	clave2 := 7
	dic.IterarRango(&clave, &clave2, func(clave int, dato *int) bool {
		*dato = *dato + 1
		return true
	})

	require.EqualValues(t, 3, *dic.Obtener(4))
	require.EqualValues(t, 2, *dic.Obtener(2))
	require.EqualValues(t, 1, *dic.Obtener(3))
	require.EqualValues(t, 2, *dic.Obtener(6))
	require.EqualValues(t, 4, *dic.Obtener(7))
	require.EqualValues(t, 3, *dic.Obtener(8))
}

func TestIteradorExternoVacio(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, string](func(a, b int) int { return a - b })
	iterador := dic.Iterador()
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.Siguiente() })
	require.False(t, iterador.HaySiguiente())

}

func TestIteradorExternoElementos(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](func(a, b int) int { return a - b })
	dic.Guardar(4, 4)
	dic.Guardar(2, 2)
	dic.Guardar(3, 3)
	dic.Guardar(1, 1)
	dic.Guardar(5, 5)
	dic.Guardar(6, 6)

	iterador := dic.Iterador()
	clave, valor := iterador.VerActual()
	require.EqualValues(t, 1, clave)
	require.EqualValues(t, 1, valor)
	iterador.Siguiente()

	clave, valor = iterador.VerActual()
	require.EqualValues(t, 2, clave)
	require.EqualValues(t, 2, valor)
	iterador.Siguiente()

	clave, valor = iterador.VerActual()
	require.EqualValues(t, 3, clave)
	require.EqualValues(t, 3, valor)
	iterador.Siguiente()

	clave, valor = iterador.VerActual()
	require.EqualValues(t, 4, clave)
	require.EqualValues(t, 4, valor)
	iterador.Siguiente()

	clave, valor = iterador.VerActual()
	require.EqualValues(t, 5, clave)
	require.EqualValues(t, 5, valor)
	iterador.Siguiente()

	clave, valor = iterador.VerActual()
	require.EqualValues(t, 6, clave)
	require.EqualValues(t, 6, valor)
	iterador.Siguiente()

	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.Siguiente() })
	require.False(t, iterador.HaySiguiente())
}

func Test2IteradoresExternos(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](func(a, b int) int { return a - b })
	dic.Guardar(3, 3)
	dic.Guardar(5, 5)
	dic.Guardar(2, 2)
	dic.Guardar(1, 1)
	dic.Guardar(4, 4)
	dic.Guardar(6, 6)

	iterador := dic.Iterador()
	clave, valor := iterador.VerActual()
	require.EqualValues(t, 1, clave)
	require.EqualValues(t, 1, valor)
	iterador.Siguiente()

	clave, valor = iterador.VerActual()
	require.EqualValues(t, 2, clave)
	require.EqualValues(t, 2, valor)
	iterador.Siguiente()

	clave, valor = iterador.VerActual()
	require.EqualValues(t, 3, clave)
	require.EqualValues(t, 3, valor)

	iterador2 := dic.Iterador()
	clave, valor = iterador2.VerActual()
	require.EqualValues(t, 1, clave)
	require.EqualValues(t, 1, valor)
	iterador2.Siguiente()
	clave, valor = iterador2.VerActual()
	require.EqualValues(t, 2, clave)
	require.EqualValues(t, 2, valor)
	iterador2.Siguiente()
	clave, valor = iterador2.VerActual()
	require.EqualValues(t, 3, clave)
	require.EqualValues(t, 3, valor)
}

func TestIteradorExternoConRangoInicial(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](func(a, b int) int { return a - b })
	dic.Guardar(3, 3)
	dic.Guardar(2, 2)
	dic.Guardar(1, 1)
	dic.Guardar(4, 4)
	dic.Guardar(5, 5)
	dic.Guardar(6, 6)
	claveInicio := 2
	iterador := dic.IteradorRango(&claveInicio, nil)
	clave, valor := iterador.VerActual()
	require.EqualValues(t, 2, clave)
	require.EqualValues(t, 2, valor)
	iterador.Siguiente()
	clave, valor = iterador.VerActual()
	require.EqualValues(t, 3, clave)
	require.EqualValues(t, 3, valor)
	iterador.Siguiente()
	clave, valor = iterador.VerActual()
	require.EqualValues(t, 4, clave)
	require.EqualValues(t, 4, valor)
	iterador.Siguiente()
	clave, valor = iterador.VerActual()
	require.EqualValues(t, 5, clave)
	require.EqualValues(t, 5, valor)
	iterador.Siguiente()
	clave, valor = iterador.VerActual()
	require.EqualValues(t, 6, clave)
	require.EqualValues(t, 6, valor)
	iterador.Siguiente()

	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.Siguiente() })
	require.False(t, iterador.HaySiguiente())
}

func TestIteradorExternoConRangoFinal(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](func(a, b int) int { return a - b })
	dic.Guardar(3, 3)
	dic.Guardar(2, 2)
	dic.Guardar(1, 1)
	dic.Guardar(4, 4)
	dic.Guardar(5, 5)
	dic.Guardar(6, 6)
	claveFinal := 3
	iterador := dic.IteradorRango(nil, &claveFinal)
	clave, valor := iterador.VerActual()
	require.EqualValues(t, 1, clave)
	require.EqualValues(t, 1, valor)
	iterador.Siguiente()
	clave, valor = iterador.VerActual()
	require.EqualValues(t, 2, clave)
	require.EqualValues(t, 2, valor)
	iterador.Siguiente()
	clave, valor = iterador.VerActual()
	require.EqualValues(t, 3, clave)
	require.EqualValues(t, 3, valor)
	iterador.Siguiente()

	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.Siguiente() })
	require.False(t, iterador.HaySiguiente())
}

func TestIteradorExternoConRangoInicialYFinal(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](func(a, b int) int { return a - b })
	dic.Guardar(3, 3)
	dic.Guardar(2, 2)
	dic.Guardar(1, 1)
	dic.Guardar(4, 4)
	dic.Guardar(5, 5)
	dic.Guardar(6, 6)
	claveInicial := 2
	claveFinal := 4
	iterador := dic.IteradorRango(&claveInicial, &claveFinal)
	clave, valor := iterador.VerActual()
	require.EqualValues(t, 2, clave)
	require.EqualValues(t, 2, valor)
	iterador.Siguiente()
	clave, valor = iterador.VerActual()
	require.EqualValues(t, 3, clave)
	require.EqualValues(t, 3, valor)
	iterador.Siguiente()
	clave, valor = iterador.VerActual()
	require.EqualValues(t, 4, clave)
	require.EqualValues(t, 4, valor)
	iterador.Siguiente()

	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.Siguiente() })
	require.False(t, iterador.HaySiguiente())
}

func TestIterarRangoCombinaciones(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](func(a, b int) int { return a - b })
	dic.Guardar(1, 1)
	dic.Guardar(2, 2)
	dic.Guardar(3, 3)
	dic.Guardar(4, 4)
	dic.Guardar(5, 5)
	dic.Guardar(6, 6)
	dic.Guardar(7, 7)
	inicio := 2
	fin := 5
	iterador := dic.IteradorRango(&inicio, &fin)
	clave, valor := iterador.VerActual()
	require.EqualValues(t, 2, clave)
	require.EqualValues(t, 2, valor)
	require.True(t, iterador.HaySiguiente())
	iterador.Siguiente()
	clave, valor = iterador.VerActual()
	require.EqualValues(t, 3, clave)
	require.EqualValues(t, 3, valor)
	require.True(t, iterador.HaySiguiente())
	iterador.Siguiente()
	clave, valor = iterador.VerActual()
	require.EqualValues(t, 4, clave)
	require.EqualValues(t, 4, valor)
	require.True(t, iterador.HaySiguiente())
	iterador.Siguiente()
	clave, valor = iterador.VerActual()
	require.EqualValues(t, 5, clave)
	require.EqualValues(t, 5, valor)
	require.True(t, iterador.HaySiguiente())
	iterador.Siguiente()
	require.False(t, iterador.HaySiguiente())
}

func TestIterarRangoCombinacionesParte2(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](func(a, b int) int { return a - b })
	dic.Guardar(3, 3)
	dic.Guardar(2, 2)
	dic.Guardar(1, 1)
	dic.Guardar(4, 4)
	dic.Guardar(5, 5)
	dic.Guardar(6, 6)
	dic.Guardar(7, 7)
	inicio := 2
	fin := 5
	iterador := dic.IteradorRango(&inicio, &fin)
	clave, valor := iterador.VerActual()
	require.EqualValues(t, 2, clave)
	require.EqualValues(t, 2, valor)
	require.True(t, iterador.HaySiguiente())
	iterador.Siguiente()
	clave, valor = iterador.VerActual()
	require.EqualValues(t, 3, clave)
	require.EqualValues(t, 3, valor)
	require.True(t, iterador.HaySiguiente())
	iterador.Siguiente()
	clave, valor = iterador.VerActual()
	require.EqualValues(t, 4, clave)
	require.EqualValues(t, 4, valor)
	require.True(t, iterador.HaySiguiente())
	iterador.Siguiente()
	clave, valor = iterador.VerActual()
	require.EqualValues(t, 5, clave)
	require.EqualValues(t, 5, valor)
	require.True(t, iterador.HaySiguiente())
	iterador.Siguiente()
	require.False(t, iterador.HaySiguiente())
}
