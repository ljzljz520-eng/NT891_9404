package draw

import (
	"testing"
	"traveldeck/domain"
)

func TestRotate(t *testing.T) {
	d := domain.Deck{Cards: []string{"a", "b", "c"}}
	d = RotateDeck(d, Rotation{Offset: 1})
	if d.Cards[1] != "a" {
		t.Fatal(d.Cards)
	}
}
