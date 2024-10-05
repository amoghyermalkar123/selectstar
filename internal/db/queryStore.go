package db

import (
	"database/sql"
	"fmt"
)

type QueryStore struct {
	db         *sql.DB
	queryStore []string
}

func NewQueryStore(db *sql.DB) *QueryStore {
	db.Exec("create table if not exists queries(name varchar(250), query text)")
	db.Exec("CREATE UNIQUE INDEX query_name_idx ON queries(name)")
	db.Exec("CREATE UNIQUE INDEX query_idx ON queries(query)")
	return &QueryStore{
		db:         db,
		queryStore: []string{},
	}
}

func (q *QueryStore) AddQuery(name, query string) error {
	q.queryStore = append(q.queryStore, name)
	_, err := q.db.Exec("insert into queries(name, query) values(?, ?)", name, query)
	if err != nil {
		return err
	}
	return nil
}

func (q *QueryStore) DeleteQuery(name string) {
	for idx, n := range q.queryStore {
		if n == name {
			q.queryStore = removeIndex(q.queryStore, idx)
		}
	}
	q.db.Exec("delete from queries where name=?", name)
}

func removeIndex(s []string, index int) []string {
	ret := make([]string, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func (q *QueryStore) GetQueries() []string {
	return q.queryStore
}

func (q *QueryStore) GetQuery(name string) (string, string, error) {
	var queryname, query string
	res := q.db.QueryRow("select name, query from queries where name=?", name)
	if err := res.Scan(&queryname, &query); err != nil {
		return "", "", err
	}
	return queryname, query, nil
}

func (q *QueryStore) LoadQueries() []string {
	rows, err := q.db.Query("select name from queries")
	if err != nil {
		fmt.Println("LoadQueries err:", err)
		return []string{}
	}

	var initialQueries []string
	for rows.Next() {
		var query string
		err := rows.Scan(&query)
		if err != nil {
			break
		}
		initialQueries = append(initialQueries, query)
	}
	q.queryStore = nil
	q.queryStore = append(q.queryStore, initialQueries...)
	return initialQueries
}
