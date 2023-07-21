package main

import (
	"log"
	"os"

	"github.com/xuri/excelize/v2"
)

func main() {
    file, err := os.Open("/home/elias/Documentos/go_proyects/excelParser/example.xlsx")
    if err != nil {
        log.Fatal(err.Error())
    }

    _, err = excelize.OpenReader(file)
    // cols, _ := data.GetCols(data.GetSheetName(6))
}
