package catalog

import (
	"testing"
	"traveldeck/domain"
)

func TestActiveOnly(t *testing.T) {
	x := ActiveOnly([]domain.Destination{{ID: "a", Active: true}, {ID: "b"}})
	if len(x) != 1 {
		t.Fatal(len(x))
	}
}
