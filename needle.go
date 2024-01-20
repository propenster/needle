package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func searchFiles(root, searchText string) error {
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			if err := processEachFile(path, searchText); err != nil {
				fmt.Printf("Error while processing file %s: %v\n", path, err)
			}
		}

		return nil

	})
}

func processEachFile(path, searchText string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	line := 0
	for scanner.Scan() {
		line++
		lineText := scanner.Text()
		if strings.Contains(strings.ToLower(lineText), strings.ToLower(searchText)) {
			fmt.Printf("Found: '%s' in '%s' at line %d\n", searchText, path, line)
		}
	}
	return scanner.Err()
}

func main() {
	fmt.Println("Welcome to needle")
	fmt.Println("Needle is a utility tool that I use to scour through hundreds of log files for particular search terms")

	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <directory> <searchTerm>")
		os.Exit(1)
	}
	startDir := os.Args[1]
	searchText := os.Args[2]
	//fmt.Printf("Start/Root Directory: %s, SearchTerm: %s\n", startDir, searchText)
	err := searchFiles(startDir, searchText)
	if err != nil {
		fmt.Printf("Error while searching for term: %s %v\n", searchText, err)
	}

}
