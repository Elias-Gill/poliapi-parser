package main

import (
	"log"
	"os"
	"strings"

	"github.com/xuri/excelize/v2"
)

// Los nombres de los encabezados
const (
	asignatura = iota
	item
	departamento
	nivel
	seccion
	apellido
	nombre
	correo
	dia // dia y hora se van a usar para los examenes
	hora
	lunes
	martes
	miercoles
	jueves
	viernes
	sabado
	parcial1
	parcial2
	final1
	final2
)

type f struct {
	fila    int
	columna int
}

var (
	fila_encabezados f
	// map de la lista de encabezados ([nombre]id_encabezado) correspondiente
	encabezados = map[string]int{
		"Asignatura":   asignatura,
		"Item":         item,
		"DPTO":         departamento,
		"Grupo":        nivel,
		"Sección":      seccion,
		"Apellido":     apellido,
		"Nombre":       nombre,
		"Correo":       correo,
		"Día":          dia,
		"Hora":         hora,
		"Lunes":        lunes,
		"Martes":       martes,
		"Miércoles":    miercoles,
		"Jueves":       jueves,
		"Viernes":      viernes,
		"Sábado":       sabado,
		"1er. Parcial": parcial1,
		"2do. Parcial": parcial2,
		"1er. Final":   final1,
		"2do. Final":   final2,
	}
	// las posiciones de columna donde se ubican los encabezados ([id]columna)
	columnas_encabezados = map[int]int{}
)

func main() {
	file, err := os.Open("/home/elias/Documentos/go_proyects/excelParser/Planificacion-de-clases-y-examenes-Primer-Periodo-2023-version-web-22062023.xlsx")
	if err != nil {
		log.Fatal(err.Error())
	}

	data, err := excelize.OpenReader(file)
	ParsearArchivo(data)
}

// determina el inicio de la fila de encabezados
func determInicioEncabezados(data *excelize.File) {
	cols, _ := data.GetCols(data.GetSheetName(6))
	for i, col := range cols {
		for k, row := range col {
			if strings.Contains(row, "Ítem") || strings.Contains(row, "Item") {
				fila_encabezados = f{fila: k, columna: i}
				return
			}
		}
	}
}

// determina las columnas donde se alojan los encabezados
// anadir el numero de columna en la lista de encabezados
func determinarColumnasDeEncabezados(data *excelize.File) {
	cols, _ := data.GetCols(data.GetSheetName(1))
	// recorrer todas las columnas
	for i := fila_encabezados.columna; i < len(cols); i++ {
		col := cols[i]
		row := col[fila_encabezados.fila]
		// buscar el encabezado dentro de la lista
		encabezado, exists := encabezados[row]
		if exists {
			// Los encabezados de hora y dia se repiten para los examenes
			if encabezado == dia || encabezado == hora{
				// determinamos a cual examen pertenece
				// v, _ := encabezados[cols[i][fila_encabezados.fila-1]]

			} else { // para todos los encabezados se carga solo un numeor
				columnas_encabezados[encabezado] = i
			}
		}
	}
}

func armarMaterias(file *excelize.File, sheet string) []*Materia {
	rows, _ := file.GetRows(sheet)
	var res []*Materia
	for _, row := range rows {
		res = append(res,
			&Materia{
				Nombre:   row[columnas_encabezados[asignatura]],
				Id:       row[columnas_encabezados[item]],
				Semestre: row[columnas_encabezados[nivel]],
				Seccion:  row[columnas_encabezados[seccion]],
				Profesor: row[columnas_encabezados[nombre]] + row[columnas_encabezados[apellido]],
				Dias: Dias{
					Lunes:     row[columnas_encabezados[lunes]],
					Martes:    row[columnas_encabezados[martes]],
					Miercoles: row[columnas_encabezados[miercoles]],
					Jueves:    row[columnas_encabezados[jueves]],
					Viernes:   row[columnas_encabezados[viernes]],
					Sabado:    row[columnas_encabezados[sabado]],
				},
			})
	}
	return res
}

func ParsearArchivo(data *excelize.File) []*Materia {
	determInicioEncabezados(data)
	determinarColumnasDeEncabezados(data)
	// var res []Materia
	// for _, v := range data.GetSheetList() {
	// armarMaterias(data, data.GetSheetName(6))
	// }
	return armarMaterias(data, data.GetSheetName(6))
}

type Dias struct {
	Lunes     string
	Martes    string
	Miercoles string
	Jueves    string
	Viernes   string
	Sabado    string
}

type Materia struct {
	// general
	Id       string `json:"id"`
	Nombre   string `json:"nombre"`
	Semestre string `json:"semestre"`
	Seccion  string `json:"seccion"`
	Profesor string `json:"profesor"`
	// examenes
	Parcial1 string `json:"parcial_1"`
	Parcial2 string `json:"parcial_2"`
	Final1   string `json:"final_1"`
	Final2   string `json:"final_2"`
	// horario de clase
	Dias Dias `json:"dias"`
}

type Carrera struct {
	Asignaturas []*Materia `json:"Asignaturas"`
	Nombre      string     `json:"nombre"`
	Index       int        `json:"index"`
}
