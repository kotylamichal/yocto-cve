package main

import (
	"fmt"
	"os"
)

func printMdLine(filename *os.File, str string) {
	n, err := filename.WriteString(str)
	if err != nil {
		fmt.Printf("Error while trying to write %d bytes:\n", n)
		fmt.Println(err)
		fmt.Println("Exiting program...")
	}

}

func printMdLineln(filename *os.File, str string) {
	printMdLine(filename, str+"\n")
}

func printMdHeader(filename *os.File, level int, str string) {
	if level > 3 {
		fmt.Println("[WARNING]: Do not use header depth higher than 3")
		level = 3
	}
	for i := 1; i <= level; i++ {
		filename.WriteString("#")
	}

	//filename.WriteString(" " + str + "\n\n")
	printMdLineln(filename, " "+str+"\n")
}
func printMdEmptyLine(filename *os.File) {
	filename.WriteString("\n")
}

/* CVE table functions */
func printCVEHeader(filename *os.File) {
	printMdLineln(filename, "| CVEHello ID | Score V2 | Score V3 | Attack vector |")
	printMdLineln(filename, "|--------|----------|----------|---------------|")
}

func printCVEItem(filename *os.File, id string, s2 string, s3 string, vector string) {
	printMdLine(filename, "| "+id+" | ")
	printMdLine(filename, s2+" |")
	printMdLine(filename, s3+" |")
	printMdLine(filename, vector+" |\n")
}
