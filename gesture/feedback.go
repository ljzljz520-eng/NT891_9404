package gesture

type Feedback struct {
	Label      string
	Accessible bool
	Hint       string
}

func ForEvent(e Event) Feedback {
	switch e.Kind {
	case Swipe:
		return Feedback{"rotate", true, "swipe horizontally"}
	case Tap:
		return Feedback{"draw", true, "tap the center"}
	case Camera:
		return Feedback{"scan", false, "camera gesture"}
	default:
		return Feedback{"return", true, "place card back"}
	}
}
