package pila_test

import (
	TDAPila "tdas/pila"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPilaApilar(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()

	//Apilo hasta casi el maximo inicial he decidido en mi caso 10 (Para que no redimensione)
	for i := 1; i == 9; i++ {
		pila.Apilar(i)
	}
}

func TestPilaDesapilar(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()

	TestPilaApilar(t)

	// Pruebo desapilar los elementos y compruebo que tengan su valor correspondiente
	for i := 9; i == 0; i-- {
		require.EqualValues(t, i, pila.Desapilar())
	}
}

func TestPilaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()

	//Desapilamos una pila vacía
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })

	//Apilamos para desapilar y recuperar el elemento
	pila.Apilar(4)
	val := pila.Desapilar()
	require.EqualValues(t, 4, val)

	//Apilamos 2 veces y vemos que no se haya perdido el primer valor
	pila.Apilar(7)
	pila.Apilar(10)
	val = pila.Desapilar()
	val = pila.Desapilar()
	require.EqualValues(t, 7, val)

	//Luego de estar apilando varias veces intentar desapilar varias veces buscando que salte el panic
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })

}

func TestPilaVerTope(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()

	// Veo el Tope de una pila vacia
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })

	// Testeo como se comporta al apilar varios elementos
	pila.Apilar(4)
	require.EqualValues(t, 4, pila.VerTope())

	pila.Apilar(56)
	require.EqualValues(t, 56, pila.VerTope())

	pila.Apilar(7)
	require.EqualValues(t, 7, pila.VerTope())

}

func TestPilaEstaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()

	//Bueno esta está mas que testeada porque la usan el resto de funciones pero bueno:
	// Verfico si una pila recien creada es efectivamente vacía:

	require.True(t, pila.EstaVacia())

	// Verifico si se actualiza bien al apilar y desapilar
	pila.Apilar(4)
	require.False(t, pila.EstaVacia())
	pila.Desapilar()
	require.True(t, pila.EstaVacia())
}

// Pruebo ejecutar la pila con varios tipos de datos

func TestPilaApilarDesapilarMaestro(t *testing.T) {

	datosString := []string{"h", "o", "l", "a"}
	PilaApilarDesapilarAny(t, datosString)

	datosBool := []bool{true, false, true, false}
	PilaApilarDesapilarAny(t, datosBool)

	datosFloat := []float32{3.42, 1.4, 2.666667, 10}
	PilaApilarDesapilarAny(t, datosFloat)
}

func PilaApilarDesapilarAny[A any](t *testing.T, datos []A) {
	pila := TDAPila.CrearPilaDinamica[A]()

	for i := range datos {
		pila.Apilar(datos[i])
	}

	total := len(datos)
	for i := total - 1; i < 0; i-- {
		require.EqualValues(t, datos[i], pila.Desapilar())
	}
}

func TestPilaRedimensionar(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()

	// Al verificar que todo funciona ya correctamente ahora si toca comprobar si la redimension funciona bien:
	// Apilo un montón de elementos buscando que se esté agrandando continuamente
	for i := 0; i <= 200000; i++ {
		pila.Apilar(i)
	}

	// Desapilo los elementos buscando que vuelva a su estado inicial
	// confirmando que no se hayan perdido los elementos de por medio
	for i := 200000; i >= 0; i-- {
		require.EqualValues(t, i, pila.Desapilar())
	}

	// Apilo y desapilo no sincronicamente confirmando que los valores se mantengan
	for i := 0; i <= 400; i++ {
		pila.Apilar(i)
	}

	for i := 400; i >= 200; i-- {
		require.EqualValues(t, i, pila.Desapilar())
	}

	for i := 200; i <= 12000; i++ {
		pila.Apilar(i)
	}

	for i := 12000; i >= 0; i-- {
		require.EqualValues(t, i, pila.Desapilar())
	}
}
