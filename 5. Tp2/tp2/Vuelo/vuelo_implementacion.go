package Vuelo

import (
	"strconv"

	TDADict "tdas/diccionario"
)

const (
	NUM_VUELO      = 0
	DESDE          = 2
	HASTA          = 3
	PRIORIDAD      = 5
	FECHA_DESPEGUE = 6
)

type vuelo struct {
	vueloDict TDADict.Diccionario[int, string]
}

func CrearVuelo(dict TDADict.Diccionario[int, string]) Vuelo {
	vuelo := new(vuelo)
	vuelo.vueloDict = dict
	return vuelo
}

func (v *vuelo) VerNum() string {
	return v.vueloDict.Obtener(NUM_VUELO)
}

func (v *vuelo) VerPrioridad() int {
	prioridad := v.vueloDict.Obtener(PRIORIDAD)
	prioridadNum, _ := strconv.Atoi(prioridad)
	return prioridadNum
}

func (v *vuelo) ImprimirVuelo() string {
	resultado := v.vueloDict.Obtener(0)
	for i := 1; v.vueloDict.Cantidad() > i; i++ {
		resultado += " " + v.vueloDict.Obtener(i)
	}
	return resultado
}

func (v *vuelo) FechaVuelo() string {
	return v.vueloDict.Obtener(FECHA_DESPEGUE)
}

func (v vuelo) ImprimirTablero() string {
	return v.vueloDict.Obtener(FECHA_DESPEGUE) + " - " + v.vueloDict.Obtener(NUM_VUELO)
}

func (v vuelo) VerViaje() string {
	return v.vueloDict.Obtener(DESDE) + " " + v.vueloDict.Obtener(HASTA)
}

func (v vuelo) ImprimirPrioridad() string {
	return v.vueloDict.Obtener(PRIORIDAD) + " - " + v.vueloDict.Obtener(NUM_VUELO)
}
