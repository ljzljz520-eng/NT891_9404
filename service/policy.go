package service

import "traveldeck/domain"

type Policy struct {
	MaxDeck     int
	AllowRepeat bool
}

func (p Policy) Accept(d domain.Deck) bool {
	if d.Closed {
		return false
	}
	if p.MaxDeck > 0 && len(d.Cards) > p.MaxDeck {
		return false
	}
	return len(d.Cards) > 0
}
func FilterByPreference(ds []domain.Destination, p domain.UserPreference) []domain.Destination {
	out := []domain.Destination{}
	for _, d := range ds {
		if p.Likes(d) {
			out = append(out, d)
		}
	}
	return out
}
