package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"gitlab.com/idoko/foxtop/mozurl"
)

type Database struct {
	Conn *sql.DB
}

func Connect(dbfile string) (Database, error) {
	db := Database{}
	dsn := dbfile + "?mode=ro"
	conn, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return db, err
	}
	db.Conn = conn
	return db, nil
}

func (db Database) AllHosts() (mozurl.MozHostStore, error) {
	store := mozurl.MozHostStore{}
	query := `SELECT SUM(visit_count) as host_visit_count, rev_host as mozhost FROM moz_places GROUP BY rev_host ORDER BY host_visit_count DESC`
	rows, err := db.Conn.Query(query)
	if err != nil {
		return store, err
	}
	defer rows.Close()

	for rows.Next() {
		var visitCount int
		var revHost string

		err := rows.Scan(&visitCount, &revHost)
		if err != nil {
			return store, err
		}
		hostname := normalizeHost(revHost)
		host := mozurl.NewMozHost(hostname, visitCount)
		store.AddHost(host)
	}
	err = rows.Err()
	if err != nil {
		return store, err
	}
	return store, nil
}

func (db Database) URLsForHost(host *mozurl.MozHost) error {
	revhost := reverseHost(host.HostName())
	query := `SELECT visit_count, url FROM moz_places WHERE rev_host = $1 ORDER BY visit_count`
	rows, err := db.Conn.Query(query, revhost)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var visitCount int
		var url string
		if err := rows.Scan(&visitCount, &url); err != nil {
			return err
		}
		// we assume that the host already knows the visit count, so we go ahead to add the raw url alone
		host.AddRawURL(url)
	}
	return nil
}

func normalizeHost(revHost string) string {
	// strip the trailing full stop first
	n := len(revHost)
	revHost = revHost[:n-1]
	n = n - 1
	runes := make([]rune, n)
	for _, rune := range revHost {
		n--
		runes[n] = rune
	}
	return string(runes[n:])

}

// this is sort of the inverse of `normalizeHost`, used when we want to select URLs
// from the moz_places table
func reverseHost(host string) string {
	n := len(host)
	runes := make([]rune, n)
	for _, rune := range host {
		n--
		runes[n] = rune
	}
	str := string(runes[n:]) + "."
	return str
}
