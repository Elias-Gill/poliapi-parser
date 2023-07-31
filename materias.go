/*
    Aqui se encuentran las funciones necesarias para crear y cargar la informacion de las materias disponibles
*/

package excelParser

func contarMaterias(cols [][]string) int {
	for i := filaEncs.fila + 1; i < len(cols[filaEncs.columna]); i++ {
		// cuando encuentre el espacio en blanco significa que terminaron las materias
		if cols[filaEncs.columna][i] == "" {
			return i
		}
	}
	// INFO: jamas se deberia llegar a esta condicion
	return len(cols[filaEncs.columna])
}

/* for fila := filaEncs.fila + 1; fila < contarMaterias(cols); col++ {

} */
