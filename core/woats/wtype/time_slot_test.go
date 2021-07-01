package wtype

import "testing"

func TestTimeSlot_IsCrossed(t *testing.T) {
	a := NewTimeSlot(1, 1, 8, 50, 9, 30)
	b := NewTimeSlot(2, 1, 8, 40, 9, 50)
	t.Log(a.IsCrossed(b))
	t.Log(b.IsCrossed(a))
}
