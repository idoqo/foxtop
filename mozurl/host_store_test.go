package mozurl

import "testing"

func TestReverseHost(t *testing.T) {
	rawurls := []struct {
		host    string
		revhost string
	}{
		{"github.com", "moc.buhtig"},
	}

	for _, u := range rawurls {
		reversed := reverseHost(u.host)
		if reversed != revhost {
			t.Errorf("expected %q to be %q when reversed, got %q", u.host, u.revhost, reversed)
		}
	}
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
