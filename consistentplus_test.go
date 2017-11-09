package consistentplus

import "testing"

func TestConsistentHash(t *testing.T) {
  ring := NewRing()
  fooMember := NewMember("foo")
  barMember := NewMember("bar")
  ring.AddMember(fooMember)
  ring.AddMember(barMember)
  member = ring.FindMember("some object")
}
