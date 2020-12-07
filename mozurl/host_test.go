package mozurl

import "testing"

func TestAddRawURL(t *testing.T) {
	twitter := "https://twitter.com"
	h := NewMozHost(twitter, 0)
	url := "https://twitter.com/hello_world"
	h.AddRawURL(url, 1)

	wantedLen := 1
	gotLen := len(h.URLs)

	if gotLen != wantedLen {
		t.Errorf("expected number of URLs after adding new URL to be %d, got %d", wantedLen, gotLen)
	}

	wantedVisits := 1
	gotVisits := h.VisitCount
	if gotVisits != wantedVisits {
		t.Errorf("expected visit count of Host after adding new url to be %d, got %d", wantedVisits, gotVisits)
	}
}
