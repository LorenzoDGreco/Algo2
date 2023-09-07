package lista_test

import (
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.True(t, true, lista.EstaVacia())
	require.EqualValues(t, 0, lista.Largo())

}

func TestListaListar(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	// Probamos listar un par de elementos al inicio
	for i := 1; i <= 5; i++ {
		lista.InsertarPrimero(i)
		require.Equal(t, i, lista.VerPrimero())
	}

	for i := 6; i <= 11; i++ {
		lista.InsertarUltimo(i)
		require.Equal(t, i, lista.VerUltimo())
	}
	require.EqualValues(t, 11, lista.Largo())
}

func TestListaDeslistar(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	TestListaListar(t)

	// Probamos deslistar los elementos y compruebo que tengan su valor correspondiente
	for i := 11; i <= 6; i-- {
		require.Equal(t, i, lista.BorrarPrimero())
	}

	for i := 5; i <= 1; i-- {
		require.Equal(t, i, lista.BorrarPrimero())
	}
}

func TestVaciarLista(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	//listamos para deslistar y recuperar el elemento
	lista.InsertarPrimero(4)
	require.Equal(t, 4, lista.BorrarPrimero())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })

	//Enlistamos un par de veces y vemos que no se hayan perdido los valores
	TestListaDeslistar(t)

	//Ahora podemos deslistar varias veces buscando que salte el panic
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })

}

func TestVerPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	// Veo el primero de una lista vacia
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })

	// Testeo como se comporta al insertar varios elementos al principio
	lista.InsertarPrimero(4)
	require.EqualValues(t, 4, lista.VerPrimero())
	lista.BorrarPrimero()
	lista.InsertarPrimero(34)
	require.EqualValues(t, 34, lista.VerPrimero())
	lista.InsertarPrimero(7)
	require.EqualValues(t, 7, lista.VerPrimero())
	//Testeo como se comporta al insertar elementos al principio y al final
	lista.InsertarPrimero(7)
	lista.InsertarPrimero(0)
	require.EqualValues(t, 0, lista.VerPrimero())
	lista.InsertarUltimo(80)
	require.EqualValues(t, 0, lista.VerPrimero())
}

func TestVerUltimo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	// Veo el primero de una lista vacia
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })

	// Testeo como se comporta al insertar varios elementos al final
	lista.InsertarUltimo(4)
	require.EqualValues(t, 4, lista.VerUltimo())
	lista.BorrarPrimero()
	lista.InsertarUltimo(34)
	require.EqualValues(t, 34, lista.VerUltimo())
	lista.InsertarUltimo(7)
	require.EqualValues(t, 7, lista.VerUltimo())
	//Al insertar elementos al final y al principio, el valor final no cambia
	lista.InsertarUltimo(7)
	lista.InsertarUltimo(0)
	require.EqualValues(t, 0, lista.VerUltimo())
	lista.InsertarPrimero(80)
	require.EqualValues(t, 0, lista.VerUltimo())
}

func TestListaEstaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	// Verfico si una lista recien creada es efectivamente vacía:
	require.True(t, lista.EstaVacia())

	// Verifico si se actualiza bien al insertar y al deslistar
	lista.InsertarPrimero(4)
	require.False(t, lista.EstaVacia())

	lista.BorrarPrimero()
	require.True(t, lista.EstaVacia())

	lista.InsertarUltimo(7)
	require.False(t, lista.EstaVacia())

	lista.BorrarPrimero()
	require.True(t, lista.EstaVacia())

	//Con varios elementos
	TestListaDeslistar(t)
	require.True(t, lista.EstaVacia())

}

// Probamos ejecutar la cola con varios tipos de datos

func TestListaListarDeslistarMaestro(t *testing.T) {

	datosString := []string{"h", "o", "l", "a"}
	ListaListarDeslistarAny(t, datosString)

	datosBool := []bool{true, false, true, false}
	ListaListarDeslistarAny(t, datosBool)

	datosFloat := []float32{3.42, 1.4, 2.666667, 10}
	ListaListarDeslistarAny(t, datosFloat)
}

