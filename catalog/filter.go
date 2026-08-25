package catalog

import "traveldeck/domain"

func ActiveOnly(ds []domain.Destination) []domain.Destination {
	out := []domain.Destination{}
	for _, d := range ds {
		if d.Active {
			out = append(out, d)
		}
	}
	return out
}
