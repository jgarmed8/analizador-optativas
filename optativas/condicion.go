package optativas

import "errors"

type TipoCondicion string

const (
	NotaMinima   TipoCondicion = "nota_minima"
	AprobarParte TipoCondicion = "aprobar_parte"
)

var (
	ErrNotaMinimaSinNota     = errors.New("una condición de nota mínima necesita una nota")
	ErrAprobarParteConNota   = errors.New("una condición de aprobar una parte no debe incluir una nota mínima")
	ErrTipoCondicionInvalido = errors.New("tipo de condición no válido")
)

type CondicionMinima struct {
	parte      TipoEvaluacion
	tipo       TipoCondicion
	notaMinima *float64
}

func NuevaCondicionMinima(
	parte TipoEvaluacion,
	tipo TipoCondicion,
	notaMinima *float64,
) (CondicionMinima, error) {
	switch tipo {
	case NotaMinima:
		if notaMinima == nil {
			return CondicionMinima{}, ErrNotaMinimaSinNota
		}
	case AprobarParte:
		if notaMinima != nil {
			return CondicionMinima{}, ErrAprobarParteConNota
		}
	default:
		return CondicionMinima{}, ErrTipoCondicionInvalido
	}

	return CondicionMinima{
		parte:      parte,
		tipo:       tipo,
		notaMinima: notaMinima,
	}, nil
}
