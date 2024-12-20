package repository

import (
	"database/sql"
	"github.com/brunocordeiro180/go-url-shortener/utils"
	"log"
)

type Url struct {
	id  int
	url string
	key string
}

func GetOrCreateUrl(db *sql.DB, url string) (string, error) {
	key, err := GetKeyFromUrl(db, url)
	if err != nil {
		return "", err
	}
	if key != "" {
		return key, nil
	}
	key = utils.RandomId(6)
	err = CreateUrl(db, url, key)

	if err != nil {
		return "", err
	}
	return key, nil
}

func GetUrlFromKey(db *sql.DB, key string) (string, bool) {
	stmt, err := db.Prepare("select url from urls where key = $1")
	defer stmt.Close()

	if err != nil {
		log.Fatalln(err)
		return "", false
	}

	var url string
	_ = stmt.QueryRow(key).Scan(&url)
	return url, true
}

func CreateUrl(db *sql.DB, url, key string) error {
	println("vai salvar ", url, "com a chave ", key)
	stmt, err := db.Prepare("insert into urls (url, key) values ($1, $2)")
	defer stmt.Close()
	if err != nil {
		return err
	}

	_, err = stmt.Exec(url, key)

	if err != nil {
		return err
	}

	return nil
}

func GetKeyFromUrl(db *sql.DB, url string) (string, error) {
	stmt, err := db.Prepare("select key from urls where url = $1")
	if err != nil {
		return "", err
	}
	defer stmt.Close()
	var key string
	err = stmt.QueryRow(url).Scan(&key)
	if err != nil {
		return "", nil
	}

	return key, nil
}

func SelectAllUrls(db *sql.DB) ([]Url, error) {
	rows, err := db.Query("select id, url, key from urls")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var urls []Url
	for rows.Next() {
		var u Url
		err = rows.Scan(&u.id, &u.url, &u.key)
		if err != nil {
			return nil, err
		}
		urls = append(urls, u)
	}

	return urls, nil
}
