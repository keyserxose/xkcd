package main

import (
	"database/sql"
	"log"
)

func writeDB(rootPath, database, title, lastBuildDateFormatted string) {

	db, err := sql.Open("sqlite3", rootPath+"/"+database)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	_, err = db.Exec("insert into files (filename, executed, executionDate) VALUES(?, 1, ?)", title, lastBuildDateFormatted)
	if err != nil {
		log.Fatal(err)
	}
}
