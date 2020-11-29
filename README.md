## foxtop
Foxtop is a (WIP) command-line tool to visualize Firefox (on Linux) history in the
terminal.

## How it works
By default, foxtop scans your Firefox directory (`~/home/.mozilla/firefox`) and
reads the `profile.ini` to figure out your default profile. It then reads the
data from the `places.sqlite` which includes data regarding your Firefox history
such as visit count per URLs, frecency (how the URLs get ranked in the URL bar),
etc.

Foxtop also opens the database file in read-only mode, ensuring that it doesn't
affect your browsing session.

## Installation
For now, simply clone the repository and build the `main.go` file in the `cmd`
directory. Or just issue `go run ./cmd/main.go`. Whatever works for you.

## Usage
Foxtop tries to figure your profile directory by reading
`~/home/.mozilla/firefox/profiles.ini` by default. If it fails to get the
profile right, or you want to visualize the history for a different profile, you
can pass the full path to the profile with the `-p` flag and foxtop will use
that instead. E.g:
```
$ go run ./cmd/main.go -p ~/.mozilla/firefox/glqkke4.default-nightly`
```
PS: A folder must contain a `places.sqlite` file to be considered a valid
profile path.

## License
Omo, [DWTFYWTPL](http://www.wtfpl.net/)
