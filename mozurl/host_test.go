package mozurl

import "testing"

func TestReverseHost(t *testing.T) {
	rawurls := []struct {
		host    string
		revhost string
	}{
		{"github.com", "moc.buhtig"},
		{"www.twitter.com", "moc.rettiwt.www"},
		{"twitter.com", "moc.rettiwt"},
		{"www.google.com", "moc.elgoog.www"},
	}

	for _, u := range rawurls {
		h := NewMozHost(u.host, 0)
		reversed := h.ReversedHost()
		if reversed != u.revhost {
			t.Errorf("expected %q to be %q when reversed, got %q", u.host, u.revhost, reversed)
		}
	}
}
