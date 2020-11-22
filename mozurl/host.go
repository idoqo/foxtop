package mozurl

// Type MozHost wraps all URLs that share the same host
type MozHost struct {
	host         string
	urls         []MozURL
	visitCount   int
	reversedHost string
}

func NewMozHost(host string, visitCount int) *MozHost {
	return &MozHost{
		host:       host,
		visitCount: visitCount,
	}
}

func (mh *MozHost) ReversedHost() string {
	if mh.reversedHost != "" {
		return mh.reversedHost
	}
	n := len(mh.host)
	runes := make([]rune, n)
	for _, rune := range mh.host {
		n--
		runes[n] = rune
	}
	return string(runes[n:])
}
