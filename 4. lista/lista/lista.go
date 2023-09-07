package lista

type Lista[T any] interface {

	// EstaVacia devuelve verdadero si la cola no tiene elementos encolados, false en caso contrario.
	EstaVacia() bool

	// VerPrimero obtiene el valor del primero de la cola. Si está vacía, entra en pánico con un mensaje
	// "La lista esta vacia".
	VerPrimero() T

	// VerUltimo obtiene el valor del ultimo de la lista. Si está vacía, entra en pánico con un mensaje
	// "La lista esta vacia".
	VerUltimo() T

	// InsertarPrimero agrega un nuevo elemento a la lista, al principio de la misma.
	InsertarPrimero(T)

	// InsertarUltimo agrega un nuevo elemento a la lista, al final de la misma.
	InsertarUltimo(T)

	//Largo devuelve un entero que indica la cantidad de elementos en la lista.
	Largo() int

	// BorrarPrimero saca el primer elemento de la lista. Si la lista tiene elementos, se quita el primero de la misma,
	// y se devuelve ese valor. Si está vacía, entra en pánico con un mensaje "La lista esta vacia".
	BorrarPrimero() T

	//La funcion iterar toma una funcion por parametro que debe crear el usuario, y le aplica a todos los elementos
	//de la lista dicha funcion.
	Iterar(visitar func(T) bool)

	// El iterador externo le otorga unas primitivas al usuario externo para que las use a su antojo
	Iterador() IteradorLista[T]
}

type IteradorLista[T any] interface {
	// VerActual muestra el elemento en el que se encuentra la lectura.
	VerActual() T

	// Verifica si hay una siguiente lectura valida
	HaySiguiente() bool

	// Avanaza al siguiente elemento.
	//Si no hay elemento siguiente, levanta un panic que dice "No hay mas elementos"
	Siguiente()

	// Inserta a continuacion de donde se encuentra actualmente.
	Insertar(T)

	// Borra el elemento en donde se encuentra actualmente.
	Borrar() T
}
