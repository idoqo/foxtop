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

/*func TestMozDomains(t *testing.T) {
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
}*/
