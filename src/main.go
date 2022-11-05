package main

import (
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Use: " + os.Args[0] + " filename.json")
		return
	}

	fmt.Println("Yocto report generator")

	err := DBCVE_FromJSON(os.Args[1], "../cve.sqlite")
	if err != 0 {
		fmt.Println("Error while converting JSON to SQLITE")
	} else {
		fmt.Println("JSON converted to SQLITE. Generating raport...")
	}

}
