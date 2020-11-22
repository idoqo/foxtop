package mozurl

import "testing"

func TestAppendHost(t *testing.T) {
	store := NewMozHostStore()

	tco := "https://t.co"
	host := NewMozHost(tco, 0)
	store.AddHost(host)

	want := 1
	got := len(store.Hosts())
	if got != 1 {
		t.Errorf("expected host count in store to be %d after adding host, got %d", want, got)
	}
}
