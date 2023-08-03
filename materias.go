/*
Aqui se encuentran las funciones necesarias para crear y cargar la informacion de las materias disponibles
*/
package excelParser

func contarMaterias(cols [][]string, filaEncs encabezado) int {
	for i := filaEncs.fila + 1; i < len(cols[filaEncs.columna]); i++ {
		// cuando encuentre el espacio en blanco significa que terminaron las materias
		if cols[filaEncs.columna][i] == "" {
			return i
		}
	}
	// INFO: jamas se deberia llegar a esta condicion
	return len(cols[filaEncs.columna])
}

func cargarMaterias(cols [][]string, filaEncs encabezado) []Materia {
	encs := parsearEncabezados(cols, filaEncs)
	exms := parsearEncsExamenes(cols, filaEncs)

	var res []Materia
	for i := filaEncs.fila + 1; i < contarMaterias(cols, filaEncs); i++ {
		res = append(res, Materia{
			// info general de la carrera
			Profesor:   cols[encs[nombreDocente].columna][i] + " " + cols[encs[apellDocente].columna][i],
			Asignatura: cols[encs[asignatura].columna][i],
			Semestre:   cols[encs[nivel].columna][i],
			Seccion:    cols[encs[seccion].columna][i],

			// --- examenes ---
            // parciales
			Parcial1: cols[exms[parcial1].colFecha][i] + 
            "  " + cols[exms[parcial1].colHora][i],
			Parcial2: cols[exms[parcial2].colFecha][i] + 
            "  " + cols[exms[parcial1].colHora][i],

            // finales
			Final1:   cols[exms[final1].colFecha][i] + 
            "  " + cols[exms[final1].colHora][i],
			Final2:   cols[exms[final2].colFecha][i] + 
            "  " + cols[exms[final2].colHora][i],

			// --- dias de clase ---
			Dias: Dias{
				Lunes:     cols[encs[lunes].columna][i],
				Martes:    cols[encs[martes].columna][i],
				Miercoles: cols[encs[miercoles].columna][i],
				Jueves:    cols[encs[jueves].columna][i],
				Viernes:   cols[encs[viernes].columna][i],
				Sabado:    cols[encs[sabado].columna][i],
			},
		})
	}
	return res
}
