package optativas

import (
	"errors"
	"strings"
)

var ErrCodigoVacio = errors.New("la optativa tiene que tener un código")

type Optativa struct {
	codigo      string
	evaluacion  []ParteEvaluacion
	condiciones []CondicionMinima
}

func NuevaOptativa(codigo string, evaluacion []ParteEvaluacion, condiciones []CondicionMinima) (Optativa, error) {
	if strings.TrimSpace(codigo) == "" {
		return Optativa{}, ErrCodigoVacio
	}

	return Optativa{
		codigo:      codigo,
		evaluacion:  append([]ParteEvaluacion(nil), evaluacion...),
		condiciones: append([]CondicionMinima(nil), condiciones...),
	}, nil
}
