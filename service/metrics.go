package service

import "traveldeck/domain"

func GroupCounts(ds []domain.Destination) map[domain.Group]int {
	out := map[domain.Group]int{}
	for _, d := range ds {
		out[d.Group]++
	}
	return out
}
func Seasonal(ds []domain.Destination, season string) []domain.Destination {
	out := []domain.Destination{}
	for _, d := range ds {
		if d.BestSeason == season {
			out = append(out, d)
		}
	}
	return out
}
func Explain(d domain.Destination) string {
	return d.Name + " is best in " + d.BestSeason + ": " + d.Highlight
}
