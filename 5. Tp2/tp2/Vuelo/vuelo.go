package Vuelo

type Vuelo interface {
	//Devuelve el numero de vuelo como string.
	VerNum() string

	//Devuelve la Prioridad del viaje como int.
	VerPrioridad() int

	//Imprime el vuelo con el formato acorde a un vuelo completo.
	ImprimirVuelo() string

	//Devuelve la fecha del vuelo como string.
	FechaVuelo() string

	//Imprime el vuelo con el formato requerido en el tablero.
	ImprimirTablero() string

	//Devuelve el recorrido del vuelo.
	VerViaje() string

	//Devuelve la prioridad del vuelo.
	ImprimirPrioridad() string
}
