package excelParser

import (
	"github.com/xuri/excelize/v2"
	"testing"
)

var (
    file *excelize.File
    cols [][]string
)

func setupEnv(t *testing.T) {
    f, err := excelize.OpenFile("/home/elias/Documentos/go_proyects/excelParser/example.xlsx")
    if err != nil {
        t.Fatal(err.Error())
    }
    file = f

    cols, err = file.GetCols(file.GetSheetName(6))
    if err != nil {
        t.Fatal(err.Error())
    }
}

func TestBuscarLineaEncabezados(t *testing.T) {
    setupEnv(t)
    _, err := determFilaEncabezados(cols)
    if err != nil {
		t.Fatal(err.Error())
    }
}

func TestEncabezados(t *testing.T) {
	encs, _ := determFilaEncabezados(cols)
	m := parsearEncabezados(cols, encs)
	for _, v := range m {
		println(v.fila, v.columna)
	}
}

func TestEncsExamenes(t *testing.T) {
	encs, _ := determFilaEncabezados(cols)
	e := parsearEncsExamenes(cols, encs)
	for _, v := range e {
		println(v.colHora, v.filaHora, v.colFecha, v.filaFecha)
	}
}

func TestCargarMaterias(t *testing.T) {
    encs, _ := determFilaEncabezados(cols)
    e := cargarMaterias(cols, encs)
    for _, v := range e {
        println(v.Dias.Lunes)
    }
}
