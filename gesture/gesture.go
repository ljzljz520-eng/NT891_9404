package gesture

import "traveldeck/domain"

type Kind string

const (
	Swipe  Kind = "swipe"
	Tap    Kind = "tap"
	Camera Kind = "camera"
	Return Kind = "return"
)

type Event struct {
	Kind   Kind
	Delta  int
	Target string
}
type Interpreter struct{ fallback bool }

func New(fallback bool) Interpreter { return Interpreter{fallback: fallback} }
func (i Interpreter) Interpret(e Event) (domain.Capability, error) {
	switch e.Kind {
	case Swipe, Tap:
		return action{valid: e.Delta != 0 || e.Kind == Tap}, nil
	case Camera:
		return action{valid: !i.fallback}, nil
	case Return:
		return action{valid: e.Target != ""}, nil
	default:
		return action{valid: false}, nil
	}
}

type action struct{ valid bool }

func (a action) Ready() bool { return a.valid }
