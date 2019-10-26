package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// filterLines print a definide part os text
func filterLines(scanning *bufio.Scanner, start int, end int) {
	actualLine := 1
	for scanning.Scan() {
		if actualLine >= start && actualLine <= end {
			fmt.Println(scanning.Text())
		} else if actualLine >= start && end == 0 {
			fmt.Println(scanning.Text())
		}
		actualLine++
	}
	if err := scanning.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Reading Standard Input:", err)
	}
}

// Show help menu
func help() {
	fmt.Printf("Usage: flines [OPTION]... [FILE]...\n")
	fmt.Printf("flines filter lines from a file in a determinate range.\n\n")
	fmt.Printf("Options:\n")
	fmt.Printf("\t-f\t\t File to read. If not used the file is read from STDIN\n")
	fmt.Printf("\t-s\t\t Which start line of the file will be show.\n")
	fmt.Printf("\t-e\t\t Which end line of the file will be show. If not specified the file will be print until the end\n")
	fmt.Printf("\t-h\t\t This help.\n")
}

func main() {
	var err error
	startLine := 0
	endLine := 0
	fileName := "/dev/stdin"
	log.SetFlags(0)

	emptyArgsSlice := []string{}
	args := append(emptyArgsSlice, os.Args[1:]...)
	args = append(args, "--")

	// outArgs is a label to do break in for, not in switch
outArgs:
	for args[0] != "--" {
		switch {
		case args[0] == "-s":
			startLine, err = strconv.Atoi(args[1])
			if err != nil {
				log.Fatal("Parameter Start (-s) is not a int")
			}
			args = append(emptyArgsSlice, args[2:]...)
		case args[0] == "-e":
			endLine, err = strconv.Atoi(args[1])
			if err != nil {
				log.Fatal("Parameter End (-e) is not a int")
			}
			args = append(emptyArgsSlice, args[2:]...)
		case args[0] == "-f":
			fileName = args[1]
			args = append(emptyArgsSlice, args[2:]...)
		case args[0] == "-h":
			help()
			return
			args = append(emptyArgsSlice, args[1:]...)
		default:
			log.Fatalln("Unknow Option")
			break outArgs
		}
	}

	if startLine > endLine && endLine != 0 {
		log.Fatal("Parameter Start (-s) is bigger than parameter End (-e). Please verify.")
	}

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	filterLines(scanner, startLine, endLine)
}
