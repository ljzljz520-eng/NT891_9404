package domain

type DestinationSource interface {
	List() ([]Destination, error)
	Get(string) (Destination, error)
}
type DeckRepository interface {
	SaveDeck(Deck) error
	LoadDeck(string) (Deck, error)
	SaveDraw(DrawRecord) error
	ListDraws(string) ([]DrawRecord, error)
}
type Capability interface{ Ready() bool }

func NormalizeGroup(g string) Group {
	switch g {
	case "mountain":
		return Mountain
	case "coastal":
		return Coastal
	case "city":
		return City
	case "countryside":
		return Countryside
	default:
		return Group(g)
	}
}
