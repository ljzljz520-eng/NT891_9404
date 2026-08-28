package draw

import (
	"testing"
	"traveldeck/catalog"
	"traveldeck/domain"
)

func TestBuildAndInspect(t *testing.T) {
	e := New(NewCatalog())
	d, x := e.BuildDeck("x", "trip", []domain.Group{domain.City})
	if x != nil || len(d.Cards) != 2 {
		t.Fatalf("%v", x)
	}
	if _, x = e.Inspect(d, 0); x != nil {
		t.Fatal(x)
	}
}
func NewCatalog() *catalog.Catalog { return catalog.New(catalog.Seed()) }
