package draw

import (
	"errors"
	"traveldeck/domain"
)

type Engine struct{ source domain.DestinationSource }

func New(source domain.DestinationSource) *Engine { return &Engine{source: source} }
func (e *Engine) BuildDeck(id, name string, groups []domain.Group) (domain.Deck, error) {
	if id == "" || name == "" {
		return domain.Deck{}, errors.New("deck identity required")
	}
	list, _ := e.source.List()
	cards := []string{}
	for _, d := range list {
		if len(groups) == 0 || containsGroup(groups, d.Group) {
			cards = append(cards, d.ID)
		}
	}
	if len(cards) == 0 {
		return domain.Deck{}, errors.New("no cards match")
	}
	return domain.Deck{ID: id, Name: name, Cards: cards, Revision: 1}, nil
}
func containsGroup(gs []domain.Group, g domain.Group) bool {
	for _, x := range gs {
		if x == g {
			return true
		}
	}
	return false
}
func (e *Engine) Inspect(deck domain.Deck, index int) (domain.Destination, error) {
	if index < 0 || index >= len(deck.Cards) {
		return domain.Destination{}, errors.New("index out of range")
	}
	return e.source.Get(deck.Cards[index])
}
