package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"
)

type filelist []string

func (f *filelist) String() string {
	return strings.Join(*f, ", ")

}

func (f *filelist) Set(value string) error {

	*f = append(*f, value)
	return nil
}
func main() {

	var files filelist
	flag.Var(&files, "f", "Input your filenames")
	searchString := flag.String("search", "foo", "text to search")
	flag.Parse()
	if len(files) == 0 {
		fmt.Println("Please provide at least one file to search")
		return
	}
	if *searchString == "" {
		fmt.Println("Please provide a search string")

	}
	fmt.Println("files are", files, *searchString, "search text is ", *searchString)
	var wg sync.WaitGroup
	for _, v := range files {
		wg.Add(1)
		go search(v, *searchString, &wg)

	}

	wg.Wait()

}

func search(file string, s string, w *sync.WaitGroup) {
	defer w.Done()
	filename, err := os.Open(file)
	if err != nil {
		fmt.Println("Error opening file", file)
		return
	}
	defer filename.Close()

	scanner := bufio.NewScanner(filename)
	lineno := 1
	found := false
	for scanner.Scan() {

		line := scanner.Text()
		if strings.Contains(line, s) {

			fmt.Printf("Found the search text in file %v at linenumber %v \n", file, lineno)
			found = true
		}
		lineno++

	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Found error reading file", file)

	}
	if !found {
		fmt.Println("\n", "Search string was not found in ", file)
	}

}
