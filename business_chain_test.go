package traveldeck

import (
	"testing"
	"traveldeck/catalog"
	"traveldeck/gesture"
	"traveldeck/service"
	"traveldeck/store"
)

func TestWorkflowOne(t *testing.T) {
	s, _ := store.Open(t.TempDir() + "/a")
	defer s.Close()
	c := catalog.New(catalog.Seed())
	cap, _ := gesture.New(false).Interpret(gesture.Event{Kind: gesture.Tap})
	v := service.New(c, s, cap)
	d, e := v.CreateDeck("w1", "mountain", nil)
	if e != nil {
		t.Fatal(e)
	}
	if _, e = v.Rotate(d.ID, 1); e != nil {
		t.Fatal(e)
	}
	if _, _, e = v.Draw(d.ID, 0, gesture.Event{Kind: gesture.Tap}); e != nil {
		t.Fatal(e)
	}
	if e = v.Return(d.ID, 0); e != nil {
		t.Fatal(e)
	}
}
func TestWorkflowTwo(t *testing.T) {
	s, _ := store.Open(t.TempDir() + "/b")
	defer s.Close()
	c := catalog.New(catalog.Seed())
	cap, _ := gesture.New(false).Interpret(gesture.Event{Kind: gesture.Tap})
	v := service.New(c, s, cap)
	d, _ := v.CreateDeck("w2", "coast", nil)
	if _, _, e := v.Draw(d.ID, 1, gesture.Event{Kind: gesture.Tap}); e != nil {
		t.Fatal(e)
	}
	h, e := v.History(d.ID)
	if e != nil || len(h) != 1 {
		t.Fatalf("%v %d", e, len(h))
	}
}
func TestWorkflowThree(t *testing.T) {
	s, _ := store.Open(t.TempDir() + "/c")
	defer s.Close()
	c := catalog.New(catalog.Seed())
	cap, _ := gesture.New(false).Interpret(gesture.Event{Kind: gesture.Tap})
	v := service.New(c, s, cap)
	d, _ := v.CreateDeck("w3", "city", nil)
	if _, _, e := v.Draw(d.ID, 2, gesture.Event{Kind: gesture.Tap}); e != nil {
		t.Fatal(e)
	}
}
func TestBusinessChain39(t *testing.T) {
	var typedNil *service.NullCapability
	var cap interface{ Ready() bool } = typedNil
	if service.CapabilityAvailable(cap) {
		t.Fatal("typed nil accepted")
	}
}
