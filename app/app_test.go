package app

import "testing"

func TestDemo(t *testing.T) {
	a, e := New(t.TempDir() + "/d")
	if e != nil {
		t.Fatal(e)
	}
	defer a.Store.Close()
	if _, e = a.Demo(); e != nil {
		t.Fatal(e)
	}
}
