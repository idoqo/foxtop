package mozurl

type MozHostStore struct {
	mozhosts []*MozHost
}

func NewMozHostStore() *MozHostStore {
	return &MozHostStore{
		mozhosts: []*MozHost{},
	}
}

func (hs *MozHostStore) AddHost(host *MozHost) {
	hs.mozhosts = append(hs.mozhosts, host)
}

func (hs *MozHostStore) Hosts() []*MozHost {
	return hs.mozhosts
}
