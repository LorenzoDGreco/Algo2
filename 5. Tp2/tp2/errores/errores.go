package errores

import "strings"

const (
	TODO_OK          = "OK"
	ASC              = "asc"
	DESC             = "desc"
	ERROR            = "Error en comando "
	VER_TABLERO      = "ver_tablero"
	ARGEGAR_ARCHIVO  = "agregar_archivo"
	PRIORIDAD_VUELOS = "prioridad_vuelos"
	SIGUIENTE_VUELO  = "siguiente_vuelo"
	INFO_VUELO       = "info_vuelo"
	BORRAR           = "borrar"
	ACCION           = 0
)

type ErrorParametros struct {
	Args []string
}

func (e ErrorParametros) Error() string {
	if e.Args[ACCION] == ARGEGAR_ARCHIVO {
		if len(e.Args) != 2 {
			return ERROR + ARGEGAR_ARCHIVO
		}
	}
	if e.Args[ACCION] == VER_TABLERO {
		if len(e.Args) != 5 {
			return ERROR + VER_TABLERO
		}
	}
	if e.Args[ACCION] == INFO_VUELO {
		if len(e.Args) != 2 {
			return ERROR + INFO_VUELO
		}
	}
	if e.Args[ACCION] == PRIORIDAD_VUELOS {
		if len(e.Args) != 2 {
			return ERROR + PRIORIDAD_VUELOS
		}
	}
	if e.Args[ACCION] == SIGUIENTE_VUELO {
		if len(e.Args) != 4 {
			return ERROR + SIGUIENTE_VUELO
		}
	}
	if e.Args[ACCION] == BORRAR {
		if len(e.Args) != 3 {
			return ERROR + BORRAR
		}
	}

	return TODO_OK
}

type ErrorURLArchivo struct {
	ErrorArchivo error
}

func (e ErrorURLArchivo) Error() string {
	if e.ErrorArchivo != nil {
		return ERROR + ARGEGAR_ARCHIVO
	}
	return TODO_OK
}

type ErrorCantidadInvalida struct {
	Cantidad int
}

func (e ErrorCantidadInvalida) Error() string {
	if e.Cantidad > 0 {
		return TODO_OK
	}
	return ERROR + VER_TABLERO
}

type ErrorCantidadInvalidaPrioridad struct {
	Cantidad int
}

func (e ErrorCantidadInvalidaPrioridad) Error() string {
	if e.Cantidad > 0 {
		return TODO_OK
	}
	return ERROR + PRIORIDAD_VUELOS
}

type ErrorModoInvalido struct {
	Modo string
}

func (e ErrorModoInvalido) Error() string {
	if e.Modo != ASC && e.Modo != DESC {
		return ERROR + VER_TABLERO
	}
	return TODO_OK
}

type ErrorFechaInvalida struct {
	Desde string
	Hasta string
}

func (e ErrorFechaInvalida) Error() string {
	comp := strings.Compare(e.Desde, e.Hasta)
	if comp > 1 {
		return ERROR + VER_TABLERO
	}
	return TODO_OK
}

type ErrorNoHayVuelo struct {
	HayVuelo bool
}

func (e ErrorNoHayVuelo) Error() string {
	if !e.HayVuelo {
		return ERROR + INFO_VUELO
	}
	return TODO_OK
}

type ErrorNoHayVuelosNuevos struct {
	NoHayVuelo bool
	Desde      string
	Hasta      string
	Fecha      string
}

func (e ErrorNoHayVuelosNuevos) Error() string {
	if e.NoHayVuelo {
		return "No hay vuelo registrado desde " + e.Desde + " hacia " + e.Hasta + " desde " + e.Fecha
	}
	return TODO_OK
}

type ErrorRangoFechasInvalida struct {
	Desde string
	Hasta string
}

func (e ErrorRangoFechasInvalida) Error() string {
	if strings.Compare(e.Desde, e.Hasta) > 1 {
		return ERROR + BORRAR
	}
	return TODO_OK
}
