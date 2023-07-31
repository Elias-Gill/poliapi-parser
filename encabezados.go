/*
    Aqui se encuentran las funciones necesarias para parsear y determinar la posicion de los
    encabezados
*/

package excelParser

import "strings"

type encabezado struct {
	fila    int
	columna int
}

func parsearEncabezados(cols [][]string) map[int]encabezado {
	var resultado map[int]encabezado
	// recorrer coluna a columna
	for col := filaEncs.columna; col < len(cols); col++ {
		// ver a cual encabezado pertenece
		enc, ok := encabezados[cols[col][filaEncs.fila]]
		if !ok {
			continue
		}
		resultado[enc] = encabezado{fila: filaEncs.fila, columna: col}
	}
	return nil
}

// La fila de encabezados empieza donde se encuentre la columna con "Item"
func determFilaEncabezados(cols [][]string) (*encabezado, error) {
	var res encabezado
	for i, col := range cols {
		for k, row := range col {
			if strings.Contains(row, "Ítem") || strings.Contains(row, "Item") {
				res = encabezado{fila: k, columna: i}
				break
			}
		}
	}
	return &res, nil
}

// los examenes cuentan con una estructura de encabezados ligeramente diferente, por lo que se parsean 
// a parte.
func parsearEncasExamenes() map[int]encabezado {
	// primero recorrer la fila de arriba, la cual contiene los encabezados de examenes
	return nil
}