func ListaListarDeslistarAny[A any](t *testing.T, datos []A) {
	lista := TDALista.CrearListaEnlazada[A]()

	for i := range datos {
		lista.InsertarUltimo(datos[i])
	}

	total := len(datos)
	for i := 0; i < total; i++ {
		require.EqualValues(t, datos[i], lista.BorrarPrimero())
	}
}

func TestListaCantidad(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	//Insertamos un montón de elementos buscando que se esté agrandando continuamente
	for i := 0; i < 200000; i++ {
		lista.InsertarUltimo(i)
		require.EqualValues(t, i, lista.VerUltimo())
	}

	// Deslisto los elementos buscando que vuelva a su estado inicial
	// confirmando que no se hayan perdido los elementos de por medio
	for i := 0; i < 200000; i++ {
		require.EqualValues(t, i, lista.BorrarPrimero())
	}

	// Inserto y elimino no sincronicamente confirmando que los valores se mantengan
	require.True(t, true, lista.EstaVacia())
	for i := 0; i < 400; i++ {
		lista.InsertarUltimo(i)
	}
	require.False(t, false, lista.EstaVacia())
	for j := 0; j < 400; j++ {
		variable := lista.BorrarPrimero()
		require.EqualValues(t, variable, j)
	}
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.True(t, true, lista.EstaVacia())
}

// Probamos el iterador interno sobre una lista vacia
func TestIteradorInternoSobreListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.Iterar(func(v int) bool {
		return true
	})
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
}

// Probamos que el iterador interno recorra toda la lista
func TestIteradorInternoConElementos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(0)
	lista.InsertarPrimero(3)
	require.EqualValues(t, 2, lista.VerUltimo())
	sumatoria := 0
	dirSumatoria := &sumatoria
	lista.Iterar(func(v int) bool {
		*dirSumatoria += v
		return true
	})
	require.EqualValues(t, 5, sumatoria)
}

// Probamos el iterador interno con una condicion de corte
func TestIteradorInternoConCondicionDeCorte(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(0)
	lista.InsertarPrimero(9)
	lista.InsertarPrimero(302932)
	lista.InsertarPrimero(10)
	lista.InsertarPrimero(33)
	contador := 0
	dirContador := &contador
	lista.Iterar(func(v int) bool {
		*dirContador += 1
		return contador < 4
	})
	require.EqualValues(t, 4, contador)
	require.EqualValues(t, 2, lista.VerUltimo())
	require.EqualValues(t, 33, lista.VerPrimero())
}

// Probamos el iterador en una lista vacia
func TestIteradorExternoListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iterador := lista.Iterador()
	require.False(t, iterador.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.Borrar() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.VerActual() })
}

// Probamos caso borde de tener un unico elemento
func TestIteradorExternoListaUnElemento(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	iterador := lista.Iterador()

	// Probamos que el primero sea igual al actual del iterador
	require.EqualValues(t, lista.VerPrimero(), iterador.VerActual())

	// Probamos que al borrar se actualice a una lista vacia
	require.EqualValues(t, 1, iterador.Borrar())
	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.Borrar() })

	//Probamos que al avanzar nos quedamos en nil, y saltan los panic
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.Siguiente() })
}

// Probamos leer todos los elementos de la lista
func TestIteradorExternoIterarLista(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(3)
	lista.InsertarPrimero(4)
	lista.InsertarPrimero(5)
	lista.InsertarPrimero(6)
	iterador := lista.Iterador()
	contador := 6

	for iterador.HaySiguiente() {
		require.EqualValues(t, contador, iterador.VerActual())
		contador--
		iterador.Siguiente()
	}
}

// Probamos en borrar todos los elementos en una lista con elementos previos
func TestIteradorExternoBorrarElementos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	for i := 1; i <= 20; i++ {
		lista.InsertarUltimo(i)
	}

	for iterador := lista.Iterador(); iterador.HaySiguiente() != false; iterador.Borrar() {
		require.EqualValues(t, iterador.VerActual(), lista.VerPrimero())
	}
}

