package db

import "testing"

func TestNormalizeHost(t *testing.T) {
	cases := []struct {
		rev    string
		actual string
	}{
		{"moc.buhtig.", "github.com"},
	}

	for _, c := range cases {
		hs := normalizeHost(c.rev)
		if hs != c.actual {
			t.Errorf("expected normal form of %q to be %q, got %q", c.rev, c.actual, hs)
		}
	}
}

func TestReverseHost(t *testing.T) {
	cases := []struct {
		host     string
		reversed string
	}{
		{"github.com", "moc.buhtig."},
	}

	for _, c := range cases {
		rev := reverseHost(c.host)
		if rev != c.reversed {
			t.Errorf("expected reversed form of %q to be %q, got %q", c.host, c.reversed, rev)
		}
	}

}
