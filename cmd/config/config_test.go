package config

import (
	"strings"
	"testing"
)

func TestDefaultDir(t *testing.T) {
	t.Run("correctly picks default profile if it's set", func(t *testing.T) {
		ini := strings.NewReader(`
[Profile1]
Name=default
Path=aaa.default
Default=1

[Profile0]
Name=default-nightly
Path=bbb.default
`)
		config, err := LoadConfig(ini)
		assertNoErr(t, err)
		want := "aaa.default"
		if config.DefaultPath() != want {
			t.Errorf("got %q as config dir, expected %q", config.DefaultPath(), want)
		}
	})

	t.Run("errors out if no default profile is available", func(t *testing.T) {
		ini := strings.NewReader(`
[Profile0]
Name=default-nightly
Path=bbb.default
`)
		_, err := LoadConfig(ini)
		assertErr(t, err, ErrNoDefaultProfile)
	})
}

func assertNoErr(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("not expecting any error but got: %v", err)
	}
}

func assertErr(t *testing.T, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("expecting error, got none")
	}

	if got != want {
		t.Errorf("expecting error: %q, got: %q", want.Error(), got.Error())
	}
}
