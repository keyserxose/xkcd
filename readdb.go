package main

import (
	"database/sql"
)

func readDb(rootPath, database, title, lastBuildDateFormatted string) (filename, executionDate string) {

	db, err := sql.Open("sqlite3", rootPath+"/"+database)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	rows, err := db.Query("select filename from files where executed = 1 and filename = ?", title)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {

		err = rows.Scan(&filename)
		if err != nil {
			panic(err)
		}
	}

	rows2, err := db.Query("select executionDate from files where executed = 1 and executionDate = ?", lastBuildDateFormatted)
	if err != nil {
		panic(err)
	}

	defer rows2.Close()

	for rows2.Next() {

		err = rows2.Scan(&executionDate)
		if err != nil {
			panic(err)
		}
	}

	return filename, executionDate

}
