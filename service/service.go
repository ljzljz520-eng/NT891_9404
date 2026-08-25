package service

import (
	"errors"
	"fmt"
	"traveldeck/domain"
	"traveldeck/draw"
	"traveldeck/gesture"
)

type Service struct {
	source  domain.DestinationSource
	repo    domain.DeckRepository
	engine  *draw.Engine
	actions domain.Capability
}

func New(source domain.DestinationSource, repo domain.DeckRepository, cap domain.Capability) *Service {
	return &Service{source: source, repo: repo, engine: draw.New(source), actions: cap}
}

// capabilityReady reports whether the configured dependency is available.
// It delegates to capabilityReady so the empty-implementation detection
// (nil interface or a non-nil interface wrapping a nil pointer, e.g. a
// typed-nil *NullCapability) lives in exactly one place.
func (s *Service) capabilityReady() bool {
	return capabilityReady(s.actions)
}

func (s *Service) CreateDeck(id, name string, groups []domain.Group) (domain.Deck, error) {
	if !s.capabilityReady() {
		return domain.Deck{}, errors.New("input capability unavailable")
	}
	d, e := s.engine.BuildDeck(id, name, groups)
	if e != nil {
		return d, e
	}
	return d, s.repo.SaveDeck(d)
}
func (s *Service) Rotate(id string, offset int) (domain.Deck, error) {
	if !s.capabilityReady() {
		return domain.Deck{}, errors.New("input capability unavailable")
	}
	d, e := s.repo.LoadDeck(id)
	if e != nil {
		return d, e
	}
	d = draw.RotateDeck(d, draw.Rotation{Offset: offset})
	return d, s.repo.SaveDeck(d)
}
func (s *Service) Draw(id string, index int, event gesture.Event) (domain.Destination, domain.DrawRecord, error) {
	d, e := s.repo.LoadDeck(id)
	if e != nil {
		return domain.Destination{}, domain.DrawRecord{}, e
	}
	if !s.capabilityReady() {
		return domain.Destination{}, domain.DrawRecord{}, errors.New("empty action")
	}
	dest, e := s.engine.Inspect(d, index)
	if e != nil {
		return dest, domain.DrawRecord{}, e
	}
	r := domain.DrawRecord{ID: fmt.Sprintf("%s-%d", id, index), DeckID: id, DestinationID: dest.ID, Sequence: index, Gesture: string(event.Kind)}
	return dest, r, s.repo.SaveDraw(r)
}
func (s *Service) Return(id string, seq int) error {
	if !s.capabilityReady() {
		return errors.New("input capability unavailable")
	}
	rs, e := s.repo.ListDraws(id)
	if e != nil {
		return e
	}
	for _, r := range rs {
		if r.Sequence == seq && !r.Returned {
			r.Returned = true
			return s.repo.SaveDraw(r)
		}
	}
	return errors.New("draw not returnable")
}
func (s *Service) History(id string) ([]domain.DrawRecord, error) { return s.repo.ListDraws(id) }
