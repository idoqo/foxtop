package mozurl

type MozHostStore struct {
	MozHosts []*MozHost `json:"hosts"`
}

func NewMozHostStore() *MozHostStore {
	return &MozHostStore{
		MozHosts: []*MozHost{},
	}
}

func (hs *MozHostStore) AddHost(host *MozHost) {
	hs.MozHosts = append(hs.MozHosts, host)
}

func (hs *MozHostStore) Hosts() []*MozHost {
	return hs.MozHosts
}
