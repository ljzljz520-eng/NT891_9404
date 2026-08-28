package gesture

import "testing"

func TestFallbackTouch(t *testing.T) {
	a, _ := New(true).Interpret(Event{Kind: Tap})
	if !a.Ready() {
		t.Fatal("tap")
	}
	c, _ := New(true).Interpret(Event{Kind: Camera})
	if c.Ready() {
		t.Fatal("camera")
	}
}
