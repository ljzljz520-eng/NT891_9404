package domain

type Group string

const (
	Mountain    Group = "mountain"
	Coastal     Group = "coastal"
	City        Group = "city"
	Countryside Group = "countryside"
)

type Destination struct {
	ID, Name, Region, Highlight, BestSeason string
	Group                                   Group
	Active                                  bool
}
type Deck struct {
	ID, Name string
	Cards    []string
	Revision int
	Closed   bool
}
type DrawRecord struct {
	ID, DeckID, DestinationID string
	Sequence                  int
	Returned                  bool
	Gesture                   string
}
type UserPreference struct {
	ID         string
	Groups     []Group
	Seasons    []string
	HomeRegion string
}
type EntityTypeOne = Destination
type EntityTypeTwo = Deck
type EntityTypeThree = DrawRecord
type EntityTypeFour = UserPreference

func (d Destination) Valid() bool {
	return d.ID != "" && d.Name != "" && d.Highlight != "" && d.BestSeason != "" && d.Active
}
func (d Deck) Contains(id string) bool {
	for _, c := range d.Cards {
		if c == id {
			return true
		}
	}
	return false
}
func (d Deck) Size() int             { return len(d.Cards) }
func (r DrawRecord) CanReturn() bool { return !r.Returned && r.DestinationID != "" }
func (p UserPreference) Likes(d Destination) bool {
	if len(p.Groups) == 0 {
		return true
	}
	for _, g := range p.Groups {
		if g == d.Group {
			return true
		}
	}
	return false
}
