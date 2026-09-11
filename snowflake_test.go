package snowflake

import "testing"

func TestUniqueIDs(t *testing.T) {
	g := NewGenerator(1)
	id1 := g.NextID()
	id2 := g.NextID()
	if id1 == id2 {
		t.Errorf("expected unique IDs, got duplicate: %d", id1)
	}
}
