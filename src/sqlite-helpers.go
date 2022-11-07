package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func DBCVE_Create(filename string) int {

	os.Remove(filename)
	db, err := sql.Open("sqlite3", filename)

	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	sqlStmt, err := db.Prepare(`
		create table issues (
			id integer not null primary key,
			cve text,
			package text,
			scorev2 float,
			scorev3 float,
			vector text,
			source text);
	`)

	_, err = sqlStmt.Exec()
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return 1
	}

	return 0
}

func DBCVE_Insert(filename string, id int, cve string, pkg string, scorev2 string, scorev3 string, vector string, source string) int {

	db, err := sql.Open("sqlite3", filename)

	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	sqlStmt, err := db.Prepare("INSERT INTO issues VALUES (?,?,?,?,?,?,?)")

	_, err = sqlStmt.Exec(id, cve, pkg, scorev2, scorev3, vector, source)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return 1
	}

	return 0
}

func DBCVE_FromJSON(jsonname string, dbname string) int {

	jsonFile, err := os.Open(jsonname)

	DBCVE_Create(dbname)

	if err != nil {
		fmt.Println(err.Error())
		return 1
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var packages Packages

	json.Unmarshal(byteValue, &packages)

	var idCounter int

	for i := 0; i < len(packages.Packages); i++ {
		for j := 0; j < len(packages.Packages[i].Issues); j++ {
			if packages.Packages[i].Issues[j].Status == "Unpatched" {
				DBCVE_Insert(dbname, idCounter, packages.Packages[i].Issues[j].Id,
					packages.Packages[i].Name, packages.Packages[i].Issues[j].ScoreV2,
					packages.Packages[i].Issues[j].ScoreV3, packages.Packages[i].Issues[j].Vector,
					packages.Packages[i].Layer)
				idCounter = idCounter + 1
			}
		}
	}

	defer jsonFile.Close()

	return 0
}

func DBCVE_CountIssues(dbname string) int {

	db, err := sql.Open("sqlite3", dbname)

	if err != nil {
		fmt.Println(err)
		return -1
	}
	defer db.Close()

	var cnt int
	_ = db.QueryRow("SELECT count(*) from issues").Scan(&cnt)

	return cnt
}

func DBCVE_CountPackageIssues(dbname string, pkgname string) int {

	db, err := sql.Open("sqlite3", dbname)

	if err != nil {
		fmt.Println(err)
		return -1
	}
	defer db.Close()

	var cnt int
	_ = db.QueryRow("SELECT count(*) from issues WHERE source = " + "aaa").Scan(&cnt)

	return cnt
}

func DBCVE_PackagesList(dbname string) []string {

	db, err := sql.Open("sqlite3", dbname)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer db.Close()

	rows, err := db.Query("SELECT DISTINCT package from issues")
	fmt.Println("")

	var list []string
	for rows.Next() {
		var tmpname string
		rows.Scan(&tmpname)
		list = append(list, tmpname)
	}

	return list
}
