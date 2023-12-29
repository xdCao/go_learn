package player

import (
	"testing"
)

func TestOps(t *testing.T) {
	mm := NewManager(make([]Music, 0))
	if mm == nil {
		t.Error("create manager failed")
	}
	if mm.Len() != 0 {
		t.Error("create manager failed, not empty")
	}

	m0 := &Music{Id: "1", Name: "my heart will go on", Artist: "dion", Location: "http://qbox.me", FileType: "mp3"}
	mm.Add(m0)

	if mm.Len() != 1 {
		t.Error("len wrong")
	}

	findm := mm.Find(m0.Name)
	if findm == nil {
		t.Error("not found by name")
	}

	getm, _ := mm.Get(0)
	if getm == nil {
		t.Error("not found by get")
	}

	removeM, _ := mm.Remove(0)
	if removeM == nil || mm.Len() != 0 {
		t.Error("remove failed")
	}

}
