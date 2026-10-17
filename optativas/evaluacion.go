package optativas

import "errors"

type TipoEvaluacion string

const (
	Teoria        TipoEvaluacion = "teoria"
	Practicas     TipoEvaluacion = "practicas"
	Trabajos      TipoEvaluacion = "trabajos"
	Entregables   TipoEvaluacion = "entregables"
	Participacion TipoEvaluacion = "participacion"
	Examen        TipoEvaluacion = "examen"
)

var tiposEvaluacionValidos = map[TipoEvaluacion]bool{
	Teoria:        true,
	Practicas:     true,
	Trabajos:      true,
	Entregables:   true,
	Participacion: true,
	Examen:        true,
}

var (
	ErrPorcentajeInvalido     = errors.New("el porcentaje debe estar entre 0 y 100")
	ErrTipoEvaluacionInvalido = errors.New("el tipo de evaluación no está definido")
)

type ParteEvaluacion struct {
	tipo       TipoEvaluacion
	porcentaje float64
}

func NuevaParteEvaluacion(tipo TipoEvaluacion, porcentaje float64) (ParteEvaluacion, error) {
	if !tiposEvaluacionValidos[tipo] {
		return ParteEvaluacion{}, ErrTipoEvaluacionInvalido
	}

	if porcentaje < 0 || porcentaje > 100 {
		return ParteEvaluacion{}, ErrPorcentajeInvalido
	}

	return ParteEvaluacion{
		tipo:       tipo,
		porcentaje: porcentaje,
	}, nil
}
