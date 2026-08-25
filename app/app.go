package app

import (
	"fmt"
	"traveldeck/catalog"
	"traveldeck/domain"
	"traveldeck/gesture"
	"traveldeck/service"
	"traveldeck/store"
)

type App struct {
	Catalog *catalog.Catalog
	Store   *store.Store
	Service *service.Service
}

func New(path string) (*App, error) {
	st, e := store.Open(path)
	if e != nil {
		return nil, e
	}
	cat := catalog.New(catalog.Seed())
	cap, _ := gesture.New(true).Interpret(gesture.Event{Kind: gesture.Tap})
	return &App{Catalog: cat, Store: st, Service: service.New(cat, st, cap)}, nil
}
func (a *App) Demo() (string, error) {
	d, e := a.Service.CreateDeck("demo", "Weekend circle", nil)
	if e != nil {
		return "", e
	}
	if _, e = a.Service.Rotate(d.ID, 2); e != nil {
		return "", e
	}
	x, _, e := a.Service.Draw(d.ID, 0, gesture.Event{Kind: gesture.Tap})
	if e != nil {
		return "", e
	}
	return fmt.Sprintf("%s | %s | %s", x.Name, x.Highlight, x.BestSeason), nil
}
func DefaultGroups() []domain.Group {
	return []domain.Group{domain.Mountain, domain.Coastal, domain.City, domain.Countryside}
}
