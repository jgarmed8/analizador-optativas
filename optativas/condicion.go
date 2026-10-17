package optativas

import (
	"errors"
	"math"
)

type TipoCondicion string

const (
	NotaMinima   TipoCondicion = "nota_minima"
	AprobarParte TipoCondicion = "aprobar_parte"
)

var (
	ErrNotaMinimaSinNota      = errors.New("una condición de nota mínima necesita una nota")
	ErrAprobarParteConNota    = errors.New("una condición de aprobar una parte no debe incluir una nota mínima")
	ErrTipoCondicionInvalido  = errors.New("tipo de condición no válido")
	ErrNotaMinimaFueraDeRango = errors.New("la nota mínima debe estar entre 0 y 10")
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
	if !tiposEvaluacionValidos[parte] {
		return CondicionMinima{}, ErrTipoEvaluacionInvalido
	}

	switch tipo {
	case NotaMinima:
		if notaMinima == nil {
			return CondicionMinima{}, ErrNotaMinimaSinNota
		}

		if math.IsNaN(*notaMinima) || *notaMinima < 0 || *notaMinima > 10 {
			return CondicionMinima{}, ErrNotaMinimaFueraDeRango
		}

		valor := *notaMinima
		notaMinima = &valor

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
