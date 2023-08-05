package excelParser

// ----------------------- structs -------------------
type Dias struct {
	Lunes     string `json:"lunes"`
	Martes    string `json:"martes"`
	Miercoles string `json:"miercoles"`
	Jueves    string `json:"jueves"`
	Viernes   string `json:"viernes"`
	Sabado    string `json:"sabado"`
}

type Materia struct {
	// general
	Id         string `json:"id"`
	Asignatura string `json:"nombre"`
	Semestre   string `json:"semestre"`
	Seccion    string `json:"seccion"`
	Profesor   string `json:"profesor"`
	// examenes
	Parcial1 string `json:"parcial_1"`
	Parcial2 string `json:"parcial_2"`
	Final1   string `json:"final_1"`
	Final2   string `json:"final_2"`
	// horario de clase
	Dias Dias `json:"dias"`
}

type examen struct {
	fecha string
	hora  string
}

type Carrera struct {
	Asignaturas []Materia `json:"Asignaturas"`
	Nombre      string    `json:"nombre"`
	Index       int       `json:"index"`
}

const (
	asignatura = iota
	item
	dpto
	nivel
	seccion
	apellDocente
	nombreDocente
	correo
	dia
	hora
	lunes
	martes
	miercoles
	jueves
	viernes
	sabado
	final1
	final2
	parcial1
	parcial2
)

var (
	// mapeo de encabezados a constantes
	encabezados = map[string]int{
		"Asignatura": asignatura,
		"Item":       item,
		"DPTO":       dpto,
		"Grupo":      nivel,
		"Sección":    seccion,
		"Apellido":   apellDocente,
		"Nombre":     nombreDocente,
		"Correo":     correo,
		"Lunes":      lunes,
		"Martes":     martes,
		"Miércoles":  miercoles,
		"Jueves":     jueves,
		"Viernes":    viernes,
		"Sábado":     sabado,
	}

	examenes = map[string]int{
		"Día":          dia,
		"Hora":         hora,
		"1er. Parcial": parcial1,
		"2do. Parcial": parcial2,
		"1er. Final":   final1,
		"2do. Final":   final2,
	}
)
