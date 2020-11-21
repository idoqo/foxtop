package mozurl

import (
	"fmt"
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
			mu := NewMozUrl(c.rawurl, c.visitCount)
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

func TestMozDomains(t *testing.T) {
	t.Run("puts a MozUrl in the correct domain", func(t *testing.T) {
		expectedHost := "t.co"
		url := "https://t.co/lorem"

		tco := NewMozUrl(url, 4)
		dm := NewDomain(tco)

		if dm.Host() != expectedHost {
			t.Fatalf("expected domain host to be %q, got %q", expectedHost, dm.Host())
		}

		assertDomainURLsCount(t, 1, dm.CountURLs())
		assertDomainVisitCount(t, dm, 4)
		assertDomainUrlAtIndex(t, dm, tco.rawurl, 0)
	})

	t.Run("appends a URL to an existing domain value", func(t *testing.T) {
		baseHost := "t.co"
		url1 := NewMozUrl(
			fmt.Sprintf("https://%s/lorem", baseHost),
			4,
		)
		url2 := NewMozUrl(
			fmt.Sprintf("https://%s/ipsum", baseHost),
			2,
		)

		dm := NewDomain(url1)
		dm.addUrl(url2)

		assertDomainURLsCount(t, 2, dm.CountURLs())
		assertDomainVisitCount(t, dm, 6)
		assertDomainUrlAtIndex(t, dm, url2.rawurl, 1)
	})

	t.Run("errors out when adding url with different host", func(t *testing.T) {
		baseHost := "t.co"
		url1 := NewMozUrl(
			fmt.Sprintf("https://%s/lorem", baseHost),
			4,
		)
		url2 := NewMozUrl(
			"https://goo.gl/ipsum",
			2,
		)

		dm := NewDomain(url1)
		err := dm.addUrl(url2)
		if err == nil {
			t.Error("expected addUrl to reject URL with different host but got no error")
		}
	})

}

func TestDomainStoreFromURLs(t *testing.T) {
	raw := map[string]int{
		"https://t.co/lorem":                 4,
		"https://twitter.com/jola_adebayor/": 2,
		"https://twitter.com/theshalvah/":    1,
		"https://mchl.xyz":                   3,
	}
	mozUrls := []MozUrl{}
	for u, c := range raw {
		url := NewMozUrl(u, c)
		mozUrls = append(mozUrls, url)
	}
	store := NewDomainStore()
	store.PackFromURLs(mozUrls)

	if store.Domains == nil {
		t.Fatal("did not expect domains map in store to be nil")
	}

	if len(store.Domains) != 3 {
		t.Errorf("expected number of domains in store to be %d, got %d", 3, len(store.Domains))
	}

	host1 := "twitter.com"
	twitter, exists := store.Domains[host1]
	if !exists {
		t.Errorf("expected to find host %q in domain store, but couldn't", host1)
	}

	if twitter.CountURLs() != 2 {
		t.Errorf("expected URL count for domain: %q to be %d, got %d", host1, 2, twitter.CountURLs())
	}

	if twitter.VisitCount() != 3 {
		t.Errorf("expected visit count for %v to be %d, got %d", twitter, 3, twitter.VisitCount())
	}
}

func TestDomainSortMethods(t *testing.T) {
	t.Run("correctly returns the number of domains", func(t *testing.T) {
		raw := map[string]int{
			"https://t.co/lorem":                 4,
			"https://twitter.com/jola_adebayor/": 2,
			"https://twitter.com/theshalvah/":    1,
		}
		mozUrls := []MozUrl{}
		for u, c := range raw {
			url := NewMozUrl(u, c)
			mozUrls = append(mozUrls, url)
		}
		store := NewDomainStore()
		store.PackFromURLs(mozUrls)
		want := 2

		if store.Len() != want {
			t.Errorf("expected number of domains in store to be %d, got %d", want, store.Len())
		}
	})

	t.Run("correctly reports higher visit count as greater than", func(t *testing.T) {
		raw := map[string]int{
			"https://t.co/lorem":                 4,
			"https://twitter.com/jola_adebayor/": 2,
			"https://twitter.com/theshalvah/":    1,
		}
		mozUrls := []MozUrl{}
		for u, c := range raw {
			url := NewMozUrl(u, c)
			mozUrls = append(mozUrls, url)
		}

		store := NewDomainStore()
		store.PackFromURLs(mozUrls)

		tcoIndex := 0     //visitCount should be 4
		twitterIndex := 1 //visitCount should be 3
		if !store.Less(tcoIndex, twitterIndex) {
			t.Errorf(
				"expected %q with visitCount: %d to be before (>) %q with visitCount: %d",
				store.hosts[tcoIndex].Host(),
				store.hosts[tcoIndex].VisitCount(),
				store.hosts[twitterIndex].Host(),
				store.hosts[twitterIndex].VisitCount(),
			)
		}
	})

	t.Run("correctly swaps two domains in the domain store", func(t *testing.T) {
		raw := map[string]int{
			"https://t.co/lorem":                 4,
			"https://twitter.com/jola_adebayor/": 2,
			"https://twitter.com/theshalvah/":    1,
		}
		mozUrls := []MozUrl{}
		for u, c := range raw {
			url := NewMozUrl(u, c)
			mozUrls = append(mozUrls, url)
		}
		store := NewDomainStore()
		store.PackFromURLs(mozUrls)

		tco := "t.co"
		twitter := "twitter.com"
		if store.hosts[0].Host() != tco {
			t.Fatalf("expected domain at index 0 to be %q, got %q", tco, store.hosts[0])
		}
		store.Swap(0, 1)
		if store.hosts[0].Host() != twitter {
			t.Errorf("expected domain at index 0 to be %q after swapping, got %q", twitter, store.hosts[0].Host())
		}
	})
}

func assertDomainURLsCount(t *testing.T, expected, actual int) {
	t.Helper()
	if expected != actual {
		t.Fatalf("expected %d URLs in domain, got %d", expected, actual)
	}
}

func assertDomainVisitCount(t *testing.T, domain *Domain, expectedVisits int) {
	t.Helper()
	if expectedVisits != domain.VisitCount() {
		t.Errorf("expected %q to have %d visit counts, got %d", domain.Host(), expectedVisits, domain.VisitCount())
	}
}

func assertDomainUrlAtIndex(t *testing.T, dm *Domain, targetUrl string, targetIndex int) {
	t.Helper()
	mozUrl := dm.URLs()[targetIndex]
	if mozUrl.rawurl != targetUrl {
		t.Errorf("expected %q at index %d in %q's URLs, got %q", targetUrl, targetIndex, dm.Host(), mozUrl.rawurl)
	}
}
