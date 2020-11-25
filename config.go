package foxtop

import (
	"io"

	"gopkg.in/ini.v1"
)

const (
	ErrNoDefaultProfile = ConfigErr("no default profile was found in the provided config")
)

type ConfigErr string

func (ce ConfigErr) Error() string {
	return string(ce)
}

type Config struct {
	defaultName string
	defaultPath string
}

func (c *Config) DefaultPath() string {
	return c.defaultPath
}

func LoadConfig(config io.Reader) (*Config, error) {
	cfg, err := ini.Load(config)
	if err != nil {
		return nil, err
	}

	var name string
	var path string

	for _, section := range cfg.Sections() {
		if section.HasKey("Default") {
			if isDefault, _ := section.Key("Default").Int(); isDefault == 1 {
				name = section.Key("Name").String()
				path = section.Key("Path").String()
			}
		}
	}

	if name == "" || path == "" {
		return nil, ErrNoDefaultProfile
	}

	return &Config{
		defaultPath: path,
		defaultName: name,
	}, nil
}
