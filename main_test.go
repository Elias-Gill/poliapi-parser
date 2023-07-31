package excelParser

import (
	"log"
    "testing"
	"github.com/xuri/excelize/v2"
)

func TestApp (t *testing.T){
    file, err := excelize.OpenFile("/home/elias/Documentos/go_proyects/excelParser/example.xlsx")
    if err != nil {
        log.Fatal(err)
    }
    /* for _, v := range data.GetSheetList() {
generarCarrera(v, data)
} */
    generarCarrera(file.GetSheetName(6), file)
}
