package grafo_test

import (
	"fmt"
	"testing"
	Grafo "tp3/grafo"

	"github.com/stretchr/testify/require"
)

func TestGrafo(t *testing.T) {
	grafo := Grafo.CrearGrafo[string, int](false, true)
	grafo.AgregarVertice("Argentina")
	grafo.AgregarVertice("Brasil")
	grafo.AgregarArista("Argentina", "Brasil", 500)
	require.EqualValues(t, 500, grafo.AristaPeso("Argentina", "Brasil"))
	require.EqualValues(t, 500, grafo.EliminarArista("Argentina", "Brasil"))
	require.PanicsWithValue(t, "No existe arista entre estos 2 vertices", func() { grafo.AristaPeso("Argentina", "Brasil") })
	require.PanicsWithValue(t, "No existe arista entre estos 2 vertices", func() { grafo.EliminarArista("Argentina", "Brasil") })
}

func TestGrafo2(t *testing.T) {
	grafo := Grafo.CrearGrafo[string, int](false, true)
	grafo.AgregarVertice("ATL")
	grafo.AgregarVertice("SHE")
	grafo.AgregarVertice("BAT")
	grafo.AgregarVertice("LAN")
	grafo.AgregarVertice("WAC")
	grafo.AgregarVertice("NAR")
	grafo.AgregarVertice("RIV")
	grafo.AgregarVertice("JFK")
	grafo.AgregarVertice("BH6")
	grafo.AgregarVertice("ASH")
	grafo.AgregarArista("JFK", "BH6", 344)
	grafo.AgregarArista("JFK", "ATL", 250)
	grafo.AgregarArista("JFK", "LAN", 459)
	grafo.AgregarArista("SHE", "ATL", 208)
	grafo.AgregarArista("RIV", "WAC", 329)
	grafo.AgregarArista("SHE", "RIV", 353)
	grafo.AgregarArista("NAR", "ATL", 164)
	grafo.AgregarArista("NAR", "BAT", 164)
	grafo.AgregarArista("WAC", "NAR", 463)
	grafo.AgregarArista("SHE", "NAR", 246)
	grafo.AgregarArista("BH6", "NAR", 196)
	grafo.AgregarArista("LAN", "NAR", 181)
	grafo.AgregarArista("ATL", "RIV", 492)
	grafo.AgregarArista("ASH", "BAT", 356)
	grafo.AgregarArista("WAC", "BAT", 348)
	grafo.AgregarArista("LAN", "ASH", 456)
	grafo.AgregarArista("BH6", "ASH", 322)
	grafo.AgregarArista("ASH", "LAN", 391)
	grafo.AgregarArista("LAN", "BAT", 181)
	grafo.AgregarArista("ATL", "WAC", 600)
	padres, dist := Grafo.CaminosMinimos(grafo, "WAC")
	fmt.Println("El padre de WAC es: " + padres.Obtener("WAC"))
	fmt.Println(dist.Obtener("WAC"))
	fmt.Println("El padre de NAR es: " + padres.Obtener("NAR"))
	fmt.Println(dist.Obtener("NAR"))
	fmt.Println("El padre de ATL es: " + padres.Obtener("ATL"))
	fmt.Println(dist.Obtener("ATL"))
}