// Probamos insertar elementos a una lista ya existente con elementos previos
func TestIteradorExternoInsertarElementos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(3)
	lista.InsertarPrimero(4)
	lista.InsertarPrimero(5)
	lista.InsertarPrimero(6)
	iterador := lista.Iterador()
	iterador.Insertar(10)
	require.EqualValues(t, 10, iterador.VerActual())
	require.EqualValues(t, 10, lista.VerPrimero())
	iterador.Insertar(-1)
	require.EqualValues(t, -1, iterador.VerActual())
	iterador.Siguiente()
	require.EqualValues(t, 10, iterador.VerActual())

}

// Probamos agregar un elemento desde el iterador con una lista vacia
func TestIteradorExternoAgregaPrincipio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iterador := lista.Iterador()
	iterador.Insertar(5)
	require.EqualValues(t, 5, lista.VerPrimero())
}

// Probamos insertar un par de elementos y agregar uno en el medio
func TestIteradorExternoAgregaMedio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	iterador := lista.Iterador()
	require.EqualValues(t, 1, iterador.VerActual())
	iterador.Siguiente()
	iterador.Insertar(10)
	require.EqualValues(t, 10, iterador.VerActual())
	iterador.Siguiente()
	require.EqualValues(t, 2, iterador.VerActual())
	iterador.Siguiente()
	require.EqualValues(t, 3, iterador.VerActual())
}

// Probamos con varios elementos
func TestIteradorExternoVolumen(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i <= 1000; i++ {
		lista.InsertarUltimo(i)
	}
	iterador := lista.Iterador()
	contador := 0
	for iterador.HaySiguiente() {
		require.EqualValues(t, contador, iterador.VerActual())
		contador++
		iterador.Siguiente()
	}
}

// Probamos las primitivas principales
func TestIteradorExternoTresValores(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	iterador := lista.Iterador()
	require.EqualValues(t, 1, iterador.VerActual())
	require.True(t, iterador.HaySiguiente())
	iterador.Siguiente()
	require.EqualValues(t, 2, iterador.VerActual())
	require.True(t, iterador.HaySiguiente())
	iterador.Siguiente()
	require.EqualValues(t, 3, iterador.VerActual())
	require.True(t, iterador.HaySiguiente())
	iterador.Siguiente()
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.Siguiente() })
	require.False(t, iterador.HaySiguiente())
}

// Probamos eliminar un elemento que estaba solo en la lista
func TestAgregarEliminarUnElemento(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	iterador := lista.Iterador()
	require.EqualValues(t, 1, iterador.VerActual())
	require.EqualValues(t, 1, iterador.Borrar())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.Borrar() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })

}

// Probamos agregar un elemento en el medio
func TestAgregarElementoMedio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	lista.InsertarUltimo(4)
	lista.InsertarUltimo(5)
	lista.InsertarUltimo(6)
	iterador := lista.Iterador()
	iterador.Siguiente()
	iterador.Siguiente()
	iterador.Siguiente()
	iterador.Insertar(50)
	for i := 1; i < 4; i++ {
		require.EqualValues(t, i, lista.BorrarPrimero())
	}
	require.EqualValues(t, 50, lista.VerPrimero())
}

func TestInsertarInicioYMedioConDosIteradores(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	lista.InsertarUltimo(4)
	lista.InsertarUltimo(5)
	lista.InsertarUltimo(6)
	iterador := lista.Iterador()
	require.EqualValues(t, 1, iterador.VerActual())
	iterador.Insertar(50)
	require.EqualValues(t, 50, iterador.VerActual())
	iterador.Siguiente()
	require.EqualValues(t, 1, iterador.VerActual())
	iterador.Siguiente()
	require.EqualValues(t, 2, iterador.VerActual())
	iterador.Siguiente()
	require.EqualValues(t, 3, iterador.VerActual())
	iterador.Insertar(50)
	require.EqualValues(t, 50, iterador.VerActual())
	iterador.Siguiente()
	require.EqualValues(t, 3, iterador.VerActual())
	iterador.Siguiente()
	require.EqualValues(t, 4, iterador.VerActual())
	iterador.Siguiente()
	require.EqualValues(t, 5, iterador.VerActual())
	iterador.Siguiente()
	require.EqualValues(t, 6, iterador.VerActual())
	iterador.Siguiente()

	require.False(t, iterador.HaySiguiente())

	iterador2 := lista.Iterador()
	require.EqualValues(t, 50, iterador2.VerActual())
	iterador2.Siguiente()
	require.EqualValues(t, 1, iterador2.VerActual())
	iterador2.Siguiente()
	require.EqualValues(t, 2, iterador2.VerActual())
	iterador2.Siguiente()
	require.EqualValues(t, 50, iterador2.VerActual())
	iterador2.Siguiente()
	require.EqualValues(t, 3, iterador2.VerActual())
	iterador2.Siguiente()
	require.EqualValues(t, 4, iterador2.VerActual())
	iterador2.Siguiente()
	require.EqualValues(t, 5, iterador2.VerActual())
	iterador2.Siguiente()
	require.EqualValues(t, 6, iterador2.VerActual())
	iterador2.Siguiente()
	require.False(t, iterador2.HaySiguiente())
}

