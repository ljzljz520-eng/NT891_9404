package draw

import "traveldeck/domain"

type Rotation struct{ Offset, Turns int }

func (r Rotation) Normalize(size int) Rotation {
	if size <= 0 {
		return Rotation{}
	}
	n := r.Offset % size
	if n < 0 {
		n += size
	}
	t := r.Turns % 4
	if t < 0 {
		t += 4
	}
	return Rotation{Offset: n, Turns: t}
}
func RotateDeck(deck domain.Deck, r Rotation) domain.Deck {
	n := r.Normalize(len(deck.Cards))
	cards := make([]string, len(deck.Cards))
	for i, c := range deck.Cards {
		cards[(i+n.Offset)%len(cards)] = c
	}
	deck.Cards = cards
	deck.Revision++
	return deck
}
func (r Rotation) Description() string {
	if r.Turns == 0 {
		return "steady"
	}
	if r.Turns > 0 {
		return "clockwise"
	}
	return "counterclockwise"
}
