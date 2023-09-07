package Tablero

import (
	V "Algueiza/Vuelo"
)

type Tablero interface {
	//Recibe una fecha de despegue y un tipo vuelo que contiene toda la informacion del vuelo.
	GuardarVuelo(string, V.Vuelo)

	//Recibe una  cantidad de vuelos, modo, desde y hasta. Para lograr imprimir toda la informacion por pantalla.
	VerTablero(string, string, string, string) string

	//Recibe un tipo vuelo para borrar toda la informacion de un vuelo viejo.
	BorrarUnitario(V.Vuelo)

	//Recibe 2 fechas desde y hasta y devuelve un Arreglo con todos los vuelos eliminados del Tablero.
	Borrar(string, string) []V.Vuelo

	//Recibe un viaje origen con un destino y una fecha para imprimir el proximo vuelo que exist. Devuelve un mensaje de
	//que no existe dicho vuelo si las condiciones pedidas no coinciden con algun vuelo.
	SiguienteVuelo(string, string) (string, string)
}
