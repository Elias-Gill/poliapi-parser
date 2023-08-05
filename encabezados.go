/*
   Aqui se encuentran las funciones necesarias para parsear y determinar la posicion de los
   encabezados
*/

package excelParser

import (
	"strings"
)

type encabezado struct {
	fila    int
	columna int
}

type encExamen struct {
	// dia
	filaFecha int
	colFecha  int
	// hora
	filaHora int
	colHora  int
}

// La fila de encabezados empieza donde se encuentre la columna con "Item"
func encontrarFilaEncabezados(cols [][]string) (encabezado, error) {
	var res encabezado
	for i, col := range cols {
		for k, row := range col {
			if strings.Contains(row, "Ítem") || strings.Contains(row, "Item") {
				res = encabezado{fila: k, columna: i}
				break
			}
		}
	}
	return res, nil
}

// guardar la posicion de los encabezados dentro de un map.
// Cada encabezado esta relacionado con una constante (int) de forma preventiva
func buscarEncabezados(cols [][]string, filaEncs encabezado) map[int]encabezado {
	resultado := make(map[int]encabezado)
	// recorrer coluna a columna
	for col := filaEncs.columna; col < len(cols); col++ {
		// ver a cual encabezado pertenece
		enc, ok := encabezados[cols[col][filaEncs.fila]]
		if !ok {
			continue
		}
		resultado[enc] = encabezado{fila: filaEncs.fila, columna: col}
	}
	return resultado
}

// los examenes cuentan con una estructura de encabezados ligeramente diferente, por lo que se parsean
// a parte.
func buscarEncsExamenes(cols [][]string, filaEncs encabezado) map[int]encExamen {
	resultado := make(map[int]encExamen)
	// recorrer columna a columna
	for col := filaEncs.columna; col < len(cols); col++ {
		// buscar en el encabezado superior a cual examen pertenece
		enc, ok := examenes[cols[col][filaEncs.fila-1]]
		if !ok {
			continue
		}

		// parsear los encabezados de dia y hora
		resultado[enc] = encExamen{
			filaHora:  filaEncs.fila,
			filaFecha: filaEncs.fila,
			colHora:   col,
			colFecha:  col + 1,
		}

		// aumentar en 1 la columna para evitar parsear de nuevo este encabezado
		col++
	}
	return resultado
}
