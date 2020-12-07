package mozurl

import (
	"net/url"
	"strings"
)

// MozURL is a representation of a row from `moz_places` but containing only the data we need
type MozURL struct {
	RawUrl     string `json:"raw_url"`
	visitCount int `json:"visit_count"`
}

func NewMozURL(rawurl string, visitCount int) MozURL {
	return MozURL{
		RawUrl:     rawurl,
		visitCount: visitCount,
	}
}

func (mu MozURL) extractHost() string {
	u, _ := url.Parse(mu.RawUrl)
	switch u.Scheme {
	case "about", "javascript":
		return u.Scheme + ":" + u.Opaque
	case "file":
		return u.Scheme + "://" + u.Host
	case "place":
		return parsePlacesUrl(u.Opaque)
	}
	return u.Host
}

func parsePlacesUrl(place string) string {
	return "place" + ":" + strings.Split(place, "=")[0]
}

func (mu MozURL) GetProtocol() string {
	u, _ := url.Parse(mu.RawUrl)
	switch u.Scheme {
	case "https", "http", "ftp", "file":
		return u.Scheme
	default:
		return ""
	}
}
