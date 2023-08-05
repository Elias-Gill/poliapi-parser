package excelParser

import "github.com/xuri/excelize/v2"

// funcion principal de la libreria
func ParseFileFromName(file string) ([]Carrera, error) {
	f, err := excelize.OpenFile(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var res []Carrera
	sheets := f.GetSheetList()
	// parsear todas las hojas
	for i := 1; i < len(sheets); i++ {
		cols, err := f.GetCols(sheets[i])
		if err != nil {
			return nil, err
		}

		encs, err := encontrarFilaEncabezados(cols)
		if err != nil {
			return nil, err
		}

		materias := parsearListaMaterias(cols, encs)
		res = append(res, Carrera{
			Asignaturas: materias,
			Nombre:      sheets[i],
			Index:       i})
	}

	return res, nil
}

func ParseFileFromIo(file *excelize.File) ([]Carrera, error) {
    var res []Carrera
    sheets := file.GetSheetList()
    // parsear todas las hojas
    for i := 1; i < len(sheets); i++ {
        cols, err := file.GetCols(sheets[i])
        if err != nil {
            return nil, err
        }

        encs, err := encontrarFilaEncabezados(cols)
        if err != nil {
            return nil, err
        }

        materias := parsearListaMaterias(cols, encs)
        res = append(res, Carrera{
            Asignaturas: materias,
            Nombre:      sheets[i],
            Index:       i})
    }

    return res, nil
}
