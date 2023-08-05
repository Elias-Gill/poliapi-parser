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
	f, err := excelize.OpenFile("/home/elias/Documentos/go/excelParser/example.xlsx")
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
	_, err := encontrarFilaEncabezados(cols)
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestEncabezados(t *testing.T) {
	encs, _ := encontrarFilaEncabezados(cols)
	m := buscarEncabezados(cols, encs)
	for _, v := range m {
		println(v.fila, v.columna)
	}
}

func TestEncsExamenes(t *testing.T) {
	encs, _ := encontrarFilaEncabezados(cols)
	e := buscarEncsExamenes(cols, encs)
	for _, v := range e {
		println(v.colHora, v.filaHora, v.colFecha, v.filaFecha)
	}
}

func TestCargarMaterias(t *testing.T) {
	encs, _ := encontrarFilaEncabezados(cols)
	e := parsearListaMaterias(cols, encs)
	for _, v := range e {
		println(v.Dias.Lunes)
	}
}

func TestIntegracion(t *testing.T) {
	_, err := ParsearArchivo("/home/elias/Documentos/go/excelParser/example.xlsx")
	if err != nil {
		t.Fatal(err.Error())
	}
}
