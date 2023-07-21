package main

import (
	"strings"

	"github.com/xuri/excelize/v2"
)

// Buscamos el inicio de la fila de encabezados, lo que buscamos es la columna con "Item"
func determFilaEncabezados(cols [][]string) {
	for i, col := range cols {
		for k, row := range col {
			if strings.Contains(row, "Ítem") || strings.Contains(row, "Item") {
				filaEncabezados = filaEncs{fila: k, columna: i}
			} 
		}
	}
}

func parsear(sheet string, data *excelize.File) error {
    rows, err := data.GetRows(sheet)
    if err != nil {
        return err
    }

    for i := 0; i < filaEncabezados.largo; i++ {
        
    }

    return nil
}

func resultado(data *excelize.File) { 

}
