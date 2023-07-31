package excelParser

import (
	"log"
	"strconv"

	"github.com/xuri/excelize/v2"
)

func generarCarrera(sheet string, exf *excelize.File) Carrera {
    // traer las columnas
    cols, err := exf.GetCols(sheet)
    if err != nil {
        log.Fatal(err.Error())
    }
    // detemrinar donde comienza la linea de encabezados
    lineaEncabezados, err := determFilaEncabezados(cols)
    if err != nil {
        log.Print(err.Error())
    }
    // comenzar el parseo
    materias, err := parsearMaterias(exf, sheet, *lineaEncabezados)
    if err != nil {
        log.Print(err.Error())
    }
    // retornar la carrera
    return Carrera{
        Asignaturas: materias,
        Nombre:      sheet,
    }
}

// Parsea y devuelve las materias contenidas dentro de hoja de un archivo excel
func parsearMaterias(exf *excelize.File, sheet string, lineaEncabezados filaEncs) ([]*Materia, error) {
    // traer las filas
    rows, err := exf.GetRows(sheet)
    if err != nil {
        return nil, err
    }

    // crear una lista precardaga con el numero de materias existentes
    var res []*Materia
    for i := 0; i < lineaEncabezados.largo; i++ {
        res = append(res, &Materia{Dias: &Dias{}})
    }

    // recorrer las columnas de la fila de encabezados
    for col := lineaEncabezados.columna; col < len(rows[lineaEncabezados.fila])-1; col++ {
        // una cadena vacia significa el fin de los encabezados
        encabezado := rows[lineaEncabezados.fila][col]
        if encabezado == "" {
            return res, nil
        }

        // recorrer las filas para el encabezado actual
        i := -1
        for fila := lineaEncabezados.fila + 1; i < lineaEncabezados.largo-1 && fila < lineaEncabezados.largo-1; fila++ {
            i++
            println(strconv.Itoa(fila) + " " + strconv.Itoa(col))
            switch encabezados[encabezado] {
            case asignatura:
                res[i].Nombre = rows[fila][col]

            case nombreDocente:
                res[i].Profesor += rows[fila][col]

            case apellDocente:
                res[i].Profesor += rows[fila][col]

            case seccion:
                res[i].Seccion = rows[fila][col]

            case nivel:
                res[i].Semestre = rows[fila][col]

            case lunes:
                res[i].Dias.Lunes = rows[fila][col]

            case martes:
                res[i].Dias.Martes = rows[fila][col]

            case miercoles:
                res[i].Dias.Miercoles = rows[fila][col]

            case jueves:
                res[i].Dias.Jueves = rows[fila][col]

            case viernes:
                res[i].Dias.Viernes = rows[fila][col]

            case sabado:
                res[i].Dias.Sabado = rows[fila][col]

            //Cuando encontramos un encabezado de dia o de hora determinamos a cual examen corresponde
            case dia, hora:
                // seleccionar el encabezado de examen
                aux := rows[lineaEncabezados.fila-1][col]
                switch encabezados[aux] {
                case parcial1:
                    res[i].Parcial1 += rows[fila][col]
                case parcial2:
                    res[i].Parcial2 += rows[fila][col]
                case final1:
                    res[i].Final1 += rows[fila][col]
                case final2:
                    res[i].Final2 += rows[fila][col]
                }
            }
        }
    }

    // INFO: no deberia de llegarse a esta condicion
    return res, nil
}
