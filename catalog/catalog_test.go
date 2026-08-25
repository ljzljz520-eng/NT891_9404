package catalog

import "testing"

func TestCatalogSeedAndSearch(t *testing.T) {
	c := New(Seed())
	ds, e := c.List()
	if e != nil || len(ds) != 8 {
		t.Fatalf("catalog %v %d", e, len(ds))
	}
	if len(c.Search("lake")) != 1 {
		t.Fatal("search")
	}
}
