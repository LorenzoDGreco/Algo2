package main

import (
	err "Algueiza/errores"
	procesar "Algueiza/procesamientos"
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	ACCION         = 0
	INICIO         = 1
	ARCHIVO_VUELOS = 1
	BORRARINI      = 1
	VUELO          = 1
	NUM_VUELO      = 1
	CANTIDAD       = 1
	BORRARFIN      = 2
	DESTINO        = 2
	MODO           = 2
	FECHA          = 3
	DESDE          = 3
	HASTA          = 4
	TODO_OK        = "OK"
	VACIO          = ""
)

func main() {
	stdin := os.Stdin
	scan := bufio.NewScanner(stdin)
	ElementosProcesados := procesar.CrearProcesamientos()
	for scan.Scan() {
		linea := scan.Text()
		comandos := strings.Split(linea, " ")
		comandos = erorresParametros(comandos)

		switch comandos[ACCION] {

		case "agregar_archivo":
			errores := ElementosProcesados.ProcesarArchivoVuelos(comandos[ARCHIVO_VUELOS])
			imprimir(errores, VACIO, comandos[ACCION])

		case "ver_tablero":
			errores := ElementosProcesados.ImprimirTablero(comandos[CANTIDAD], comandos[MODO], comandos[DESDE], comandos[HASTA])
			imprimir(errores, VACIO, comandos[ACCION])

		case "info_vuelo":
			info_vuelo, errores := ElementosProcesados.VerVuelo(comandos[NUM_VUELO])
			imprimir(errores, info_vuelo, comandos[ACCION])

		case "prioridad_vuelos":
			errores := ElementosProcesados.PrioridadVuelos(comandos[CANTIDAD])
			imprimir(errores, VACIO, comandos[ACCION])

		case "siguiente_vuelo":
			siguiente_vuelo, errores := ElementosProcesados.SiguienteVuelo(comandos[INICIO], comandos[DESTINO], comandos[FECHA])
			imprimir(errores, siguiente_vuelo, comandos[ACCION])

		case "borrar":
			errores := ElementosProcesados.Borrar(comandos[BORRARINI], comandos[BORRARFIN])
			imprimir(errores, VACIO, comandos[ACCION])
		}
	}
}

func erorresParametros(comandos []string) []string {
	errores := err.ErrorParametros{Args: comandos}.Error()
	if errores != TODO_OK {
		comandos[ACCION] = VACIO
		fmt.Fprintf(os.Stderr, "%s\n", errores)
	}
	return comandos
}

func imprimir(errores, dato, comando string) {
	if errores == TODO_OK {
		if comando == "info_vuelo" || comando == "siguiente_vuelo" {
			fmt.Fprintf(os.Stdout, "%s\n", dato)
		}
		fmt.Fprintf(os.Stdout, "OK\n")
	} else {
		fmt.Fprintf(os.Stderr, "%s\n", errores)
	}
}
