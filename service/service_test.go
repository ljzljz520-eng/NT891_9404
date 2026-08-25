package service

import (
	"testing"
	"traveldeck/catalog"
	"traveldeck/gesture"
	"traveldeck/store"
)

func TestServiceFlow(t *testing.T) {
	p := t.TempDir() + "/x.db"
	s, e := store.Open(p)
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	c := catalog.New(catalog.Seed())
	cap, _ := gesture.New(false).Interpret(gesture.Event{Kind: gesture.Tap})
	v := New(c, s, cap)
	d, e := v.CreateDeck("d", "demo", nil)
	if e != nil {
		t.Fatal(e)
	}
	if _, _, e = v.Draw(d.ID, 0, gesture.Event{Kind: gesture.Tap}); e != nil {
		t.Fatal(e)
	}
}
