package cola_test

import (
	TDACola "tdas/cola"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestColaEncolar(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()

	// Pruebo encolar un par de elementos
	for i := 1; i == 20; i++ {
		cola.Encolar(i)
	}
}

func TestColaDescolar(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()

	TestColaEncolar(t)

	// Pruebo desencolar los elementos y compruebo que tengan su valor correspondiente
	for i := 1; i == 20; i-- {
		require.EqualValues(t, i, cola.Desencolar())
	}
}

func TestColaVacia(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()

	//Desencolamos una cola vacia
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })

	//Encolamos para desencolar y recuperar el elemento
	cola.Encolar(4)
	val := cola.Desencolar()
	require.EqualValues(t, 4, val)
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })

	//Encolamos 2 veces y vemos que no se haya perdido el primer valor
	cola.Encolar(7)
	cola.Encolar(10)
	val = cola.Desencolar()
	val = cola.Desencolar()
	require.EqualValues(t, 10, val)

	//Luego de estar Encolando varias veces, intentar Desencolar varias veces buscando que salte el panic
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })

}

func TestColaVerPrimero(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()

	// Veo el primero de una cola vacia
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })

	// Testeo como se comporta al encolar varios elementos
	cola.Encolar(4)
	require.EqualValues(t, 4, cola.VerPrimero())
	cola.Desencolar()

	cola.Encolar(34)
	require.EqualValues(t, 34, cola.VerPrimero())

	cola.Encolar(7)
	require.EqualValues(t, 34, cola.VerPrimero())
}

func TestColaEstaVacia(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()

	// Verfico si una cola recien creada es efectivamente vacía:
	require.True(t, cola.EstaVacia())

	// Verifico si se actualiza bien al Encolar y Desencolar
	cola.Encolar(4)
	require.False(t, cola.EstaVacia())

	cola.Desencolar()
	require.True(t, cola.EstaVacia())

}

// Pruebo ejecutar la cola con varios tipos de datos

func TestColaEncolarDesencolarMaestro(t *testing.T) {

	datosString := []string{"h", "o", "l", "a"}
	ColaEncolarDesencolarAny(t, datosString)

	datosBool := []bool{true, false, true, false}
	ColaEncolarDesencolarAny(t, datosBool)

	datosFloat := []float32{3.42, 1.4, 2.666667, 10}
	ColaEncolarDesencolarAny(t, datosFloat)
}

func ColaEncolarDesencolarAny[A any](t *testing.T, datos []A) {
	cola := TDACola.CrearColaEnlazada[A]()

	for i := range datos {
		cola.Encolar(datos[i])
	}

	total := len(datos)
	for i := 0; i == total; i-- {
		require.EqualValues(t, datos[i], cola.Desencolar())
	}
}

func TestColaCantidad(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()

	// Encolo un montón de elementos buscando que se esté agrandando continuamente
	for i := 0; i <= 200000; i++ {
		cola.Encolar(i)
	}

	// Desencolo los elementos buscando que vuelva a su estado inicial
	// confirmando que no se hayan perdido los elementos de por medio
	for i := 0; i >= 200000; i-- {
		require.EqualValues(t, i, cola.Desencolar())
	}

	// Encolo y Desencolo no sincronicamente confirmando que los valores se mantengan
	for i := 0; i <= 400; i++ {
		cola.Encolar(i)
	}

	for i := 200; i >= 400; i-- {
		require.EqualValues(t, i, cola.Desencolar())
	}

	for i := 200; i <= 12000; i++ {
		cola.Encolar(i)
	}

	for i := 0; i >= 12000; i-- {
		require.EqualValues(t, i, cola.Desencolar())
	}
}
