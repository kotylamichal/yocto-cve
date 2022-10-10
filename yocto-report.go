package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Issue struct {
	Id      string `json:"id"`
	Summary string `json:"sumary"`
	ScoreV2 string `json:"scorev2"`
	ScoreV3 string `json:"scorev3"`
	Vector  string `json:"vector"`
	Status  string `json:"status"`
	Link    string `json:"link"`
}

type Packages struct {
	Packages []Package `json:"package"`
}

type Package struct {
	Name    string  `json:"name"`
	Layer   string  `json:"layer"`
	Version string  `json:"version"`
	Issues  []Issue `json:"issue"`
}

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Use: " + os.Args[0] + " filename.json")
		return
	}

	jsonFile, err := os.Open(os.Args[1])

	fmt.Println("Yocto report generator")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	f, err := os.Create("report.md")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer f.Close()

	printMdHeader(f, 1, "Yocto CVE report")
	printMdLine(f, "TBD: General info about report")
	printMdEmptyLine(f)

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var packages Packages

	json.Unmarshal(byteValue, &packages)

	for i := 0; i < len(packages.Packages); i++ {
		var x = false

		for j := 0; j < len(packages.Packages[i].Issues); j++ {

			if packages.Packages[i].Issues[j].Status == "Unpatched" {
				if !x {
					printMdEmptyLine(f)
					printMdHeader(f, 2, packages.Packages[i].Name)
					printCVEHeader(f)

					x = true
				}

				printCVEItem(f, packages.Packages[i].Issues[j].Id, packages.Packages[i].Issues[j].ScoreV2,
					packages.Packages[i].Issues[j].ScoreV3, packages.Packages[i].Issues[j].Vector)

			}
		}
	}

	fmt.Println("DONE")

	defer jsonFile.Close()

}
