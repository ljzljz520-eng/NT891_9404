package catalog

import (
	"errors"
	"sort"
	"traveldeck/domain"
)

type Catalog struct{ items map[string]domain.Destination }

func New(items []domain.Destination) *Catalog {
	c := &Catalog{items: map[string]domain.Destination{}}
	for _, d := range items {
		if d.Valid() {
			c.items[d.ID] = d
		}
	}
	return c
}
func (c *Catalog) List() ([]domain.Destination, error) {
	out := make([]domain.Destination, 0, len(c.items))
	for _, d := range c.items {
		out = append(out, d)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].ID < out[j].ID })
	return out, nil
}
func (c *Catalog) Get(id string) (domain.Destination, error) {
	d, ok := c.items[id]
	if !ok {
		return domain.Destination{}, errors.New("destination not found")
	}
	return d, nil
}
func (c *Catalog) ByGroup(g domain.Group) []domain.Destination {
	var out []domain.Destination
	for _, d := range c.items {
		if d.Group == g {
			out = append(out, d)
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Name < out[j].Name })
	return out
}
func (c *Catalog) Search(term string) []domain.Destination {
	var out []domain.Destination
	for _, d := range c.items {
		if contains(d.Name, term) || contains(d.Region, term) || contains(d.Highlight, term) {
			out = append(out, d)
		}
	}
	return out
}
func contains(s, t string) bool {
	if t == "" {
		return true
	}
	for i := 0; i+len(t) <= len(s); i++ {
		if equalFold(s[i:i+len(t)], t) {
			return true
		}
	}
	return false
}
func equalFold(a, b string) bool {
	for i := range a {
		x, y := a[i], b[i]
		if x >= 'A' && x <= 'Z' {
			x += 32
		}
		if y >= 'A' && y <= 'Z' {
			y += 32
		}
		if x != y {
			return false
		}
	}
	return true
}
