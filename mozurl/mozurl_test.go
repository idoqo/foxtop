package mozurl

import (
	"testing"
)

func TestMozURL(t *testing.T) {
	t.Run("correctly extracts host from domain", func(t *testing.T) {
		cases := []struct {
			rawurl     string
			visitCount int
			host       string
		}{
			{rawurl: "https://t.co/lorem", visitCount: 4, host: "t.co"},
			{rawurl: "ftp://ftp.intel.com.br", visitCount: 4, host: "ftp.intel.com.br"},
			{rawurl: "about:config", visitCount: 4, host: "about:config"},
			{rawurl: "javascript:void();", visitCount: 4, host: "javascript:void();"},
			{rawurl: "place:parent=menu_____&parent=filed", visitCount: 4, host: "place:parent"},
		}
		for _, c := range cases {
			mu := NewMozURL(c.rawurl, c.visitCount)
			if mu.rawurl != c.rawurl || mu.visitCount != c.visitCount {
				t.Fatalf("failed to make MozUrl: Expected %v, got %v", c, mu)
			}
			got := mu.extractHost()
			if got != c.host {
				t.Errorf("expected host for %q to be %q, got %q", c.rawurl, c.host, got)
			}
		}
	})
}
