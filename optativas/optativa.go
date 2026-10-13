package optativas

type Optativa struct {
	codigo      string
	nombre      string
	evaluacion  []ParteEvaluacion
	condiciones []CondicionMinima
}
