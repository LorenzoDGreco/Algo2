package tp0

// Swap intercambia dos valores enteros.
func Swap(x *int, y *int) {
	z := *x

	*x = *y
	*y = z

}

// Maximo devuelve la posición del mayor elemento del arreglo, o -1 si el el arreglo es de largo 0. Si el máximo
// elemento aparece más de una vez, se debe devolver la primera posición en que ocurre.
func Maximo(vector []int) int {

	if len(vector) == 0 {
		return -1 // Es buena o mala practica tener doble punto de retorno?
	}

	maximo_valor := vector[0]
	maximo_ind := 0

	for ind, valor := range vector {

		if vector[ind] > maximo_valor {
			maximo_valor = valor
			maximo_ind = ind
		}
	}

	return maximo_ind
}

// Comparar compara dos arreglos de longitud especificada.
// Devuelve -1 si el primer arreglo es menor que el segundo; 0 si son iguales; o 1 si el primero es el mayor.
// Un arreglo es menor a otro cuando al compararlos elemento a elemento, el primer elemento en el que difieren
// no existe o es menor.
func Comparar(vector1 []int, vector2 []int) int {
	// Lo mismo que el anterior, pero al ser una funcion recursiva
	// (si la respuesta anterior fue que está mal) se puede utilizar?
	if len(vector1) == 0 && len(vector2) == 0 {
		return 0
	} else if len(vector1) == 0 {
		return -1
	} else if len(vector2) == 0 {
		return 1
	} else if vector1[0] < vector2[0] {
		return -1
	} else if vector2[0] < vector1[0] {
		return 1
	}

	return Comparar(vector1[1:], vector2[1:])
}

// Seleccion ordena el arreglo recibido mediante el algoritmo de selección.
func Seleccion(vector []int) []int {
	slice := vector[:]

	for x := len(vector) - 1; x > 0; x-- {
		valMax := Maximo(slice)
		Swap(&vector[x], &vector[valMax])
		slice = slice[:len(slice)-1]
	}

	return vector
}

// Suma devuelve la suma de los elementos de un arreglo. En caso de no tener elementos, debe devolver 0.
// Esta función debe implementarse de forma RECURSIVA. Se puede usar una función auxiliar (que sea
// la recursiva).
func Suma(vector []int) int {
	if len(vector) == 0 {
		return 0
	}
	return vector[len(vector)-1] + Suma(vector[:len(vector)-1])
}

// EsCadenaCapicua devuelve si la cadena es un palíndromo. Es decir, si se lee igual al derecho que al revés.
// Esta función debe implementarse de forma RECURSIVA. Se puede usar una función auxiliar (que sea
// la recursiva).
func EsCadenaCapicua(cadena string) bool {
	//Creí que tenia que sacar los espacios en blanco pero si lo hago la que pide
	// " EE" sea falso porque cuenta ese espacio
	//  y me lo tira como true porque yo elimine ese espacio

	//var cadena_procesada string
	// for _, char := range cadena {
	// 	if string(char) != " " {
	// 		cadena_procesada += string(char)
	// 	}
	// }

	//Y con el codigo con la cadena_procesada este debería estár modularizado en una funcion recursiva
	if len(cadena) == 1 || len(cadena) == 0 {
		return true
	} else if len(cadena) == 2 {
		if cadena[:1] == cadena[1:] {
			return true
		} else {
			return false
		}
	} else if cadena[:1] != cadena[len(cadena)-1:] {
		return false
	}

	return EsCadenaCapicua(cadena[1 : len(cadena)-1])
}
