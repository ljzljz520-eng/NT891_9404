package service

import (
	"errors"
	"testing"

	"traveldeck/domain"
	"traveldeck/gesture"
)

// memRepo is an in-memory DeckRepository used by the capability tests.
type memRepo struct {
	decks map[string]domain.Deck
	draws map[string]domain.DrawRecord
}

func newMemRepo() *memRepo {
	return &memRepo{decks: map[string]domain.Deck{}, draws: map[string]domain.DrawRecord{}}
}

func (m *memRepo) SaveDeck(d domain.Deck) error { m.decks[d.ID] = d; return nil }
func (m *memRepo) LoadDeck(id string) (domain.Deck, error) {
	d, ok := m.decks[id]
	if !ok {
		return domain.Deck{}, errors.New("deck not found")
	}
	return d, nil
}
func (m *memRepo) SaveDraw(r domain.DrawRecord) error { m.draws[r.ID] = r; return nil }
func (m *memRepo) ListDraws(deck string) ([]domain.DrawRecord, error) {
	out := []domain.DrawRecord{}
	for _, r := range m.draws {
		if deck == "" || r.DeckID == deck {
			out = append(out, r)
		}
	}
	return out, nil
}

func seedRepo() *memRepo {
	m := newMemRepo()
	m.decks["d"] = domain.Deck{ID: "d", Name: "n", Cards: []string{"a", "b"}}
	m.draws["d-0"] = domain.DrawRecord{ID: "d-0", DeckID: "d", Sequence: 0, DestinationID: "a"}
	return m
}

// TestCapabilityGuardsBlockEmptyImplementations verifies that update operations
// refuse to run when the dependency is an empty/unavailable implementation
// rather than treating it as a valid dependency and continuing.
func TestCapabilityGuardsBlockEmptyImplementations(t *testing.T) {
	cases := []struct {
		name string
		cap  domain.Capability
	}{
		{"nil interface", nil},
		// A typed-nil *NullCapability wrapped in a Capability interface is the
		// classic Go "empty implementation": cap == nil is false, so a naive
		// guard would dereference a nil pointer.
		{"typed-nil NullCapability", (*NullCapability)(nil)},
		{"disabled NullCapability", &NullCapability{Enabled: false}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := New(nil, seedRepo(), tc.cap)
			if _, err := s.Rotate("d", 1); err == nil {
				t.Fatalf("Rotate: expected error for %s, got nil", tc.name)
			}
			if err := s.Return("d", 0); err == nil {
				t.Fatalf("Return: expected error for %s, got nil", tc.name)
			}
		})
	}
}

// TestCapabilityNeverPanicsOnEmptyImplementation ensures the empty-implementation
// guard is safe (no nil-pointer dereference) across every entry point that
// consults the dependency.
func TestCapabilityNeverPanicsOnEmptyImplementation(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("operation panicked on empty implementation: %v", r)
		}
	}()
	var n *NullCapability
	var cap domain.Capability = n
	s := New(nil, seedRepo(), cap)
	_, _ = s.Rotate("d", 1)
	_ = s.Return("d", 0)
	if CapabilityAvailable(cap) {
		t.Fatal("CapabilityAvailable should report false for a typed-nil dependency")
	}
}

// TestOperationsProceedWithAvailableCapability confirms the guard does not
// over-block: when the dependency is genuinely available, update operations run.
func TestOperationsProceedWithAvailableCapability(t *testing.T) {
	repo := seedRepo()
	s := New(nil, repo, &NullCapability{Enabled: true})

	if _, err := s.Rotate("d", 1); err != nil {
		t.Fatalf("Rotate with available capability failed: %v", err)
	}
	if err := s.Return("d", 0); err != nil {
		t.Fatalf("Return with available capability failed: %v", err)
	}

	// A returned draw must not be returnable again, and must persist.
	if err := s.Return("d", 0); err == nil {
		t.Fatal("Return should reject an already-returned draw")
	}
}

// TestCreateDeckAndDrawRespectCapability confirms the pre-existing guards on
// CreateDeck and Draw remain safe and effective for empty implementations.
func TestCreateDeckAndDrawRespectCapability(t *testing.T) {
	var n *NullCapability
	s := New(nil, seedRepo(), n)

	if _, err := s.CreateDeck("d2", "n", nil); err == nil {
		t.Fatal("CreateDeck: expected error for empty implementation, got nil")
	}
	if _, _, err := s.Draw("d", 0, gesture.Event{Kind: gesture.Tap}); err == nil {
		t.Fatal("Draw: expected error for empty implementation, got nil")
	}
}
