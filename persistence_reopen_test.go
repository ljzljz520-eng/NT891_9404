package traveldeck

import (
	"testing"
	"traveldeck/domain"
	"traveldeck/store"
)

func TestPersistenceSurvivesReopen(t *testing.T) {
	p := t.TempDir() + "/persist.db"
	s, e := store.Open(p)
	if e != nil {
		t.Fatal(e)
	}
	if e = s.SaveDeck(domain.Deck{ID: "persist", Name: "reopen", Cards: []string{"d001"}}); e != nil {
		t.Fatal(e)
	}
	s.Close()
	s, e = store.Open(p)
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	d, e := s.LoadDeck("persist")
	if e != nil || d.Name != "reopen" {
		t.Fatalf("%v", e)
	}
}