func TestIteradorExternoInsertarAlFinal(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	iterador := lista.Iterador()
	iterador.Siguiente()
	iterador.Siguiente()
	iterador.Siguiente()
	iterador.Insertar(0)
	require.True(t, iterador.HaySiguiente())
	iterador.Siguiente()
	require.False(t, iterador.HaySiguiente())

}

func TestBorrarAlInicio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	lista.InsertarUltimo(4)
	iterador := lista.Iterador()
	require.EqualValues(t, 1, iterador.Borrar())
	require.EqualValues(t, 2, iterador.Borrar())
	require.EqualValues(t, 3, iterador.Borrar())
	require.EqualValues(t, 4, iterador.VerActual())
	require.True(t, iterador.HaySiguiente())
	require.EqualValues(t, lista.VerPrimero(), iterador.VerActual())
}

func TestBorrarAlFinal(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	lista.InsertarUltimo(4)
	iterador := lista.Iterador()
	iterador.Siguiente()
	iterador.Siguiente()
	iterador.Siguiente()
	require.EqualValues(t, 4, iterador.Borrar())
	require.EqualValues(t, 3, lista.VerUltimo())
	iterador2 := lista.Iterador()
	require.EqualValues(t, lista.VerPrimero(), iterador2.Borrar())
	require.EqualValues(t, lista.VerPrimero(), iterador2.Borrar())
	require.EqualValues(t, lista.VerPrimero(), iterador2.Borrar())
}

func TestInsertarConIteradorBorrarConPrimitiva(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	lista.InsertarUltimo(4)
	iterador := lista.Iterador()
	iterador.Siguiente()
	require.EqualValues(t, 2, iterador.VerActual())
	iterador.Insertar(0)
	require.EqualValues(t, 0, iterador.VerActual())
	iterador.Siguiente()
	require.EqualValues(t, 2, iterador.VerActual())
	iterador.Siguiente()
	require.EqualValues(t, 3, iterador.VerActual())
	iterador.Siguiente()
	require.EqualValues(t, 4, iterador.VerActual())

	require.EqualValues(t, 1, lista.BorrarPrimero())
	require.EqualValues(t, 0, lista.BorrarPrimero())
	require.EqualValues(t, 2, lista.BorrarPrimero())
	require.EqualValues(t, 3, lista.BorrarPrimero())
	require.EqualValues(t, 4, lista.BorrarPrimero())
}

func TestIteradorExternoInsertarListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iterador := lista.Iterador()
	iterador.Insertar(1)
	iterador.Insertar(2)
	iterador.Siguiente()
	iterador.Insertar(3)
	iterador.Insertar(4)
	iterador.Siguiente()
	iterador.Insertar(7)
	iterador.Insertar(6)
	iterador.Siguiente()
	iterador.Siguiente()
	iterador.Siguiente()
	iterador.Siguiente()
	require.False(t, iterador.HaySiguiente())

	require.Equal(t, 2, lista.BorrarPrimero())
	require.Equal(t, 4, lista.BorrarPrimero())
	require.Equal(t, 6, lista.BorrarPrimero())
	require.Equal(t, 7, lista.BorrarPrimero())
	require.Equal(t, 3, lista.BorrarPrimero())
	require.Equal(t, 1, lista.BorrarPrimero())
}
