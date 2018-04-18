package models

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type Database struct {
	*sql.DB
}

func NewDatabase(host string, port int, useTLS bool, user string, password string, dbName string) (db *Database, err error) {
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		user,
		password,
		host,
		port,
		dbName,
	)
	if useTLS == false {
		connStr += "?sslmode=disable"
	}
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err = conn.Ping(); err != nil {
		return nil, err
	}

	db = &Database{conn}
	return db, nil
}

func (db *Database) Close() {
	db.DB.Close()
}

func (db *Database) InsertSnippet(title, content string, expires time.Duration) (id int64, err error) {
	stmt := `INSERT INTO snippets(title, content, created, expires)
                 VALUES ($1, $2, $3, $4)
                 RETURNING id`
	now := time.Now()
	result, err := db.Exec(stmt,
		title,
		content,
		now,
		now.Add(expires),
	)
	if err != nil {
		return 0, err
	}

	id, err = result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (db *Database) GetSnippet(id int) (*Snippet, error) {
	stmt := `SELECT id, title, content, created, expires
                 FROM snippets
                 WHERE expires > CURRENT_TIMESTAMP AND id=$1`

	row := db.QueryRow(stmt, id)
	s := &Snippet{}
	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return s, nil
}

func (db *Database) LatestSnippets() (snippets Snippets, err error) {
	stmt := `SELECT id, title, content, created, expires
                 FROM snippets
                 WHERE expires > CURRENT_TIMESTAMP
                 ORDER BY created DESC
                 LIMIT 10`
	rows, err := db.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		s := &Snippet{}
		err := rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}
