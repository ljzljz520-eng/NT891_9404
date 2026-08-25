package store

import (
	"testing"
	"traveldeck/domain"
)

func TestStoreRoundTrip(t *testing.T) {
	p := t.TempDir() + "/db"
	s, e := Open(p)
	if e != nil {
		t.Fatal(e)
	}
	d := domain.Deck{ID: "d", Name: "n", Cards: []string{"a"}}
	if e = s.SaveDeck(d); e != nil {
		t.Fatal(e)
	}
	s.Close()
	s, e = Open(p)
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	x, e := s.LoadDeck("d")
	if e != nil || x.Name != "n" {
		t.Fatalf("%v %+v", e, x)
	}
}
