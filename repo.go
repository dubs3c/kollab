package main

import (
	"database/sql"
	"encoding/json"
	"log"
)

func PathExists(db *sql.DB, path string) (bool, error) {
	found := 0
	err := db.QueryRow("SELECT id FROM paths WHERE url=?", path).Scan(&found)
	if found > 0 {
		return true, err
	}
	return false, err
}

func InsertPath(db *sql.DB, path *PathResponse) error {
	headers, err := json.Marshal(path.Headers)
	if err != nil {
		return err
	}
	_, err = db.Exec("INSERT INTO paths(uuid, url, verb, headers, fk_server) VALUES(?, ?, ?, ?, ?)", path.Id, path.Path, path.Verb, string(headers), 0)
	return err
}

func GetAllPaths(db *sql.DB) (*[]PathResponse, error) {
	paths := &[]PathResponse{}
	rows, err := db.Query("SELECT uuid, url, verb, headers, body FROM paths")
	if err != nil {
		return paths, err
	}

	defer rows.Close()

	for rows.Next() {
		p := PathResponse{}
		a := ""
		if err := rows.Scan(&p.Id, &p.Path, &p.Verb, &a, &p.Body); err != nil {
			log.Println("GetAllPaths:", err)
			continue
		}

		err = json.Unmarshal([]byte(a), &p.Headers)
		if err != nil {
			log.Println("GetAllPaths:", err)
		}

		*paths = append(*paths, p)
	}

	return paths, err
}

func GetPath(db *sql.DB, id string) (*PathResponse, error) {
	p := &PathResponse{}
	a := ""
	row := db.QueryRow("SELECT uuid, url, verb, headers, body FROM paths WHERE uuid=?", id)
	err := row.Scan(&p.Id, &p.Path, &p.Verb, &a, &p.Body)

	if err != nil {
		log.Println("GetPath:", err)
		return p, err
	}

	err = json.Unmarshal([]byte(a), &p.Headers)
	if err != nil {
		log.Println("GetPath JSON:", err)
	}

	return p, err
}

func DeletePath(db *sql.DB, id string) (int64, error) {
	result, err := db.Exec("DELETE FROM paths WHERE uuid=?", id)
	rows, _ := result.RowsAffected()
	return rows, err
}

func UpdatePath(db *sql.DB, path *PathResponse) (int64, error) {
	//row := db.QueryRow("SELECT uuid, url, verb, headers, body FROM paths WHERE url=?", path.Id)
	headers, err := json.Marshal(&path.Headers)
	if err != nil {
		return 0, err
	}
	result, err := db.Exec("UPDATE paths SET url=?, verb=?, headers=?, body=? WHERE uuid=?", &path.Path, &path.Verb, string(headers), &path.Body, &path.Id)
	rows, _ := result.RowsAffected()
	return rows, err
}

func InsertServer(db *sql.DB, path PathResponse) error {
	_, err := db.Exec("INSERT INTO servers(uuid, name, type, port) VALUES(?, ?, ?)", path.Id, path.Path, path.Verb, path.Headers, 0)
	return err
}
